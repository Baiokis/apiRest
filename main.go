package main

import (
	"apiRest/database"
	"apiRest/models"
	"apiRest/routes"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Albert Einstein", Historia: "Físico teórico, desenvolveu a teoria da relatividade."},
		{Id: 2, Nome: "Marie Curie", Historia: "Pioneira na pesquisa sobre radioatividade, primeira mulher a ganhar um Prêmio Nobel."},
		{Id: 3, Nome: "Isaac Newton", Historia: "Físico e matemático, formulou as leis do movimento e a lei da gravitação universal."},
	}

	database.DatabaseConnect()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleResquest()
}
