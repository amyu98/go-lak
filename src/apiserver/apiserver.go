package apiserver

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	dbaccessinstance "github.com/amyu98/go-lak/src/dbacessinstance"
	"github.com/amyu98/go-lak/src/gamehandler"
)

type Apiserverinstance struct {
	db         *dbaccessinstance.DBAccessInstance
	controller *gamehandler.Controller
}

func (server *Apiserverinstance) Run(db *dbaccessinstance.DBAccessInstance) {
	server.db = db
	server.controller = &gamehandler.Controller{Db: db}
	http.HandleFunc("/", server.RootHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (server *Apiserverinstance) RootHandler(w http.ResponseWriter, r *http.Request) {
	server.enableCors(&w)
	path := r.URL.Path
	if strings.HasPrefix(path, "/api/v1") {
		server.apiControl(w, r)
	}
	w.Header().Set("Content-Type", "application/json")
}

func (server *Apiserverinstance) enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "text/html; charset=utf-8")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func (server *Apiserverinstance) apiControl(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	path = strings.TrimPrefix(path, "/api/v1/")

	if strings.HasPrefix(path, "slug=") {
		server.ControlStatefull(w, r, path)
	} else {
		server.ControlStateless(w, r, path)
	}
}

func (server *Apiserverinstance) ControlStateless(w http.ResponseWriter, r *http.Request, path string) {
	switch path {
	case "new_game":
		server.controller.NewGame(w)
	}
}

func (server *Apiserverinstance) ControlStatefull(w http.ResponseWriter, r *http.Request, path string) {

	slug := strings.SplitN(strings.TrimPrefix(path, "slug="), "/", 2)[0]
	path = strings.TrimPrefix(path, fmt.Sprintf("slug=%s/", slug))
	state := server.db.ReadState(slug)
	cellTarget := server.getTargetCell(r)

	switch path {
	case "get_game":
		server.controller.GetGame(w, state)
	case "select_cell":
		server.controller.SelectCell(w, state, cellTarget)
	case "move_piece":
		server.controller.MovePiece(w, state, cellTarget)
	case "ai_recommendation":
		server.controller.AiRecommendation(w, state)
	}
}

func (server *Apiserverinstance) getTargetCell(r *http.Request) int {
	targetCellAsString := r.URL.Query().Get("target_cell")
	targetCell, err := strconv.Atoi(targetCellAsString)
	if err != nil {
		targetCell = -1
	}
	return targetCell
}
