package routes

import (
	"net/http"

	"github.com/MAHcodes/lets_go/teendo/controllers"
)

func RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /item", controllers.GetAllItems)
	mux.HandleFunc("GET /item/{id}", controllers.GetItem)
	mux.HandleFunc("POST /item", controllers.CreateItem)
	mux.HandleFunc("PUT /item/{id}", controllers.UpdateItem)
	mux.HandleFunc("DELETE /item/{id}", controllers.DeleteItem)
	return mux
}
