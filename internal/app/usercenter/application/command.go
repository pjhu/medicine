package application

// SigninCommand for signup request
type SigninCommand struct {
	Phone    string `json:"phone" binding:"required"`
	Code     string `json:"code" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ValidateTokenCommand for signup request
type ValidateTokenCommand struct {
	Token string `json:"token" binding:"required"`
}
