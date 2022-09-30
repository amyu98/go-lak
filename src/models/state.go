package models

import (
	"golang.org/x/exp/slices"
)

type State struct {
	Slug          string
	Board         [24]Cell
	CurrentPlayer string
	WhiteJail     int
	BlackJail     int
	DiceRoll      [2]int
	SelectedCell  int
	PossibleMoves []int
	UsedMoves     []int
	Tick          int
	Logs          []GameLog
}

func (s *State) EnemiesAt(cellIndex int) *int {
	if s.CurrentPlayer == "black" {
		if cellIndex == 24 {
			return &s.WhiteJail
		} else {
			return &s.Board[cellIndex].WhitePieces
		}
	} else {
		if cellIndex == -1 {
			return &s.BlackJail
		} else {
			return &s.Board[cellIndex].BlackPieces
		}
	}
}

func (s *State) FriendsAt(cellIndex int) *int {
	if s.CurrentPlayer == "black" {
		if cellIndex == -1 {
			return &s.BlackJail
		} else {
			return &s.Board[cellIndex].BlackPieces
		}
	} else {
		if cellIndex == 24 {
			return &s.WhiteJail
		} else {
			return &s.Board[cellIndex].WhitePieces
		}
	}
}

func (s *State) FriendlyJail() *int {
	if s.CurrentPlayer == "black" {
		return &s.BlackJail
	} else {
		return &s.WhiteJail
	}
}

func (s *State) EnemyJail() *int {
	if s.CurrentPlayer == "black" {
		return &s.WhiteJail
	} else {
		return &s.BlackJail
	}
}

func (s *State) FriendlyJailIndex() int {
	if s.CurrentPlayer == "black" {
		return -1
	} else {
		return 24
	}
}

func (s *State) EnemyJailIndex() int {
	if s.CurrentPlayer == "black" {
		return 24
	} else {
		return -1
	}
}

func (s *State) FriendlyInEndGame() bool {
	var home []int
	if s.CurrentPlayer == "black" {
		home = []int{0, 1, 2, 3, 4, 5}
	} else {
		home = []int{18, 19, 20, 21, 22, 23}
	}
	if *s.FriendlyJail() > 0 {
		return false
	}
	for _, cell := range s.Board {
		if !slices.Contains(home, cell.Index) && *s.FriendsAt(cell.Index) > 0 {
			return false
		}
	}
	return true
}

type GameLog struct {
	Tick int
	Msg  string
}
