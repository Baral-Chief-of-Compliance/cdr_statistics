package main

import (
	"net/http"

	"github.com/cdr/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.Use(cors.Default())

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK!",
		})
	})

	server.POST("/statistics", controller.GetStatistics)

	server.Run(":8080")
}
