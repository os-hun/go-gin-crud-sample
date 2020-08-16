package server

import (
	"github.com/gin-gonic/gin"
)

func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()

	u := r.Group("/users")
	{}
	
	p := r.Group("/posts")
	{}
	
	return r
}
