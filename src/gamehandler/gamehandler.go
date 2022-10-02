package gamehandler

import (
	"encoding/json"
	"net/http"

	"github.com/amyu98/go-lak/src/constants"
	dbaccessinstance "github.com/amyu98/go-lak/src/dbacessinstance"
	"github.com/amyu98/go-lak/src/gameai"
	"github.com/amyu98/go-lak/src/gamelogger"
	"github.com/amyu98/go-lak/src/models"
	"github.com/amyu98/go-lak/src/statemanager"
)

type Controller struct {
	Db *dbaccessinstance.DBAccessInstance
}

func (controller *Controller) NewGame(w http.ResponseWriter) {
	initState := constants.GetInitState()
	slug := controller.Db.GenerateSlug()
	startingPlayer := "black"
	state := models.State{
		Slug:          slug,
		Board:         initState,
		CurrentPlayer: startingPlayer,
		WhiteJail:     0,
		BlackJail:     0,
		WhiteGoals:    0,
		BlackGoals:    0,
		DiceRoll:      [2]int{0, 0},
		SelectedCell:  -99,
		PossibleMoves: []int{},
		Tick:          0,
		Logs:          []models.GameLog{},
		Victory:       "",
		ShouldRecord:  true,
		PlayersType:   "AI-HUMAN",
	}
	gamelogger.LogMessage(&state, "Starting new game with slug: "+slug)
	gamelogger.LogMessage(&state, "Starting player: "+startingPlayer)
	statemanager.RollDice(&state)
	controller.Db.WriteState(&state)
	res, err := json.Marshal(state)
	if err != nil {
		panic(err)
	}
	w.Write(res)
}

func (controller *Controller) GetGame(w http.ResponseWriter, s *models.State) {
	res, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	w.Write(res)
}

func (controller *Controller) SelectCell(w http.ResponseWriter, s *models.State, cellIndex int) {
	piecesInCell := controller.getPieceCountInCell(s, cellIndex)
	if piecesInCell == 0 {
		s.SelectedCell = -99
		s.PossibleMoves = []int{}
	} else {

		canSelect := cellIndex != s.EnemyJailIndex() && *s.FriendsAt(cellIndex) > 0 &&
			(*s.FriendlyJail() == 0 || s.FriendlyJailIndex() == cellIndex)

		if canSelect {
			statemanager.SelectCell(s, cellIndex)
			possibleMoves := statemanager.GetPossibleMoves(s)
			s.PossibleMoves = possibleMoves
			if len(possibleMoves) == 0 {
				s.SelectedCell = -99
			}
		}
	}
	res, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	controller.Db.WriteState(s)
	w.Write(res)
}

func (controller *Controller) getPieceCountInCell(s *models.State, cellIndex int) int {
	if cellIndex == -1 {
		return s.BlackJail
	}
	if cellIndex == 24 {
		return s.WhiteJail
	}
	if cellIndex == 30 {
		return s.WhiteGoals
	}
	if cellIndex == -30 {
		return s.BlackGoals
	}
	return s.Board[cellIndex].BlackPieces + s.Board[cellIndex].WhitePieces
}

func (controller *Controller) MovePiece(w http.ResponseWriter, s *models.State, targetCell int) {
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
	controller.Db.WriteState(s)
	w.Write(res)
}

func (controller *Controller) AiRecommendation(w http.ResponseWriter, s *models.State) {
	gameai.GenerateActionsForState(s)
	m := s.ToMap()
	res, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	w.Write(res)
}
