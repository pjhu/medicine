package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"bff/common/main/lib"
	"bff/common/main/errors"
)

type tokenCommond struct {
	Token string
}

// UserAuth validate token, get user meta
func UserAuth() gin.HandlerFunc {
	var userMeta UserMeta
	return func(ctx *gin.Context) {
		fullTokenString := ctx.GetHeader("Authorization")
		resp, err := lib.Client.R().
			SetBody(tokenCommond{Token: fullTokenString}).
			SetResult(&userMeta).
			Post(viper.GetString("microservice.usercenter") + "/internal-api/v1/varify-token")
		if resp.StatusCode() != http.StatusOK|| err != nil {
			log.Info(resp,err)
			ctx.Header("Content-Type", "application/json")
			ctx.AbortWithError(http.StatusForbidden, errors.NewErrorWithCode(errors.SystemInternalError, "no auth"))
			return
		}
		ctx.Set(AuthUserKey, userMeta)
		ctx.Next()
	}
}