package game_handler

import (
	"encoding/json"
	"github.com/amyu98/lak/models"
	"net/http"
)

func NewGameHandler(w http.ResponseWriter, r *http.Request) {

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
