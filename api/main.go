package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gabrieldeespindula/goapi/db"
	"github.com/gabrieldeespindula/goapi/router"
)

func main() {
	db.Connect()

	r := gin.Default()
	router.SetupRoutes(r)

	r.Run(":8080")
}
