package statemanager

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"github.com/amyu98/go-lak/src/gamelogger"
	"github.com/amyu98/go-lak/src/models"
)

func GetPossibleMoves(s *models.State) []int {
	usableMoves := GetUsableMoves(s)
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
			if s.FriendlyInEndGame() {
				possibleCells = append(possibleCells, s.FriendlyGoalIndex())
			}
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

func GetMovesFromDice(diceRoll [2]int) []int {
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
	if enemiesAtTarget == nil {
		enemiesAtTarget = new(int)
	}
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

	gamelogger.LogMessage(s, s.CurrentPlayer+" Moved piece from "+strconv.Itoa(s.SelectedCell)+" to "+strconv.Itoa(targetCell))

	victory := checkVictory(s)
	if victory {
		gamelogger.LogMessage(s, s.CurrentPlayer+" wins!")
		s.Victory = s.CurrentPlayer
		return
	}


	moveValue := int(math.Abs(float64(int(targetCell - s.SelectedCell))))
	if targetCell == 30 || targetCell == -30 {
		moveValue = selectGoalScorer(s)
	}
	s.UsedMoves = append(s.UsedMoves, moveValue)

	s.PossibleMoves = GetPossibleMoves(s)

	if len(s.UsedMoves) == len(GetMovesFromDice(s.DiceRoll)) {
		nextTurn(s)
	}

	if len(s.PossibleMoves) == 0 || *friendsAtHomeCell == 0 {
		s.SelectedCell = -99
		s.PossibleMoves = []int{}
	}
	skipImpossibleTurn(s)

	logUsableMoves(s)
}

func getAllPotentialMoves(s *models.State) map[int][]int {
	poentialMoves := make(map[int][]int)
	for i, _ := range s.Board {
		if *s.FriendsAt(i) == 0 {
			continue
		}
		cloneState := reflect.ValueOf(s).Elem().Interface().(models.State)
		SelectCell(&cloneState, i)
		possibleMoves := GetPossibleMoves(&cloneState)
		poentialMoves[i] = possibleMoves
	}

	if *s.FriendlyJail() > 0 {
		cloneState := reflect.ValueOf(s).Elem().Interface().(models.State)
		SelectCell(&cloneState, s.FriendlyJailIndex())
		possibleMoves := GetPossibleMoves(&cloneState)
		poentialMoves[s.FriendlyJailIndex()] = possibleMoves
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

	gamelogger.LogMessage(s, "Next turn: "+s.CurrentPlayer)
	RollDice(s)

	skipImpossibleTurn(s)
}

func RollDice(s *models.State) {
	s.DiceRoll[0] = rand.Intn(6) + 1
	s.DiceRoll[1] = rand.Intn(6) + 1
	gamelogger.LogMessage(s, "Dice roll: "+strconv.Itoa(s.DiceRoll[0])+" "+strconv.Itoa(s.DiceRoll[1]))
}

func GetUsableMoves(s *models.State) []int {
	usableMoves := GetMovesFromDice(s.DiceRoll)
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
	usableMoves := GetUsableMoves(s)
	gamelogger.LogMessage(s, "Usable moves: "+strings.Trim(strings.Join(strings.Fields(fmt.Sprint(usableMoves)), ", "), "[]"))
}

func skipImpossibleTurn(s *models.State) {
	potenitalMoves := getAllPotentialMoves(s)
	anyPotentialMoves := false

	for _, moves := range potenitalMoves {
		if len(moves) > 0 {
			anyPotentialMoves = true
			break
		}
	}

	stuckInJail := *s.FriendlyJail() > 0 && len(potenitalMoves[s.FriendlyJailIndex()]) == 0

	if !anyPotentialMoves || stuckInJail {
		log := "No possible moves for " + s.CurrentPlayer + ", skipping turn"
		gamelogger.LogMessage(s, log)
		nextTurn(s)
	}
}

func selectGoalScorer(s *models.State) int {
	// TODO! only logic not implemented yet
	// What scorer can score what? Also forbiden to score is not active on this on one
	lowestDelta := 100
	usableMoves := GetUsableMoves(s)
	moveValue := 0
	for _, move := range usableMoves {
		deltaToGoal := int(math.Abs(float64(int(s.EnemyJailIndex() - s.SelectedCell))))
		if deltaToGoal > move {
			continue
		}
		if deltaToGoal < lowestDelta {
			lowestDelta = deltaToGoal
			moveValue = move
		}
	}
	return moveValue
}

func checkVictory(s *models.State) bool {
	for i, _ := range s.Board {
		if *s.FriendsAt(i) != 0 {
			return false
		}
	}
	return *s.FriendlyJail() == 0
}
