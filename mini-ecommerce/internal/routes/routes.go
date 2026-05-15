package routes

import (
	"github.com/gin-gonic/gin"

	"sre/internal/handler"
	"sre/internal/middleware"
)

func SetupRoutes(router *gin.Engine) {

	router.POST("/register",
		handler.RegisterHandler)

	router.POST("/login",
		handler.LoginHandler)

	router.GET("/products",
		handler.GetProductsHandler)

	router.GET("/products/:id",
		handler.GetProductByIDHandler)

	router.GET("/metrics",
		handler.MetricsHandler())

	router.GET("/heavy",
		handler.HeavyHandler)

	protected := router.Group("/")

	protected.Use(middleware.AuthMiddleware())

	{
		protected.GET("/protected",
			handler.ProtectedHandler)

		protected.POST("/orders",
			handler.CreateOrderHandler)

		protected.GET("/orders",
			handler.GetOrdersHandler)
	}
}
