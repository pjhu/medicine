package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/pjhu/medicine/internal/app/ordercenter/adapter/persistence"
	"github.com/pjhu/medicine/internal/app/ordercenter/application"
	"github.com/pjhu/medicine/internal/pkg/cache"
	"github.com/pjhu/medicine/internal/pkg/datasource"
	"github.com/pjhu/medicine/pkg/errors"
	"github.com/pjhu/medicine/pkg/httpclient"
)

// InitRouters for order
func InitRouters(router *gin.Engine) {

	r := router.Group("/api/v1/customer/")
	r.POST("orders", placeOrder)
	r.GET("orders/:id", getOrderDetail)
}

func placeOrder(ctx *gin.Context) {
	var placeOrderCommand application.PlaceOrderCommand
	if err := ctx.ShouldBind(&placeOrderCommand); err != nil {
		logrus.Error(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			errors.NewErrorWithCode(errors.InvalidParameter, err.Error()))
		return
	}

	logrus.Info("controller info: ", placeOrderCommand)
	appSvc := buildApplicationService()
	placeOrderResponse, err := appSvc.PlaceOrderHandler(placeOrderCommand)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,
			errors.NewErrorWithCode(errors.SystemInternalError, err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"data": placeOrderResponse,
	})
}

func getOrderDetail(ctx *gin.Context) {
	var req application.GetAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			errors.NewErrorWithCode(errors.InvalidParameter, err.Error()))
		return
	}

	appSvc := buildApplicationService()
	order, err := appSvc.GetOrderDetail(req.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,
			errors.NewErrorWithCode(errors.SystemInternalError, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": order,
	})
}

func buildApplicationService() *application.OrderApplicationService {
	db := datasource.NewDBSession()
	repo := persistence.NewOrderRepository(db)
	return application.NewOrderApplicationService(db, repo, cache.Repository(), httpclient.Request())
}
