// main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Article - Our struct for all articles
type Player struct {
	Id    string `json:"id"`
	Pname string `json:"name"`
	Pteam string `json:"team"`
}
type Score struct {
	Id      string `json:"id,omitempty"`
	Matches string `json:"match"`
	Runs    string `json:"runs"`
	Wickets string `json:"wickets"`
}

var players []Player
var scores []Score

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllPlayers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllPlayers")
	json.NewEncoder(w).Encode(players)
}

func returnSinglePlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, player := range players {
		if player.Id == key {
			json.NewEncoder(w).Encode(player)
		}
	}
}

func createNewPlayer(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var player Player
	json.Unmarshal(reqBody, &player)
	// update our global Articles array to include
	// our new Article
	players = append(players, player)

	json.NewEncoder(w).Encode(player)
}
func createNewScore(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Score struct
	// append this to our Score array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var score Score
	json.Unmarshal(reqBody, &score)
	// update our global Articles array to include
	// our new Article
	scores = append(scores, score)

	json.NewEncoder(w).Encode(score)
}

func deletePlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, player := range players {
		if player.Id == id {
			players = append(players[:index], players[index+1:]...)
		}
	}

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/players", returnAllPlayers)
	myRouter.HandleFunc("/player", createNewPlayer).Methods("POST")
	myRouter.HandleFunc("/player/{id}/score", createNewScore).Methods("POST")
	myRouter.HandleFunc("/player/{id}", deletePlayer).Methods("DELETE")
	myRouter.HandleFunc("/player/{id}", returnSinglePlayer)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	players = []Player{
		{Id: "1", Pname: "VIRAT", Pteam: "RCB"},
		{Id: "2", Pname: "Dhoni", Pteam: "Chennai"},
	}
	scores = []Score{
		{"id": 1, "scores": [{"match": 1, "runs": 20, "wickets": 0}],
	},
		handleRequests(),
	}