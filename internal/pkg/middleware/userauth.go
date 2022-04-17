package middleware

import (
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"pjhu/medicine/pkg/httpclient"
)

type tokenCommand struct {
	Token string
}

func init() {
	hystrix.ConfigureCommand("user_auth", hystrix.CommandConfig{
		Timeout:                1000, // 执行command的超时时间。默认时间是1000毫秒
		MaxConcurrentRequests:  1,    // command的最大并发量 默认值是10
		RequestVolumeThreshold: 1,    // 一个统计窗口10秒内请求数量。达到这个请求数量后才去判断是否要开启熔断。默认值是20
		SleepWindow:            5000, // 当熔断器被打开后，SleepWindow的时间就是控制过多久后去尝试服务是否可用了。默认值是5000毫秒
		ErrorPercentThreshold:  1,    // 错误百分比，请求数量大于等于RequestVolumeThreshold并且错误率到达这个百分比后就会启动熔断 默认值是50
	})
}

// UserAuth returns a Basic HTTP Authorization middleware
// the key is the username and the value is the password.
func UserAuth() gin.HandlerFunc {
	var userMeta UserMeta
	return func(ctx *gin.Context) {
		fullTokenString := ctx.GetHeader("Authorization")
		var resp *resty.Response
		var err error
		hystrixErr := hystrix.Do("user_auth",
			func() error {
				resp, err = httpclient.BuildRestClient().R().
					SetBody(tokenCommand{Token: fullTokenString}).
					SetResult(&userMeta).
					Post(viper.GetString("microservice.usercenter") + "/internal-api/v1/varify-token")

				if resp.StatusCode() != http.StatusOK || err != nil {
					logrus.Info(resp, err)
					ctx.Header("Content-Type", "application/json")
					ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
						"data":    nil,
						"message": "no auth",
					})
					return nil
				}
				ctx.Set(AuthUserKey, userMeta)
				return nil
			},
			func(err error) error {
				logrus.Error("do this when services are down", err)
				ctx.Header("Content-Type", "application/json")
				ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
					"data":    nil,
					"message": "hystrix open",
				})
				return nil
			})
		if hystrixErr != nil {
			logrus.Error("hystrix error", err)
			ctx.Header("Content-Type", "application/json")
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"data":    nil,
				"message": "hystrix error",
			})
			return
		}
		ctx.Set(AuthUserKey, userMeta)
		ctx.Next()
	}
}
