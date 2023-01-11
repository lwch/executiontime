package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lwch/executiontime"
)

func main() {
	h := gin.Default()

	h.Use(
		executiontime.New(),
	)

	// Example ping request.
	h.GET("/ping", func(g *gin.Context) {
		g.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})

	h.Run()
}
