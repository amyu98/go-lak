package gamedb

import (
	"encoding/json"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/amyu98/go-lak/src/models"
)

func LoadGame(slug string) (*models.State) {
	allStates := readDB()
	return allStates[slug]
}

func SaveGame(s *models.State) {
	allStates := readDB()

	if allStates == nil {
		allStates = make(map[string]*models.State)
	}
	allStates[s.Slug] = s
	writeDB(allStates)
}

func readDB() (map[string]*models.State) {
	pathToDB := dbPath()
	fileContents, err := os.Open(pathToDB)
	if err != nil {
		panic(err)
	}
	defer fileContents.Close()
	byteValue, _ := io.ReadAll(fileContents)
	var states map[string]*models.State
	json.Unmarshal(byteValue, &states)
	return states
}

func writeDB(states map[string]*models.State) {
	pathToDB := dbPath()
	if _, err := os.Stat(pathToDB); os.IsNotExist(err) {
		os.Create(pathToDB)
	}

	err := os.WriteFile(pathToDB, []byte(statesToJSON(states)), 0644)
    if err != nil {
		panic(err)
    }
}

func statesToJSON(states map[string]*models.State) (string) {
	json, err := json.Marshal(states)
	if err != nil {
		panic(err)
	}
	return string(json)
}

func GenerateSlug(n int) (string) {
		rand.Seed(time.Now().UnixNano())
		var letters = []rune("abcdefghijklmnopqrstuvwxyz")
		b := make([]rune, n)
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
		return string(b)
}

func dbPath() (string) {
	// Print working directory
	wd := os.Getenv("PWD")
	return wd + "/src/gamedb/db.json"
}