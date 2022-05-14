package domain

type IOrderRepository interface {
	InsertOne(userOrder *UserOrder) error
	FindBy(userOrder *UserOrder) error
}
