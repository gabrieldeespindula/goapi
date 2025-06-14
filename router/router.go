package router

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gabrieldeespindula/goapi/handlers"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	r.GET("/records", handlers.GetRecords)
	r.GET("/records/:id", handlers.GetRecordByID)
	r.POST("/records", handlers.CreateRecord)
	r.PATCH("/records/:id", handlers.UpdateRecord)
	r.DELETE("/records/:id", handlers.DeleteRecord)
}
