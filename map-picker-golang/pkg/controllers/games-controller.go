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
	//ID, err:=strconv.ParseInt(gameId, 0 , 0)
	//if err != nil {
	//	fmt.Fprintf(w, "Error while parsing. '%v' is not a number", gameId)
	//}

	//game := gojsonq.New().FromString(json).Where("id", "=", ID)

	var b bytes.Buffer

	//jq := gojsonq.New().JSONString(json)
	//res:= jq.From("games").WhereEqual("id", gameId).Get()

	gojsonq.New().JSONString(json).From("games.["+ vars["gameId"] +"]").Writer(&b)

	//if err != nil {
	//	fmt.Fprintf(w, "Error")
	//}


	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, b.String())



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
