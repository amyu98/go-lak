package gameai

import (
	"github.com/amyu98/go-lak/src/models"
	"github.com/amyu98/go-lak/src/statemanager"
)

const (
	MaxDepth         = 3
	AggressiveWeight = 1
	DefensiveWeight  = 2
)

type Bot struct {
	Depth            int
	AggressiveWeight int
	DefensiveWeight  int
	Side			 string
}

type Decision struct {
	CellIndex int
	TargetCell int
	Profit float32
}

func NewBot(side string) *Bot {
	return &Bot{
		Depth:            MaxDepth,
		AggressiveWeight: AggressiveWeight,
		DefensiveWeight:  DefensiveWeight,
		Side:			 side,
	}
}

func (b *Bot) GetDecision(s *models.State) Decision {
	possibleActions := b.getPossibleActions(s)
}

func (b *Bot) getPossibleActions(s *models.State) []Decision {
	usableMoves := statemanager.GetUsableMoves(s)

}

