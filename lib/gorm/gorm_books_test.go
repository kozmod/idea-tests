package gorm

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"testing"
)

func Test(t *testing.T) {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	var Books []Book

	row := db.Find(&Books)
	if row.Error != nil {
		fmt.Println("error in querying")
	}

	for _, b := range Books {
		fmt.Println("book title: ", b.Title)
	}

}

type Book struct {
	ID    int64  `gorm:"primary_key"`
	Title string `json:"title"`
}
