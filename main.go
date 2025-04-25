package main

import (
	"apiRest/database"
	"apiRest/routes"
	"fmt"
)

func main() {
	database.DatabaseConnect()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleResquest()
}
