package statemanager

import (
	"math"
	"math/rand"
	"strings"
	"fmt"
	"reflect"
	"strconv"
	"github.com/amyu98/go-lak/src/gamelogger"
	"github.com/amyu98/go-lak/src/models"
)

func GetPossibleMoves(s *models.State) []int {
	usableMoves := getUsableMoves(s)
	fmt.Println("Usable moves: ", usableMoves)
	cellIndex := s.SelectedCell
	possibleCells := []int{}
	for _, dice := range usableMoves {
		var possibleCell int
		if s.CurrentPlayer == "black" {
			possibleCell = cellIndex + dice
		} else {
			possibleCell = cellIndex - dice
		}
		if possibleCell > 23 || possibleCell < 0 {
			continue
		}
		if s.CurrentPlayer == "black" {
			if s.Board[possibleCell].WhitePieces > 1 {
				continue
			}
		} else {
			if s.Board[possibleCell].BlackPieces > 1 {
				continue
			}
		}
		possibleCells = append(possibleCells, possibleCell)
	}
	return possibleCells
}

func getMovesFromDice(diceRoll [2]int) []int {
	if diceRoll[0] == diceRoll[1] {
		return []int{diceRoll[0], diceRoll[0], diceRoll[0], diceRoll[0]}
	} else {
		return []int{diceRoll[0], diceRoll[1]}
	}
}

func SelectCell(s *models.State, cellIndex int) {
	if s.CurrentPlayer == "black" && s.BlackJail > 0 {
		if cellIndex != -1 {
			return
		}
	} else if s.CurrentPlayer == "white" && s.WhiteJail > 0 {
		if cellIndex != 24 {
			return
		}
	}
	s.SelectedCell = cellIndex
}

func MovePiece(s *models.State, targetCell int) {
	cellIndex := s.SelectedCell
	var homeCell int
	if s.CurrentPlayer == "black" {
		if cellIndex == -1 {
			homeCell = s.BlackJail
		} else {
			homeCell = s.Board[cellIndex].BlackPieces
		}
	} else {
		if cellIndex == 24 {
			homeCell = s.WhiteJail
		} else {
			homeCell = s.Board[cellIndex].WhitePieces
		}
	}
	if s.CurrentPlayer == "black" {
		if s.Board[targetCell].WhitePieces > 1 || homeCell < 1 {
			panic("Invalid move")
		}
		s.Board[targetCell].BlackPieces++
		homeCell--
		if s.Board[targetCell].WhitePieces == 1 {
			s.Board[targetCell].WhitePieces--
			s.WhiteJail++
		}
	} else {
		if s.Board[targetCell].BlackPieces > 1 || homeCell < 1 {
			panic("Invalid move")
		}
		s.Board[targetCell].WhitePieces++
		homeCell--
		if s.Board[targetCell].BlackPieces == 1 {
			s.Board[targetCell].BlackPieces--
			s.BlackJail++
		}
	}
	if s.CurrentPlayer == "black" {
		if cellIndex == -1 {
			s.BlackJail = homeCell
		} else {
			s.Board[cellIndex].BlackPieces = homeCell
		}
	} else {
		if cellIndex == 24 {
			s.WhiteJail = homeCell
		} else {
			s.Board[cellIndex].WhitePieces = homeCell
		}
	}
	moveValue := int(targetCell - cellIndex)
	s.UsedMoves = append(s.UsedMoves, int(math.Abs(float64(moveValue))))
	s.PossibleMoves = GetPossibleMoves(s)
	if len(s.PossibleMoves) == 0 || homeCell == 0 {
		s.SelectedCell = -99
		s.PossibleMoves = []int{}
	}
	if len(s.UsedMoves) == len(getMovesFromDice(s.DiceRoll)) {
		nextTurn(s)
	}
	potenitalMoves := getAllPotentialMoves(s)
	var anyPotentialMoves bool
	for _, moves := range potenitalMoves {
		if len(moves) > 0 {
			anyPotentialMoves = true
			break
		}
	}
	if !anyPotentialMoves {
		log := "No possible moves for " + s.CurrentPlayer + ", skipping turn"
		gamelogger.LogMessage(s, log)
		nextTurn(s)
	}
	logUsableMoves(s)
}

func getAllPotentialMoves(s *models.State) (map[int][]int){
	// Map of int to int array:
	poentialMoves := make(map[int][]int)
	for i, cell := range s.Board {
		if s.CurrentPlayer == "black" && cell.BlackPieces == 0 ||
			s.CurrentPlayer == "white" && cell.WhitePieces == 0 {
			continue
		}
		cloneState := reflect.ValueOf(s).Elem().Interface().(models.State)
		SelectCell(&cloneState, i)
		possibleMoves := GetPossibleMoves(&cloneState)
		poentialMoves[i] = possibleMoves
	}
	if s.CurrentPlayer == "black" && s.BlackJail > 0 {
		cloneState := reflect.ValueOf(s).Elem().Interface().(models.State)
		SelectCell(&cloneState, -1)
		possibleMoves := GetPossibleMoves(&cloneState)
		poentialMoves[-1] = possibleMoves
	} else if s.CurrentPlayer == "white" && s.WhiteJail > 0 {
		cloneState := reflect.ValueOf(s).Elem().Interface().(models.State)
		SelectCell(&cloneState, 24)
		possibleMoves := GetPossibleMoves(&cloneState)
		poentialMoves[24] = possibleMoves
	}
	return poentialMoves
}


func nextTurn(s *models.State) {
	s.SelectedCell = -99
	s.UsedMoves = []int{}
	s.PossibleMoves = []int{}
	s.Tick++
	if s.CurrentPlayer == "black" {
		s.CurrentPlayer = "white"
	} else {
		s.CurrentPlayer = "black"
	}

	gamelogger.LogMessage(s, "Next turn: " + s.CurrentPlayer)
	RollDice(s)

	if s.CurrentPlayer == "black" && s.BlackJail > 0 || s.CurrentPlayer == "white" && s.WhiteJail > 0 {
		if s.CurrentPlayer == "black" {
			s.SelectedCell = -1
		} else {
			s.SelectedCell = 24
		}
		possibleMoves := GetPossibleMoves(s)

		if len(possibleMoves) == 0 {
			gamelogger.LogMessage(s, "No possible moves for " + s.CurrentPlayer)
			gamelogger.LogMessage(s, "Returning turn to other player")
			nextTurn(s)
		}
	}
}

func RollDice(s *models.State) {
	s.DiceRoll[0] = rand.Intn(6) + 1
	s.DiceRoll[1] = rand.Intn(6) + 1
	gamelogger.LogMessage(s, "Dice roll: " + strconv.Itoa(s.DiceRoll[0]) + " " + strconv.Itoa(s.DiceRoll[1]))
}

func getUsableMoves(s *models.State) ([]int) {
	usableMoves := getMovesFromDice(s.DiceRoll)
	for _, usedMove := range s.UsedMoves {
		for i, move := range usableMoves {
			if move == usedMove {
				usableMoves = append(usableMoves[:i], usableMoves[i+1:]...)
				break
			}
		}
	}
	return usableMoves
}

func logUsableMoves(s *models.State) {
	usableMoves := getUsableMoves(s)
	gamelogger.LogMessage(s, "Usable moves: " + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(usableMoves)), ", "), "[]"))
}
