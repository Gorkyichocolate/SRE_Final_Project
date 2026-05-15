package service

import (
	"sre/internal/model"
	"sre/internal/repository"
)

func CreateOrder(order model.Order) error {

	return repository.CreateOrder(order)
}

func GetOrders(userID int64) ([]model.Order, error) {

	return repository.GetOrdersByUserID(userID)
}
