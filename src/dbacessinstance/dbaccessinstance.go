package dbaccessinstance

import (
	"sync"

	"github.com/amyu98/go-lak/src/gamedb"
	"github.com/amyu98/go-lak/src/models"
)

type DBAccessInstance struct {
	lock sync.Mutex
}

func (d *DBAccessInstance) ReadState(slug string) *models.State{
	d.lock.Lock()
	s := gamedb.LoadGame(slug)
	d.lock.Unlock()
	return s
}

func (d *DBAccessInstance) WriteState(state *models.State) {
	d.lock.Lock()
	gamedb.SaveGame(state)
	d.lock.Unlock()
}

func (d *DBAccessInstance) ReadAllStates() map[string]*models.State {
	d.lock.Lock()
	s := gamedb.ReadDB()
	d.lock.Unlock()
	return s
}

func (d *DBAccessInstance) GenerateSlug() string {
	return gamedb.GenerateSlug(10)
}