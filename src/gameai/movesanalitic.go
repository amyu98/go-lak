package gameai

import (
	"fmt"
	"sort"

	"github.com/amyu98/go-lak/src/models"
	"github.com/amyu98/go-lak/src/statemanager"
)

type Move struct {
	// MoveId                    int
	Actions [][]Action
	// EatenEnemies              int
	// ExpousedFriendlies        int
	// ExpousedFriendliesAtRange int
	// NewHousesCreated          int
	// OddsToLosePieces          []any
	// OverallProfit             int
	// StateResultHash           string
}

type Action struct {
	From int
	To   int
}

func GenerateActionsForState(state *models.State) [][]Action {
	everyPossibleAction := generateEveryPossibleAction(state)
	for i := 0; i < len(everyPossibleAction); i++ {
		everyPossibleAction[i] = sortActions(everyPossibleAction[i])
	}
	removeDuplicates(&everyPossibleAction)
	return everyPossibleAction
}

func generateEveryPossibleAction(state *models.State) [][]Action {
	stateAsMap := state.ToMap()
	fmt.Println("Map: ", stateAsMap)
	actionsList := [][]Action{}
	for key, value := range stateAsMap {
		if value > 0 {
			fmt.Println("Generating actions for ", key)
			addActionsFromEntryPoint(state, key, []Action{}, &actionsList)
		}
	}
	return actionsList
}

func removeDuplicates(actionsList *[][]Action) {
	for i := 0; i < len(*actionsList); i++ {
		for j := i + 1; j < len(*actionsList); j++ {
			if isSameAction((*actionsList)[i], (*actionsList)[j]) {
				*actionsList = append((*actionsList)[:j], (*actionsList)[j+1:]...)
				j--
			}
		}
	}
}

func sortActions(actions []Action) []Action {
	sort.Slice(actions, func(i, j int) bool {
		if actions[i].From == -1 || actions[i].From == 24 {
			return true
		}
		if actions[i].To == 30 || actions[i].To == -30 {
			return false
		}
		iIsLowerThanJ := actions[i].From < actions[j].From || actions[i].To < actions[j].To
		return iIsLowerThanJ
	})
	return actions
}

func isSameAction(action1 []Action, action2 []Action) bool {
	if len(action1) != len(action2) {
		return false
	}
	for i := 0; i < len(action1); i++ {
		if action1[i] != action2[i] {
			return false
		}
	}
	return true
}

func addActionsFromEntryPoint(state *models.State, entryPoint int, previosActions []Action, actionsList *[][]Action) {
	fmt.Println("Entry point: ", entryPoint)
	color := state.CurrentPlayer
	stateClone := state.CloneState()
	statemanager.SelectCell(stateClone, entryPoint)
	possibleMoves := statemanager.GetPossibleMoves(stateClone)
	fmt.Println("Dice:", state.DiceRoll)
	fmt.Println("Possible moves: ", possibleMoves)
	if *state.FriendlyJail() > 0 {
		if entryPoint != state.FriendlyJailIndex() {
			fmt.Println("Friendly jail is not empty, cannot play")
			return
		}
	}
	for _, move := range possibleMoves {
		stateClone2 := stateClone.CloneState()
		fmt.Println("Moving from ", entryPoint, " to ", move)
		statemanager.MovePiece(stateClone2, move)
		fmt.Println("Actions before: ", previosActions)
		previosActionsClone := make([]Action, len(previosActions))
		copy(previosActionsClone, previosActions)
		previosActionsClone = append(previosActionsClone, Action{From: entryPoint, To: move})
		fmt.Println("Actions after: ", previosActionsClone)

		// On tv mode this can go wrong
		if stateClone2.CurrentPlayer == color {
			fmt.Println("Same player, adding actions")
			stateAsMap := stateClone2.ToMap()
			for key, value := range stateAsMap {
				if value > 0 {
					addActionsFromEntryPoint(stateClone2, key, previosActionsClone, actionsList)
				}
			}
		} else {
			*actionsList = append(*actionsList, previosActionsClone)
		}
	}
}
