package apiserver

import (
	"fmt"
	"github.com/amyu98/go-lak/src/gamehandler"
	"log"
	"net/http"
)

func Run() {
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	path := r.URL.Path
	switch path {
	case "/":
		fmt.Fprintf(w, "Hello, %q", r.URL.Path)
	case "/new_game":
		gamehandler.NewGame(w, r)
	case "/possible_moves":
		gamehandler.PossibleMoves(w, r)
	}
	w.Header().Set("Content-Type", "application/json")
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "text/html; charset=utf-8")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
