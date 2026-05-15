package service

import (
	"sre/internal/model"
	"sre/internal/repository"
)

func GetProducts() ([]model.Product, error) {

	return repository.GetAllProducts()
}

func GetProduct(id string) (*model.Product, error) {

	return repository.GetProductByID(id)
}
