package domain

type IAuthRepository interface {
	InsertOne(member *Member) error
	FindBy(member *Member) error
}
