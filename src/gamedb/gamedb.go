package gamedb

import (
	"encoding/json"
	"io"
	"os"
    "math/rand"
	"time"
	"github.com/amyu98/go-lak/src/models"
)

func LoadGame(slug string) (*models.State) {
	allStates := ReadDB()
	return allStates[slug]
}

func SaveGame(s *models.State) {
	allStates := ReadDB()
	if allStates == nil {
		allStates = make(map[string]*models.State)
	}
	allStates[s.Slug] = s
	WriteDB(allStates)
}

func ReadDB() (map[string]*models.State) {
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

func WriteDB(states map[string]*models.State) {
	pathToDB := dbPath()
	if _, err := os.Stat(pathToDB); os.IsNotExist(err) {
		os.Create(pathToDB)
	}
	file, err := os.OpenFile(pathToDB, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	json.NewEncoder(file).Encode(states)
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