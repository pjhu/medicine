package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/vulcand/oxy/forward"
	"github.com/vulcand/oxy/testutils"
	log "github.com/sirupsen/logrus"
)

// ReverseProxy returns
func ReverseProxy(target string, path string) gin.HandlerFunc {
	fwd, _ := forward.New()
	return func(ctx *gin.Context) {
		req := ctx.Request
		log.Info(req)
		req.Header = ctx.Request.Header
		req.URL = testutils.ParseURI(target + ctx.Param(path))
		fwd.ServeHTTP(ctx.Writer, ctx.Request)
	}
}