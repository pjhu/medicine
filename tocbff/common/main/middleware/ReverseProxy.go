package middleware

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/vulcand/oxy/forward"
	"github.com/vulcand/oxy/testutils"
)

var fwd, _ = forward.New()

// ReverseProxy returns
func ReverseProxy(target string, path string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, _ := ctx.Get(AuthUserKey)
		userMeta := user.(UserMeta)
		req := ctx.Request
		req.Header = ctx.Request.Header
		req.Header.Add("userId", strconv.FormatInt(userMeta.Id, 10))
		req.URL = testutils.ParseURI(target + ctx.Param(path))
		fwd.ServeHTTP(ctx.Writer, ctx.Request)
	}
}