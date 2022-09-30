package statemanager

import (
    "math/rand"
	"github.com/amyu98/go-lak/src/models"
)

func InitGame() {

}

func LoadGame() {

}

func GetPossibleMoves(s *models.State) ([]int) {
	cellIndex := s.SelectedCell
	possibleCells := []int{}
	for _, dice := range s.DiceRoll {
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

func SelectPiece(s *models.State, cellIndex int) {
	s.SelectedCell = cellIndex
}

func MovePiece(s *models.State, targetCell int) {
	cellIndex := s.SelectedCell
	if s.CurrentPlayer == "black" {
		if s.Board[targetCell].WhitePieces > 1 || s.Board[cellIndex].BlackPieces < 1{
			panic("Invalid move")
		}
		s.Board[targetCell].BlackPieces++
		s.Board[cellIndex].BlackPieces--
		if s.Board[targetCell].WhitePieces == 1 {
			s.Board[targetCell].WhitePieces--
			s.WhiteJail++
		}
	} else {
		if s.Board[targetCell].BlackPieces > 1 || s.Board[cellIndex].WhitePieces < 1{
			panic("Invalid move")
		}
		s.Board[targetCell].WhitePieces++
		s.Board[cellIndex].WhitePieces--
		if s.Board[targetCell].BlackPieces == 1 {
			s.Board[targetCell].BlackPieces--
			s.BlackJail++
		}
	}
}

func RollDice(s *models.State) {
	s.DiceRoll[0] = rand.Intn(6) + 1
	s.DiceRoll[1] = rand.Intn(6) + 1
}

