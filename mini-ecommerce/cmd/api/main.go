package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"sre/internal/config"
	"sre/internal/database"
	"sre/internal/metrics"
	"sre/internal/middleware"
	"sre/internal/routes"
)

func main() {

	cfg := config.LoadConfig()

	err := database.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}

	metrics.InitMetrics()

	router := gin.Default()

	router.SetTrustedProxies(nil)

	router.Use(middleware.PrometheusMiddleware())

	routes.SetupRoutes(router)

	router.Run(":" + cfg.Port)
}
