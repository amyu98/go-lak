package server

import (
	"fmt"
	"github.com/amyu98/go-lak/game_handler"
	"log"
	"net/http"
)

func Run() {
	print("Starting server...")
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	path := r.URL.Path
	switch path {
	case "/":
		print("Echo!")
		fmt.Fprintf(w, "Hello, %q", r.URL.Path)
	case "/new_game":
		print("New game!")
		game_handler.NewGameHandler(w, r)
	}
	w.Header().Set("Content-Type", "application/json")
}

func enableCors(w *http.ResponseWriter) {
	print("Enabling CORS...")
	(*w).Header().Set("Content-Type", "text/html; charset=utf-8")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
