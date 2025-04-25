package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DatabaseConnect() {
	stringdeConexao := "host=localhost user=root dbname=root password=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringdeConexao))
	if err != nil {
		log.Panic("Erro ao conectar no banco de dados: ")
	}
}
