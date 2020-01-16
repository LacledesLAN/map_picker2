package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lacledeslan/map_picker2/pkg/models"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

var NewGame models.Game



func CreateGame(w http.ResponseWriter, r *http.Request) {
	//CreateGame := &models.Game{}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "You want to create a game")
}

func DeleteGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "You want to delete a game")
}

func GetAllGames(w http.ResponseWriter, r *http.Request) {
	pwd, _ := os.Getwd()
	//read in json file
	data, err := ioutil.ReadFile(pwd+ "/data-responses/games.json")
	if err != nil {
		fmt.Println(err)
	}

	dataBytes := []byte(data)

	obj := models.Games{}

	// unmarshall json file
	err = json.Unmarshal(dataBytes, &obj)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, obj)
}


func GetGameByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameId := vars["gameId"]

	ID, err:=strconv.ParseInt(gameId, 0 , 0)
	if err != nil {
		fmt.Fprintf(w, "Error while parsing. '%v' is not a number", gameId)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "You are looking for a game with the ID of:  %v", ID)
}

func GetGameByName(w http.ResponseWriter, r *http.Request) {
	vars:= mux.Vars(r)
	gameName := vars["gameName"]
	//if err != nil {
	//	fmt.Fprint(w, "You must provide a game name")
	//}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "You searched for %v", gameName)
}

func UpdateGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "You want to update a game")
}
