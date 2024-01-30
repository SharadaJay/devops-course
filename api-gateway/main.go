package main

import (
	"example.com/api-gateway/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/messages", handlers.GetMessagesHandler)
	r.PUT("/state", handlers.PutStateHandler)
	r.GET("/state", handlers.GetStateHandler)
	r.GET("/run-log", handlers.GetRunLogHandler)
	r.GET("/mqstatistic", handlers.GetMQStatisticHandler)

	if err := r.Run(":8083"); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
