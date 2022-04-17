package domain

type IRepository interface {
	InsertOne(userOrder *UserOrder) (int64, error)
	Get(userOrder *UserOrder) (*UserOrder, error)
}
