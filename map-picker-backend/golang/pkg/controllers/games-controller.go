package controllers

import (
	"bytes"
	"fmt"
	"github.com/lacledeslan/map_picker2/pkg/models"
	"github.com/thedevsaddam/gojsonq"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
)

var NewGame models.Game

var json = readFile()

func readFile() string {
	pwd, _ := os.Getwd()
	//read in json file
	filePath := filepath.Join(pwd, "data-responses", "games.json")

	//convert to absolute path to reduce issues with linux docker images/OS
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		fmt.Println(err)
	}

	//read in json file
	data, err := ioutil.ReadFile(absPath)
	if err != nil {
		fmt.Println(err)
	}

	return string(data)
}


func CreateGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "You want to create a game")
}

func DeleteGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "You want to delete a game")
}

func GetAllGames(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}


func GetGameByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameId, _ := strconv.Atoi(vars["gameId"])
	fmt.Println(gameId)

	var b bytes.Buffer

	//jq := gojsonq.New().JSONString(json)
	//res:= jq.From("games").WhereEqual("id", gameId).Get()

	gojsonq.New().JSONString(json).From("games.["+ vars["gameId"] +"]").Writer(&b)


	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, b.String())

}

func GetGameByName(w http.ResponseWriter, r *http.Request) {
	vars:= mux.Vars(r)
	gameName := vars["gameName"]


	var b bytes.Buffer
	gojsonq.New().JSONString(json).From("games").WhereEqual("name", gameName).Writer(&b)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, b.String())
}

func GetGameMapList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)
	//gameId, _ := strconv.Atoi(vars["gameId"])
	//mapListOnly := vars["maplistonly"]
	//fmt.Println(gameId)
	//fmt.Println(mapListOnly)

	//var b bytes.Buffer
	//
	//if  mapListOnly == "yes" {
	//	gojsonq.New().JSONString(json).From("games.["+ vars["gameId"] +"].gameMapList").Writer(&b)
	//} else {
	//	gojsonq.New().JSONString(json).From("games.["+ vars["gameId"] +"]").Writer(&b)
	//}
	//
	//w.Header().Set("Content-Type", "application/json")
	//fmt.Fprintf(w, b.String())
}

func UpdateGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "You want to update a game")
}
