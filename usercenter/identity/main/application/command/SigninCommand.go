package identitycommand

// SignoutCommand for signup request
type SignoutCommand struct {
	Phone string `json:"phone" binding:"required"`
	Code string `json:"code" binding:"required"`
	Password string `json:"password" binding:"required"`
}