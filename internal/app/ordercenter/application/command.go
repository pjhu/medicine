package application

type GetAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// PlaceOrderCommand map request body for create order
type PlaceOrderCommand struct {
	ProductID int64  `json:"productId" binding:"required"`
	Sku       string `json:"sku" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required"`
	Address   string `json:"address" binding:"required"`
}
