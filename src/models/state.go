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
	WhiteGoals     int
	BlackGoals     int
	DiceRoll      [2]int
	SelectedCell  int
	PossibleMoves []int
	UsedMoves     []int
	Tick          int
	Logs          []GameLog
	Victory	   	  string
}

func (s *State) EnemiesAt(cellIndex int) *int {
	if s.EnemyJailIndex() == cellIndex {
		return s.EnemyJail()
	}
	if s.EnemyGoalIndex() == cellIndex {
		return s.EnemyGoal()
	}
	if cellIndex == s.FriendlyJailIndex() || cellIndex == s.FriendlyGoalIndex() {
		return nil
	}
	if s.CurrentPlayer == "black" {
		return &s.Board[cellIndex].WhitePieces
	} else {
		return &s.Board[cellIndex].BlackPieces
	}
}

func (s *State) FriendsAt(cellIndex int) *int {
	if s.FriendlyJailIndex() == cellIndex {
		return s.FriendlyJail()
	} else if s.FriendlyGoalIndex() == cellIndex {
		return s.FriendlyGoal()
	} else if s.EnemyJailIndex() == cellIndex {
		panic("Trying to get friends at enemy jail")
	} else if s.EnemyGoalIndex() == cellIndex {
		panic("Trying to get friends at enemy goal")
	} else {
		if s.CurrentPlayer == "black" {
			return &s.Board[cellIndex].BlackPieces
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

func (s *State) FriendlyGoalIndex() int {
	if s.CurrentPlayer == "black" {
		return -30
	} else {
		return 30
	}
}

func (s *State) EnemyGoalIndex() int {
	if s.CurrentPlayer == "black" {
		return 30
	} else {
		return -30
	}
}

func (s *State) FriendlyGoal() *int {
	if s.CurrentPlayer == "black" {
		return &s.BlackGoals
	} else {
		return &s.WhiteGoals
	}
}

func (s *State) EnemyGoal() *int {
	if s.CurrentPlayer == "white" {
		return &s.BlackGoals
	} else {
		return &s.WhiteGoals
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
