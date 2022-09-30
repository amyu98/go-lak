package gamehandler

import (
	"encoding/json"
	"net/http"
	"github.com/amyu98/go-lak/src/gamelogger"
	"github.com/amyu98/go-lak/src/constants"
	"github.com/amyu98/go-lak/src/gamedb"
	"github.com/amyu98/go-lak/src/models"
	"github.com/amyu98/go-lak/src/statemanager"
)

func NewGame(w http.ResponseWriter) {
	initState := constants.GetInitState()
	slug := gamedb.GenerateSlug(10)
	startingPlayer := "black"
	state := models.State{
		Slug:          slug,
		Board:         initState,
		CurrentPlayer: startingPlayer,
		WhiteJail:     0,
		BlackJail:     0,
		DiceRoll:      [2]int{0, 0},
		SelectedCell:  -99,
		PossibleMoves: []int{},
		Tick:          0,
		Logs:          []models.GameLog{},
	}
	gamelogger.LogMessage(&state, "Starting new game with slug: " + slug)
	gamelogger.LogMessage(&state, "Starting player: " + startingPlayer)
	statemanager.RollDice(&state)
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

func SelectCell(w http.ResponseWriter, s *models.State, cellIndex int) {
	piecesInCell := getPieceCountInCell(s, cellIndex)
	if piecesInCell == 0 {
		s.SelectedCell = -99
		s.PossibleMoves = []int{}
	} else {
		statemanager.SelectCell(s, cellIndex)
		possibleMoves := statemanager.GetPossibleMoves(s)
		s.PossibleMoves = possibleMoves
		if len(possibleMoves) == 0 {
			s.SelectedCell = -99
		}
	}
	res, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	gamedb.SaveGame(s)
	w.Write(res)
}

func getPieceCountInCell(s *models.State, cellIndex int) int {
	if cellIndex == -1 {
		return s.BlackJail
	}
	if cellIndex == 24 {
		return s.WhiteJail
	}
	return s.Board[cellIndex].BlackPieces + s.Board[cellIndex].WhitePieces
}

func MovePiece(w http.ResponseWriter, s *models.State, targetCell int) {
	isPossibleMove := false
	for _, possibleMove := range s.PossibleMoves {
		if possibleMove == targetCell {
			isPossibleMove = true
			break
		}
	}
	if isPossibleMove {
		statemanager.MovePiece(s, targetCell)
	} else {
		s.SelectedCell = -99
		s.PossibleMoves = []int{}
	}
	res, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	gamedb.SaveGame(s)
	w.Write(res)
}
