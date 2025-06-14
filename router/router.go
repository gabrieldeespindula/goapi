package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gabrieldeespindula/goapi/handlers"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/records", handlers.GetRecords)
	r.GET("/records/:id", handlers.GetRecordByID)
	r.POST("/records", handlers.CreateRecord)
	r.PATCH("/records/:id", handlers.UpdateRecord)
	r.DELETE("/records/:id", handlers.DeleteRecord)
}
