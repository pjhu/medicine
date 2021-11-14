package command

// SigninCommand for signup request
type SigninCommand struct {
	Phone string `json:"phone" binding:"required"`
	Code string `json:"code" binding:"required"`
	Password string `json:"password" binding:"required"`
}