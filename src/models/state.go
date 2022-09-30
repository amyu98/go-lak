package models

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
	Tick		  int
	Logs		  []GameLog
}

type GameLog struct {
	Tick int
	Msg string
}