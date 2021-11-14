package cache

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

 const (
    // AuthorizationHeader for http header
    AuthorizationHeader = "Authorization"
    // MiniProgramTokenPrefix for mp token prefix
    MiniProgramTokenPrefix = "MP "
    tokenLength = 32
)

// CreateTokenKey for user
func CreateTokenKey() (token string){
    return MiniProgramTokenPrefix + strings.Replace(uuid.New().String(), "-", "", -1)
}

// ExtractTokenKey form http header
func ExtractTokenKey(fullTokenString string) (token string, err error){
	if !strings.HasPrefix(fullTokenString, MiniProgramTokenPrefix) {
		return token, errors.New("no correct token prefix")
	}
	token = strings.TrimPrefix(fullTokenString, MiniProgramTokenPrefix)
    if len(token) != tokenLength {
        return token, errors.New("not correct token")
    }
    return token, nil
}