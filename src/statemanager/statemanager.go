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
	s.SelectedCell = cellIndex
}

func MovePiece(s *models.State, targetCell int) {
	friendsAtHomeCell := s.FriendsAt(s.SelectedCell)
	enemiesAtTarget := s.EnemiesAt(targetCell)
	friendliesAtTarget := s.FriendsAt(targetCell)

	if *friendsAtHomeCell == 0 {
		gamelogger.LogMessage(s, "No pieces at home cell")
		return
	}
	if *enemiesAtTarget > 1 {
		gamelogger.LogMessage(s, "Target cell is full")
		return
	}

	*friendsAtHomeCell--
	*friendliesAtTarget++

	if *enemiesAtTarget == 1 {
		*enemiesAtTarget--
		*s.EnemyJail()++
	}

	gamelogger.LogMessage(s, s.CurrentPlayer + " Moved piece from " + strconv.Itoa(s.SelectedCell) + " to " + strconv.Itoa(targetCell))
	moveValue := int(math.Abs(float64(int(targetCell - s.SelectedCell))))
	s.UsedMoves = append(s.UsedMoves, moveValue)

	s.PossibleMoves = GetPossibleMoves(s)

	if len(s.PossibleMoves) == 0 || *friendsAtHomeCell == 0 {
		s.SelectedCell = -99
		s.PossibleMoves = []int{}
	}
	if len(s.UsedMoves) == len(getMovesFromDice(s.DiceRoll)) {
		nextTurn(s)
	}

	potenitalMoves := getAllPotentialMoves(s)
	gamelogger.LogMessage(s, "Potential moves: " + fmt.Sprint(potenitalMoves))
	var anyPotentialMoves bool
	for _, moves := range potenitalMoves {
		if len(moves) > 0 {
			anyPotentialMoves = true
			break
		}
	}
	// potenitalMoves at position 5
	stuckInJail := *s.FriendlyJail() > 0 && len(potenitalMoves[s.FriendlyJailIndex()]) == 0
	if !anyPotentialMoves || stuckInJail {
		log := "No possible moves for " + s.CurrentPlayer + ", skipping turn"
		gamelogger.LogMessage(s, log)
		nextTurn(s)
	}
	logUsableMoves(s)
}

func getAllPotentialMoves(s *models.State) (map[int][]int){
	poentialMoves := make(map[int][]int)
	for i,_ := range s.Board {
		if *s.FriendsAt(i) == 0 {
			continue
		}
		cloneState := reflect.ValueOf(s).Elem().Interface().(models.State)
		SelectCell(&cloneState, i)
		possibleMoves := GetPossibleMoves(&cloneState)
		poentialMoves[i] = possibleMoves
	}

	cloneState := reflect.ValueOf(s).Elem().Interface().(models.State)
	SelectCell(&cloneState, s.FriendlyJailIndex())
	possibleMoves := GetPossibleMoves(&cloneState)
	poentialMoves[s.FriendlyJailIndex()] = possibleMoves

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
