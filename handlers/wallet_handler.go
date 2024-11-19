package handlers

import (
	"net/http"
	"omnichain-dashboard/services"

	"github.com/gin-gonic/gin"
)

func GetWalletBalance(c *gin.Context) {
	address := c.Param("address")
	balance, err := services.FetchWalletBalance(address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"address": address, "balance": balance})
}

func GetTransactionHistory(c *gin.Context) {}
