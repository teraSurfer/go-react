package services

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

type Movie struct {
	gorm.Model
	Name string
	Director string
}

func InitialMigration () {
	db, err = gorm.Open("postgres", "host=localhost port=5432 dbname=movie_info user=postgres password=admin")
	if(err) {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	defer db.Close()

	db.AutoMigrate(&Movie{})
}

