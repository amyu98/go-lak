package gamehandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/amyu98/go-lak/src/models"
)

func NewGame(w http.ResponseWriter, r *http.Request) {

	initState := [24]models.Cell{
		models.Cell{Index: 0, WhitePieces: 0, BlackPieces: 2},
		models.Cell{Index: 1, WhitePieces: 0, BlackPieces: 0},
		models.Cell{Index: 2, WhitePieces: 0, BlackPieces: 0},
		models.Cell{Index: 3, WhitePieces: 0, BlackPieces: 0},
		models.Cell{Index: 4, WhitePieces: 0, BlackPieces: 0},
		models.Cell{Index: 5, WhitePieces: 5, BlackPieces: 0},
		models.Cell{Index: 6, WhitePieces: 0, BlackPieces: 0},
		models.Cell{Index: 7, WhitePieces: 3, BlackPieces: 0},
		models.Cell{Index: 8, WhitePieces: 0, BlackPieces: 0},
		models.Cell{Index: 9, WhitePieces: 0, BlackPieces: 0},
		models.Cell{Index: 10, WhitePieces: 0, BlackPieces: 0},
		models.Cell{Index: 11, WhitePieces: 0, BlackPieces: 5},
		models.Cell{Index: 12, WhitePieces: 5, BlackPieces: 0},
		models.Cell{Index: 13, WhitePieces: 0, BlackPieces: 0},
		models.Cell{Index: 14, WhitePieces: 0, BlackPieces: 0},
		models.Cell{Index: 15, WhitePieces: 0, BlackPieces: 0},
		models.Cell{Index: 16, WhitePieces: 0, BlackPieces: 3},
		models.Cell{Index: 17, WhitePieces: 0, BlackPieces: 0},
		models.Cell{Index: 18, WhitePieces: 0, BlackPieces: 5},
		models.Cell{Index: 19, WhitePieces: 0, BlackPieces: 0},
		models.Cell{Index: 20, WhitePieces: 0, BlackPieces: 0},
		models.Cell{Index: 21, WhitePieces: 0, BlackPieces: 0},
		models.Cell{Index: 22, WhitePieces: 0, BlackPieces: 0},
		models.Cell{Index: 23, WhitePieces: 2, BlackPieces: 0},
	}

	res, err := json.Marshal(initState)
	if err != nil {
		panic(err)
	}
	w.Write(res)
}

func PossibleMoves(w http.ResponseWriter, r *http.Request) {
	requestBody := r.Body
	var boardState models.State
	err := json.NewDecoder(requestBody).Decode(&boardState)
	selectedCell := r.URL.Query().Get("selectedCell")
	selectedCellAsInt, err := strconv.Atoi(selectedCell)
	if err != nil {
		panic(err)
	}

	possibleMoves := PossibleMovesByCell(boardState, selectedCellAsInt)
	res, err := json.Marshal(possibleMoves)
	if err != nil {
		panic(err)
	}
	w.Write(res)
}

func PossibleMovesByCell(boardState models.State, selectedCellIndex int) []int {

	possibleCells := []int{}
	// A possible cell - a cell withing one dice range, where there is no more than one enemy piece.
	for _, dice := range boardState.DiceRoll {
		// Declre possibleCell int:
		var possibleCell int
		if boardState.CurrentPlayer == "black" {
			possibleCell = selectedCellIndex + dice
		} else {
			possibleCell = selectedCellIndex - dice
		}
		if possibleCell > 23 || possibleCell < 0 {
			continue
		}
		if boardState.CurrentPlayer == "black" {
			if boardState.BoardState[possibleCell].WhitePieces > 1 {
				continue
			}
		} else {
			if boardState.BoardState[possibleCell].BlackPieces > 1 {
				continue
			}
		}
		possibleCells = append(possibleCells, possibleCell)
	}
	fmt.Println(possibleCells)
	return possibleCells
}
