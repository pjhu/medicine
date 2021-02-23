package order

import (
	"net/http"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
	ordercommand "medicine/order/main/application/command"
	orderapplicationservice "medicine/order/main/application/services"
)

// Routers for order
func Routers(e *gin.Engine) {
	e.POST("/api/v1/admin/orders", func(c *gin.Context) {
		var placeOrderCommand ordercommand.PlaceOrderCommand
		if err := c.ShouldBind(&placeOrderCommand); err == nil {
			log.Info("controller info:%#v\n", placeOrderCommand)
			placeOrderResponse := orderapplicationservice.PlaceOrderHandler(placeOrderCommand)
			c.JSON(http.StatusOK, placeOrderResponse)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
}
