package routes

import (
	"github.com/gorilla/mux"
	"github.com/lacledeslan/map_picker2/pkg/controllers"
)



var RegisterMiscRoutes = func(router *mux.Router) {
		router.NotFoundHandler = controllers.HTTP404
}