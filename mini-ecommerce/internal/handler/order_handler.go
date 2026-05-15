package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"sre/internal/metrics"
	"sre/internal/model"
	"sre/internal/service"
)

type CreateOrderRequest struct {
	ProductID int64 `json:"product_id"`
	Quantity  int   `json:"quantity"`
}

func CreateOrderHandler(c *gin.Context) {

	var req CreateOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})

		return
	}

	userID := c.GetInt64("user_id")

	order := model.Order{
		UserID:     userID,
		ProductID:  req.ProductID,
		Quantity:   req.Quantity,
		TotalPrice: 100 * float64(req.Quantity),
	}

	err := service.CreateOrder(order)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not create order",
		})

		return
	}

	metrics.OrdersCreatedTotal.Inc()

	c.JSON(http.StatusCreated, gin.H{
		"message": "order created",
	})
}

func GetOrdersHandler(c *gin.Context) {

	userID := c.GetInt64("user_id")

	orders, err := service.GetOrders(userID)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not get orders",
		})

		return
	}

	c.JSON(http.StatusOK, orders)
}
