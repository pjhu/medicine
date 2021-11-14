package command

// ValidateTokenCommand for signup request
type ValidateTokenCommand struct {
	Token string `json:"token" binding:"required"`
}