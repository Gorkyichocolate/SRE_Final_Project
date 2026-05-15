package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"sre/internal/metrics"
	"sre/internal/service"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterHandler(c *gin.Context) {

	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})

		return
	}

	err := service.RegisterUser(
		req.Username,
		req.Email,
		req.Password,
	)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not create user",
		})

		return
	}

	metrics.AuthRequestsTotal.WithLabelValues(
		"register",
	).Inc()

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created",
	})
}

func LoginHandler(c *gin.Context) {

	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})

		return
	}

	token, err := service.LoginUser(
		req.Email,
		req.Password,
	)

	if err != nil {

		metrics.AuthRequestsTotal.WithLabelValues(
			"login_failed",
		).Inc()

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid credentials",
		})

		return
	}

	metrics.AuthRequestsTotal.WithLabelValues(
		"login_success",
	).Inc()

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
