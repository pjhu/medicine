package domain

type IRepository interface {
	InsertOne(member *Member) (int64, error)
	Get(member *Member) (*Member, error)
	Exist(member *Member) (bool, error)
}
