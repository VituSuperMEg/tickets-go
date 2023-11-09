package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "3640"
	dbname   = "bilheteira"
)

func InitDB() *gorm.DB {
	pg := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open("postgres", pg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexão bem-sucedida!🎯")
	return db
}
