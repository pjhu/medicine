package domain

type IRepository interface {
	InsertOne(userOrder *UserOrder) error
	FindBy(userOrder *UserOrder) error
}
