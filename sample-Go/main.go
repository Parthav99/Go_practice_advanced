package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()                  //
	server.GET("/", func(ctx *gin.Context) { //handler

		ctx.JSON(http.StatusOK, gin.H{
			"message": "hi",
		})
	})
	// server.Use(gin.Logger())
	server.Run(":3000")
}
