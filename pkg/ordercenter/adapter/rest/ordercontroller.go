package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"pjhu/medicine/pkg/ordercenter/application"
)

type IOrderController interface {
	InitRouters(router *gin.Engine)
}

type OrderController struct {
	appSvc application.IApplicationService
}

func Build(app application.IApplicationService) IOrderController {
	return OrderController{
		appSvc: app,
	}
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// InitRouters for order
func (oc OrderController) InitRouters(router *gin.Engine) {

	//customerGroup := router.Group("/api/v1/customer")
	//customerGroup.Use(middleware.UserAuth())
	r := router.Group("/api/v1/customer/")
	r.POST("orders", oc.placeOrder)
	r.GET("orders/:id", oc.getOrderDetail)
}

func (oc OrderController) placeOrder(ctx *gin.Context) {
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

	placeOrderResponse, err := oc.appSvc.PlaceOrderHandler(placeOrderCommand)
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

func (oc OrderController) getOrderDetail(ctx *gin.Context) {
	var req getAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	order, err := oc.appSvc.GetOrderDetail(req.ID)
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
