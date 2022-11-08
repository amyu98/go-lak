package gameai

import (
	"encoding/json"
	// "encoding/json"
	"fmt"
	"sort"

	"github.com/amyu98/go-lak/src/models"
	"github.com/amyu98/go-lak/src/statemanager"
)

type Proposal struct {
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

type Move []Action

type PlanMetaData struct {
	color         string
	optionalMoves []Move
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
	stateClone := state.CloneState()
	stateAsMap := stateClone.ToMap()
	actionsList := [][]Action{}
	planMetaData := PlanMetaData{color: stateClone.CurrentPlayer, optionalMoves: []Move{}}
	for key, value := range stateAsMap {
		if value > 0 {
			fmt.Println("Generating actions for ", key)
			addActionsFromEntryPoint(stateClone, key, Move{}, &planMetaData)
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

func addActionsFromEntryPoint(state models.State, entryPoint int, move Move, planMetaData *PlanMetaData) {

	stateClone := state.CloneState()

	stateClone.SelectedCell = -99
	statemanager.SelectCell(&stateClone, entryPoint)
	possibleMoves := statemanager.GetPossibleMoves(&stateClone)

	if !canMove(&stateClone, entryPoint, possibleMoves) {
		return
	}

	fmt.Println("Number of possible moves: ", len(possibleMoves))

	for _, possibleMove := range possibleMoves {

		fmt.Println("Possible move: ", possibleMove)

		stateClone2 := stateClone.CloneState()
		statemanager.MovePiece(&stateClone2, possibleMove)
		moveClone := make([]Action, len(move))
		copy(moveClone, move)
		moveClone = append(moveClone, Action{From: entryPoint, To: possibleMove})

		stillAiTurn := stateClone2.CurrentPlayer == planMetaData.color
		if stillAiTurn {
			fmt.Println("Still AI turn, going deeper")
			stateAsMap := stateClone2.ToMap()
			for key, value := range stateAsMap {
				if value > 0 {
					addActionsFromEntryPoint(stateClone2, key, moveClone, planMetaData)
				}
			}
		} else {
			// fmt.Println("Move: ", moveClone)
		}
	}
}

func canMove(state *models.State, entryPoint int, possibleMoves []int) bool {
	if len(possibleMoves) < 1 {
		return false
	}
	if *state.FriendlyJail() > 0 {
		if entryPoint != state.FriendlyJailIndex() {
			return false
		}
	}
	return true
}

func asJsonString(o any) string {
	output, _ := json.Marshal(o)
	return string(output)

}
