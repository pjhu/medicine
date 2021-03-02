package identitydomain

import "time"

// Member for identity domain model
type Member struct {
	Id int64 `xorm:"pk"`
	Phone string `xorm:"varchar(32) not null"`
	Nickname string
	Password string
	CreatedAt time.Time `xorm:"created"`
	LastModifiedAt time.Time `xorm:"updated"`
}

// NewMember for user register
func NewMember(id int64, phone string) *Member {
	return &Member{Id:id, Phone: phone,}
}