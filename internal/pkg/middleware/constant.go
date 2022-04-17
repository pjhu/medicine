package middleware

// UserMeta store in token
type UserMeta struct {
	Id       int64
	Phone    string
	Nickname string
}

// AuthUserKey is the cookie name for user credential in basic auth.
const AuthUserKey = "user"
