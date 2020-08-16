package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/os-hun/go-gin-crud-sample/models"
)

var (
	db *gorm.DB
	err error
)

// Init is initialize database
func Init() {
	db, err = gorm.Open("postgres", "host=db port=5432 user=go-gin-crud-sample dbname=go-gin-crud-sample password=go-gin-crud-sample sslmode=disable")
	if err != nil {
		panic(err)
	}
	autoMigration()
	user := models.User{
		ID: 1,
		Name: "aoki",
		Posts: []models.Post{{ID: 1, Content: "tweet1"}, {ID: 2, Content: "tweet2"}},
	}
	db.Create(&user)
}

// GetDB is called in models
func GetDB() *gorm.DB {
	return db
}

// Close is closing the database
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func autoMigration() {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Post{})
}
