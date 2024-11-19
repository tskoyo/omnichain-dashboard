package main

import (
	"omnichain-dashboard/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Routes
	router.GET("/wallet/:address", handlers.GetWalletBalance)
	router.GET("/transactions/:address", handlers.GetTransactionHistory)

	// Start the server
	router.Run(":8080")
}
