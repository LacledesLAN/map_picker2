package routes

import (
	"github.com/gorilla/mux"
	"github.com/lacledeslan/map_picker2/pkg/controllers"
)


var RegisterGameRoutes = func(router *mux.Router) {
	router.HandleFunc("/game/", controllers.CreateGame).Methods("POST")
	router.HandleFunc("/game/{gameId}", controllers.DeleteGame).Methods("DELETE")
	router.HandleFunc("/games" , controllers.GetAllGames).Methods("GET")
	router.HandleFunc("/game/{gameId}", controllers.GetGameByID).Methods("GET")
	router.HandleFunc("/game?search={gameName}", controllers.GetGameByID).Methods("GET")
	router.HandleFunc("/game/{gameId}", controllers.UpdateGame).Methods("PUT")
}