package routes

import (
	"apiRest/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleResquest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.PersonalidadePorId).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
