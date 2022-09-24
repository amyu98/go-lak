package models

type State struct {
	BoardState [24]Cell
	CurrentPlayer string
	WhiteJail int
	BlackJail int
	DiceRoll [2]int
}