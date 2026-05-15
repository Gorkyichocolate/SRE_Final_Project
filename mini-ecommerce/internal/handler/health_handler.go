package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"sre/internal/database"
)

func HealthHandler(c *gin.Context) {

	err := database.DB.Ping()

	if err != nil {

		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "database unavailable",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
