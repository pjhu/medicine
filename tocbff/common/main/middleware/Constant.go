package middleware

// UserMeta store in token
type UserMeta struct {
	Id int64
	Phone string
	Nickname string
}

const AuthUserKey = "user"