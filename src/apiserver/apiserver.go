package apiserver

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/amyu98/go-lak/src/gamedb"
	"github.com/amyu98/go-lak/src/gamehandler"
)

func Run() {
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	path := r.URL.Path
	if strings.HasPrefix(path, "/api/v1") {
		ApiControl(w, r)
	}
	w.Header().Set("Content-Type", "application/json")
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "text/html; charset=utf-8")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func ApiControl(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	path = strings.TrimPrefix(path, "/api/v1/")

	if strings.HasPrefix(path, "slug=") {
		ControlStatefull(w, r, path)
	} else {
		ControlStateless(w, r, path)
	}
}

func ControlStateless(w http.ResponseWriter, r *http.Request, path string) {
	switch path {
	case "new_game":
		gamehandler.NewGame(w)
	}
}

func ControlStatefull(w http.ResponseWriter, r *http.Request, path string) {

	slug := strings.SplitN(strings.TrimPrefix(path, "slug="), "/", 2)[0]
	path = strings.TrimPrefix(path, fmt.Sprintf("slug=%s/", slug))
	state := gamedb.LoadGame(slug)
	cellTarget := getTargetCell(r)

	switch path {
	case "get_game":
		gamehandler.GetGame(w, state)
	case "roll_dice":
		gamehandler.RollDice(w, state)
	case "select_cell":
		gamehandler.SelectCell(w, state, cellTarget)
	case "move_piece":
		gamehandler.MovePiece(w, state, cellTarget)
	}
}

func getTargetCell(r *http.Request) int {
	targetCellAsString := r.URL.Query().Get("target_cell")
	targetCell, err := strconv.Atoi(targetCellAsString)
	if err != nil {
		targetCell = -1
	}
	return targetCell
}
