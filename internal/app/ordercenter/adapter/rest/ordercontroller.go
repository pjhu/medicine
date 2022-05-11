package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/pjhu/medicine/internal/app/ordercenter/application"
	"github.com/pjhu/medicine/internal/pkg/cache"
	"github.com/pjhu/medicine/internal/pkg/datasource"
	"github.com/pjhu/medicine/pkg/httpclient"
)

// InitRouters for order
func InitRouters(router *gin.Engine) {

	//customerGroup := router.Group("/api/v1/customer")
	//customerGroup.Use(middleware.UserAuth())
	r := router.Group("/api/v1/customer/")
	r.POST("orders", placeOrder)
	r.GET("orders/:id", getOrderDetail)
}

func placeOrder(ctx *gin.Context) {
	var placeOrderCommand application.PlaceOrderCommand
	if err := ctx.ShouldBind(&placeOrderCommand); err != nil {
		logrus.Error(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	logrus.Info("controller info: ", placeOrderCommand)

	appSvc := application.Builder(datasource.GetDB(), cache.Repository(), httpclient.Request())
	placeOrderResponse, err := appSvc.PlaceOrderHandler(placeOrderCommand)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"data":    nil,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"data":    placeOrderResponse,
		"message": nil,
	})
}

func getOrderDetail(ctx *gin.Context) {
	var req application.GetAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	appSvc := application.Builder(datasource.GetDB(), cache.Repository(), httpclient.Request())
	order, err := appSvc.GetOrderDetail(req.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"data":    nil,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    order,
		"message": nil,
	})
}
