package handler

import (
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"sre/internal/metrics"
)

func HeavyHandler(c *gin.Context) {

	start := time.Now()

	for i := 0; i < 100000000; i++ {
		_ = math.Sqrt(float64(i))
	}

	metrics.HeavyRequestsTotal.Inc()

	c.JSON(http.StatusOK, gin.H{
		"message": "heavy request completed",
		"duration": time.Since(start).String(),
	})
}