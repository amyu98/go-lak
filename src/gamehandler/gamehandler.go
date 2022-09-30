package gamehandler

import (
	"encoding/json"
	"net/http"
	"github.com/amyu98/go-lak/src/constants"
	"github.com/amyu98/go-lak/src/gamedb"
	"github.com/amyu98/go-lak/src/models"
	"github.com/amyu98/go-lak/src/statemanager"
)

func NewGame(w http.ResponseWriter) {
	initState := constants.GetInitState()
	slug := gamedb.GenerateSlug(10)
	state := models.State{
		Slug:          slug,
		Board:         initState,
		CurrentPlayer: "black",
		WhiteJail:     0,
		BlackJail:     0,
		DiceRoll:      [2]int{0, 0},
		SelectedCell:  -1,
		PossibleMoves: []int{},
	}
	gamedb.SaveGame(&state)
	res, err := json.Marshal(state)
	if err != nil {
		panic(err)
	}
	w.Write(res)
}

func RollDice(w http.ResponseWriter, s *models.State) {
	statemanager.RollDice(s)
	res, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	gamedb.SaveGame(s)
	w.Write(res)
}

func GetGame(w http.ResponseWriter, s *models.State) {
	res, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	w.Write(res)
}

func GetPossibleMoves(w http.ResponseWriter, s *models.State) {
	possibleMoves := statemanager.GetPossibleMoves(s)
	res, err := json.Marshal(possibleMoves)
	if err != nil {
		panic(err)
	}
	w.Write(res)
}

func SelectPiece(w http.ResponseWriter, s *models.State, cellIndex int) {
	statemanager.SelectPiece(s, cellIndex)
	res, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	gamedb.SaveGame(s)
	w.Write(res)
}

func MovePiece(w http.ResponseWriter, s *models.State, targetCell int) {
	statemanager.MovePiece(s, targetCell)
	res, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	gamedb.SaveGame(s)
	w.Write(res)
}
