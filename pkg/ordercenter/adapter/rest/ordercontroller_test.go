package rest

import (
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"

	"pjhu/medicine/pkg/ordercenter/application"
	"pjhu/medicine/pkg/ordercenter/domain"
	"pjhu/medicine/pkg/ordercenter/mock"
)

func TestOrderController_placeOrder(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var rst = application.PlaceOrderResponse{
		ID: 1457312267728064512,
	}

	m := mock.NewMockIApplicationService(ctrl)
	m.
		EXPECT().
		PlaceOrderHandler(gomock.Any()).
		Return(rst, nil).
		Times(1)

	router := gin.Default()
	oc := Build(m)
	oc.InitRouters(router)
	apitest.New().
		Handler(router).
		Post("/api/v1/customer/orders").
		JSON(`{
			"productId": 1,
			"sku": "1",
			"quantity": 1,
			"address": "长宁区"
		}`).
		Expect(t).
		Assert(jsonpath.Equal(`$.data.id`, float64(1457312267728064512))).
		Status(http.StatusCreated).
		End()
}

func TestOrderController_getOrderDetail(t *testing.T) {

	gin.SetMode(gin.TestMode)
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	var newOrder = domain.UserOrder{
		Id:               1457312267728064512,
		OrderAmountTotal: 1,
		PayChannel:       "alipay",
		OrderStatus:      "OrderStatus",
		CreatedAt:        time.Now(),
		CreatedBy:        "CreatedBy",
		LastModifiedAt:   time.Now(),
		LastModifiedBy:   "LastModifiedBy",
	}
	m := mock.NewMockIApplicationService(ctl)
	m.
		EXPECT().
		GetOrderDetail(gomock.Any()).
		Return(newOrder, nil).
		Times(1)

	router := gin.Default()
	oc := Build(m)
	oc.InitRouters(router)
	apitest.New().
		Handler(router).
		Get("/api/v1/customer/orders/1457312267728064512").
		Expect(t).
		Assert(jsonpath.Equal(`$.data.Id`, float64(1457312267728064512))).
		Assert(jsonpath.Equal(`$.data.PayChannel`, "alipay")).
		Assert(jsonpath.Equal(`$.data.OrderStatus`, "OrderStatus")).
		Assert(jsonpath.Equal(`$.data.CreatedBy`, "CreatedBy")).
		Assert(jsonpath.Equal(`$.data.LastModifiedBy`, "LastModifiedBy")).
		Status(http.StatusOK).
		End()
}
