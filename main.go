package main

import (
	"github.com/os-hun/go-gin-crud-sample/db"
	"github.com/os-hun/go-gin-crud-sample/config/server"
)

func main() {
	db.Init()
	server.Init()
	db.Close()
}
