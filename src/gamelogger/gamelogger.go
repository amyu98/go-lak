package gamelogger

import (
	"github.com/amyu98/go-lak/src/models"
)

func LogMessage(s *models.State, message string) {
	s.Logs = append(s.Logs, models.GameLog{Tick: s.Tick, Msg: message})
}