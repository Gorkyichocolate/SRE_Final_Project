package model

type Order struct {
	ID         int64   `json:"id"`
	UserID     int64   `json:"user_id"`
	ProductID  int64   `json:"product_id"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}

type CreateOrderRequest struct {
	ProductID int64 `json:"product_id"`
	Quantity  int   `json:"quantity"`
}
