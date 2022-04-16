package application

// PlaceOrderResponse respionse body for create order
type PlaceOrderResponse struct {
	ID int64 `json:"id"`
}

type OrderResponse struct {
	// ID int64 `json:"id"`
	// OrderAmountTotal int
	// PayChannel string `xorm:"varchar(32)"`
	// OrderStatus string `xorm:"varchar(25) not null"`
	// CreatedAt time.Time `xorm:"created"`
	// CreatedBy string `xorm:"not null"`
	// LastModifiedAt time.Time `xorm:"updated"`
	// LastModifiedBy string `xorm:"not null"`
}
