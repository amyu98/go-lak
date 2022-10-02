package main

import (
	"github.com/amyu98/go-lak/src/apiserver"
	dbaccessinstance "github.com/amyu98/go-lak/src/dbacessinstance"
	"github.com/amyu98/go-lak/src/gameai"
)

func main() {

	db := dbaccessinstance.DBAccessInstance{}
	bot := gameai.Bot{Name: "BotDeLaBest", Db: &db, Side: "black"}
	go bot.WaitForYourTurn()

	server := apiserver.Apiserverinstance{}
	server.Run(&db)
}
