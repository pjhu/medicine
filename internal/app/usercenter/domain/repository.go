package domain

type IRepository interface {
	InsertOne(member *Member) error
	FindBy(member *Member) error
}
