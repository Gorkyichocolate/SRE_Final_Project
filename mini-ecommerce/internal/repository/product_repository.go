package repository

import (
	"sre/internal/database"
	"sre/internal/model"
)

func GetAllProducts() ([]model.Product, error) {

	rows, err := database.DB.Query(
		`
		SELECT id, name, description, price, stock
		FROM products
		`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []model.Product

	for rows.Next() {

		var product model.Product

		rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Stock,
		)

		products = append(products, product)
	}

	return products, nil
}

func GetProductByID(id string) (*model.Product, error) {

	var product model.Product

	err := database.DB.QueryRow(
		`
		SELECT id, name, description, price, stock
		FROM products
		WHERE id=$1
		`,
		id,
	).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Stock,
	)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

