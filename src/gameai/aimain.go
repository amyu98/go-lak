package gameai

import (
	"fmt"
	"time"

	dbaccessinstance "github.com/amyu98/go-lak/src/dbacessinstance"
	"github.com/amyu98/go-lak/src/models"
	"github.com/amyu98/go-lak/src/statemanager"
)

const (
	MaxDepth         = 3
	AggressiveWeight = 1
	DefensiveWeight  = 2
)

type Bot struct {
	Name             string
	Depth            int
	AggressiveWeight int
	DefensiveWeight  int
	Side             string
	Db               *dbaccessinstance.DBAccessInstance
}

func (bot *Bot) WaitForYourTurn() {
	for {
		states := bot.Db.ReadAllStates()
		bot.playPlayableStates(states)
		time.Sleep(1 * time.Second)
	}
}

func (bot *Bot) playPlayableStates(states  map[string]*models.State) {
	for slug := range states {
		state := states[slug]
		if state.CurrentPlayer == bot.Side {
			fmt.Println("Playing state: ", state.Slug)
			bot.playState(state)
		}
	}
}

func (bot *Bot) playState(state *models.State) {
	actions := GenerateActionsForState(state)
	fmt.Println("Actions: ", actions)
	action := ChooseAction(state, actions)
	bot.playAction(state, action)
}

func (bot *Bot) playAction(state *models.State, actions []Action) {
	for _, action := range actions {
		statemanager.SelectCell(state, action.From)
		bot.Db.WriteState(state)
		statemanager.MovePiece(state, action.To)
		bot.Db.WriteState(state)
	}
}


func ChooseAction(state *models.State, actions [][]Action) []Action {
	return actions[0]
}

