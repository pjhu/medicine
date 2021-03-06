package order

import (
	"net/http"
	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
	ordercommand "medicine/order/main/application/command"
	orderapplicationservice "medicine/order/main/application/services"
)

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// Routers for order
func Routers(e *gin.RouterGroup) {
	e.POST("orders", func(ctx *gin.Context) {
		var placeOrderCommand ordercommand.PlaceOrderCommand
		if err := ctx.ShouldBind(&placeOrderCommand); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		log.Info("controller info: ", placeOrderCommand)
		placeOrderResponse, err := orderapplicationservice.PlaceOrderHandler(placeOrderCommand)
		if err != nil {
			ctx.AbortWithError(err.GetHTTPStatus(), err)
			return
		}
		ctx.JSON(http.StatusCreated, placeOrderResponse)
	})

	e.GET("orders/:id", func(ctx *gin.Context) {
		var req getAccountRequest
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		order, err := orderapplicationservice.GetOrderDetail(req.ID)
		if err != nil {
			ctx.AbortWithError(err.GetHTTPStatus(), err)
			return
		}
		ctx.JSON(http.StatusCreated, order)
	})
}
