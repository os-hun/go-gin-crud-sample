package main

import (
	"github.com/os-hun/go-gin-crud-sample/db"
	"github.com/os-hun/go-gin-crud-sample/config/routes"
)

func main() {
	db.Init()
	routes.Init()
	db.Close()
}
