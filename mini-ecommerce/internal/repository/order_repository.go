package repository

import (
	"sre/internal/database"
	"sre/internal/model"
)

func CreateOrder(order model.Order) error {

	_, err := database.DB.Exec(
		`
		INSERT INTO orders(user_id, product_id, quantity, total_price)
		VALUES($1, $2, $3, $4)
		`,
		order.UserID,
		order.ProductID,
		order.Quantity,
		order.TotalPrice,
	)

	return err
}

func GetOrdersByUserID(userID int64) ([]model.Order, error) {

	rows, err := database.DB.Query(
		`
		SELECT id, user_id, product_id, quantity, total_price
		FROM orders
		WHERE user_id=$1
		`,
		userID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var orders []model.Order

	for rows.Next() {

		var order model.Order

		rows.Scan(
			&order.ID,
			&order.UserID,
			&order.ProductID,
			&order.Quantity,
			&order.TotalPrice,
		)

		orders = append(orders, order)
	}

	return orders, nil
}
