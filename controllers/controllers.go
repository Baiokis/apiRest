package controllers

import (
	"apiRest/database"
	"apiRest/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(models.Personalidades)
}

func PersonalidadePorId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	for _, p := range models.Personalidades {
		if strconv.Itoa(p.Id) == id {
			json.NewEncoder(w).Encode(p)
			return
		}
	}

}
