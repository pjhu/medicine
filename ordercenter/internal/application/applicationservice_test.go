package service

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/golang/mock/gomock"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"

	"ordercenter/internal/application/command"
	"ordercenter/mock"
)

func TestOrderApplicationService_PlaceOrderHandler(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock.NewMockIRepository(ctrl)
	m.
		EXPECT().
		InsertOne(gomock.Any()).
		Return(int64(1), nil).
		Times(1)

	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	httpmock.RegisterResponder("POST", "http://localhost:48080/api/v1/accounts/decrease",
		httpmock.NewStringResponder(200, `{"orderId": "1", "orderAmount": 1}`))

	orderApplicationService := Build(m, client)
	newCommand := command.PlaceOrderCommand {
		ProductID: 1,
		Sku: "1",
		Quantity: 1,
		Address: "长宁区",
	}
	result, _ := orderApplicationService.PlaceOrderHandler(newCommand)
	assert.Less(t, int64(0), result.ID)
}
