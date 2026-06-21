package futures_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/extension/interceptor"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/api"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/common/logger"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/generate/futures/order"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/types"
)

var orderApi order.OrderAPI

func init() {
	key := os.Getenv("API_KEY")
	secret := os.Getenv("API_SECRET")
	passphrase := os.Getenv("API_PASSPHRASE")

	defaultLogger := logger.NewDefaultLogger()
	logger.SetLogger(defaultLogger)

	httpOptionBuilder := types.NewTransportOptionBuilder().
		SetKeepAlive(true).
		AddInterceptors(interceptor.NewLoggingInterceptor(false, defaultLogger))

	option := types.NewClientOptionBuilder().
		WithKey(key).
		WithSecret(secret).
		WithPassphrase(passphrase).
		WithSpotEndpoint(types.GlobalApiEndpoint).
		WithFuturesEndpoint(types.GlobalFuturesApiEndpoint).
		WithBrokerEndpoint(types.GlobalBrokerApiEndpoint).
		WithTransportOption(httpOptionBuilder.Build()).
		Build()

	client := api.NewClient(option)

	orderApi = client.RestService().GetFuturesService().GetOrderAPI()
}

func TestFillsGetTradeHistoryReq(t *testing.T) {
	// GetTradeHistory
	// Get Trade History
	// /api/v1/fills

	builder := order.NewGetTradeHistoryReqBuilder()
	builder.SetSide("buy").SetType("market")
	req := builder.Build()

	resp, err := orderApi.GetTradeHistory(req, context.TODO())
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(resp.ToMap())
	if err != nil {
		panic(err)
	}
	fmt.Println("code:", resp.CommonResponse.Code)
	fmt.Println("message:", resp.CommonResponse.Message)
	fmt.Println("data:", string(data))
}

func TestFillsGetOpenOrderValueReq(t *testing.T) {
	// GetOpenOrderValue
	// Get Open Order Value
	// /api/v1/openOrderStatistics

	builder := order.NewGetOpenOrderValueReqBuilder()
	builder.SetSymbol("DOGEUSDTM")
	req := builder.Build()

	resp, err := orderApi.GetOpenOrderValue(req, context.TODO())
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(resp.ToMap())
	if err != nil {
		panic(err)
	}
	fmt.Println("code:", resp.CommonResponse.Code)
	fmt.Println("message:", resp.CommonResponse.Message)
	fmt.Println("data:", string(data))
}

func TestOrderGetOrderByClientOidReq(t *testing.T) {
	// GetOrderByClientOid
	// Get Order By ClientOid
	// /api/v1/orders/byClientOid

	builder := order.NewGetOrderByClientOidReqBuilder()
	builder.SetClientOid("df2e52bb-1d11-44d9-91a9-26899f3f46cb")
	req := builder.Build()

	resp, err := orderApi.GetOrderByClientOid(req, context.TODO())
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(resp.ToMap())
	if err != nil {
		panic(err)
	}
	fmt.Println("code:", resp.CommonResponse.Code)
	fmt.Println("message:", resp.CommonResponse.Message)
	fmt.Println("data:", string(data))
}

func TestOrderCancelOrderByClientOidReq(t *testing.T) {
	// CancelOrderByClientOid
	// Cancel Order By ClientOid
	// /api/v1/orders/client-order/{clientOid}

	builder := order.NewCancelOrderByClientOidReqBuilder()
	builder.SetSymbol("XBTUSDTM").SetClientOid("e59071a0-382e-45f0-b934-6c0cea16908b")
	req := builder.Build()

	resp, err := orderApi.CancelOrderByClientOid(req, context.TODO())
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(resp.ToMap())
	if err != nil {
		panic(err)
	}
	fmt.Println("code:", resp.CommonResponse.Code)
	fmt.Println("message:", resp.CommonResponse.Message)
	fmt.Println("data:", string(data))
}

func TestOrderCancelAllOrdersV1Req(t *testing.T) {
	// CancelAllOrders
	// Cancel All Orders
	// /api/v1/orders

	builder := order.NewCancelAllOrdersV1ReqBuilder()
	builder.SetSymbol("XBTUSDTM")
	req := builder.Build()

	resp, err := orderApi.CancelAllOrdersV1(req, context.TODO())
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(resp.ToMap())
	if err != nil {
		panic(err)
	}
	fmt.Println("code:", resp.CommonResponse.Code)
	fmt.Println("message:", resp.CommonResponse.Message)
	fmt.Println("data:", string(data))
}
func TestOrderGetOrderListReq(t *testing.T) {
	// GetOrderList
	// Get Order List
	// /api/v1/orders

	builder := order.NewGetOrderListReqBuilder()
	builder.SetStatus("active").SetSymbol("XBTUSDTM").SetSide("buy")
	req := builder.Build()

	resp, err := orderApi.GetOrderList(req, context.TODO())
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(resp.ToMap())
	if err != nil {
		panic(err)
	}
	fmt.Println("code:", resp.CommonResponse.Code)
	fmt.Println("message:", resp.CommonResponse.Message)
	fmt.Println("data:", string(data))
}

func TestOrderBatchCancelOrdersReq(t *testing.T) {
	// BatchCancelOrders
	// Batch Cancel Orders
	// /api/v1/orders/multi-cancel

	builder := order.NewBatchCancelOrdersReqBuilder()
	builder.SetOrderIdsList([]string{"250442580993564673", "250442581069062144"}).SetClientOidsList([]order.BatchCancelOrdersClientOidsList{
		*order.NewBatchCancelOrdersClientOidsList("XBTUSDTM", "4d7571f4-8a26-43a0-9fae-a49161682d4a"),
		*order.NewBatchCancelOrdersClientOidsList("XBTUSDTM", "8ae42301-86d5-4df5-8c43-86561d7e8b13"),
	})
	req := builder.Build()

	resp, err := orderApi.BatchCancelOrders(req, context.TODO())
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(resp.ToMap())
	if err != nil {
		panic(err)
	}
	fmt.Println("code:", resp.CommonResponse.Code)
	fmt.Println("message:", resp.CommonResponse.Message)
	fmt.Println("data:", string(data))
}
func TestOrderBatchAddOrdersReq(t *testing.T) {
	// BatchAddOrders
	// Batch Add Orders
	// /api/v1/orders/multi

	order1 := order.NewBatchAddOrdersItemBuilder().SetClientOid(uuid.NewString()).SetSide("buy").SetSymbol("XBTUSDTM").
		SetLeverage(3.0).SetType("limit").SetRemark("order_test").SetMarginMode("CROSS").
		SetPrice("10000").SetSize(1).Build()
	order2 := order.NewBatchAddOrdersItemBuilder().SetClientOid(uuid.NewString()).SetSide("buy").SetSymbol("XBTUSDTM").
		SetLeverage(3.0).SetType("limit").SetRemark("order_test").SetMarginMode("CROSS").
		SetPrice("10000").SetSize(1).Build()

	builder := order.NewBatchAddOrdersReqBuilder()
	builder.SetItems([]order.BatchAddOrdersItem{*order1, *order2})
	req := builder.Build()

	resp, err := orderApi.BatchAddOrders(req, context.TODO())
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(resp.ToMap())
	if err != nil {
		panic(err)
	}
	fmt.Println("code:", resp.CommonResponse.Code)
	fmt.Println("message:", resp.CommonResponse.Message)
	fmt.Println("data:", string(data))
}

func TestOrderCancelOrderByIdReq(t *testing.T) {
	// CancelOrderById
	// Cancel Order By OrderId
	// /api/v1/orders/{orderId}

	builder := order.NewCancelOrderByIdReqBuilder()
	builder.SetOrderId("250442004150878209")
	req := builder.Build()

	resp, err := orderApi.CancelOrderById(req, context.TODO())
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(resp.ToMap())
	if err != nil {
		panic(err)
	}
	fmt.Println("code:", resp.CommonResponse.Code)
	fmt.Println("message:", resp.CommonResponse.Message)
	fmt.Println("data:", string(data))
}
func TestOrderGetOrderByOrderIdReq(t *testing.T) {
	// GetOrderByOrderId
	// Get Order By OrderId
	// /api/v1/orders/{order-id}

	builder := order.NewGetOrderByOrderIdReqBuilder()
	builder.SetOrderId("250441187108892672")
	req := builder.Build()

	resp, err := orderApi.GetOrderByOrderId(req, context.TODO())
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(resp.ToMap())
	if err != nil {
		panic(err)
	}
	fmt.Println("code:", resp.CommonResponse.Code)
	fmt.Println("message:", resp.CommonResponse.Message)
	fmt.Println("data:", string(data))
}
func TestOrderAddOrderReq(t *testing.T) {
	// AddOrder
	// Add Order
	// /api/v1/orders

	builder := order.NewAddOrderReqBuilder()
	builder.SetClientOid(uuid.NewString()).SetSide("buy").SetSymbol("XBTUSDTM").
		SetLeverage(3.0).SetType("limit").SetRemark("order_test").SetMarginMode("CROSS").
		SetPrice("10000").SetSize(1)
	req := builder.Build()

	resp, err := orderApi.AddOrder(req, context.TODO())
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(resp.ToMap())
	if err != nil {
		panic(err)
	}
	fmt.Println("code:", resp.CommonResponse.Code)
	fmt.Println("message:", resp.CommonResponse.Message)
	fmt.Println("data:", string(data))
}
func TestOrderAddOrderTestReq(t *testing.T) {
	// AddOrderTest
	// Add Order Test
	// /api/v1/orders/test

	builder := order.NewAddOrderTestReqBuilder()
	builder.SetClientOid(uuid.NewString()).SetSide("buy").SetSymbol("XBTUSDTM").
		SetLeverage(3.0).SetType("limit").SetRemark("order_test").SetMarginMode("ISOLATED").
		SetPrice("10000").SetSize(1)
	req := builder.Build()

	resp, err := orderApi.AddOrderTest(req, context.TODO())
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(resp.ToMap())
	if err != nil {
		panic(err)
	}
	fmt.Println("code:", resp.CommonResponse.Code)
	fmt.Println("message:", resp.CommonResponse.Message)
	fmt.Println("data:", string(data))
}

func TestOrderGetRecentClosedOrdersReq(t *testing.T) {
	// GetRecentClosedOrders
	// Get Recent Closed Orders
	// /api/v1/recentDoneOrders

	builder := order.NewGetRecentClosedOrdersReqBuilder()
	builder.SetSymbol("XBTUSDTM")
	req := builder.Build()

	resp, err := orderApi.GetRecentClosedOrders(req, context.TODO())
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(resp.ToMap())
	if err != nil {
		panic(err)
	}
	fmt.Println("code:", resp.CommonResponse.Code)
	fmt.Println("message:", resp.CommonResponse.Message)
	fmt.Println("data:", string(data))
}

func TestFillsGetRecentTradeHistoryReq(t *testing.T) {
	// GetRecentTradeHistory
	// Get Recent Trade History
	// /api/v1/recentFills

	builder := order.NewGetRecentTradeHistoryReqBuilder()
	builder.SetSymbol("DOGEUSDTM")
	req := builder.Build()

	resp, err := orderApi.GetRecentTradeHistory(req, context.TODO())
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(resp.ToMap())
	if err != nil {
		panic(err)
	}
	fmt.Println("code:", resp.CommonResponse.Code)
	fmt.Println("message:", resp.CommonResponse.Message)
	fmt.Println("data:", string(data))
}

func TestOrderAddTPSLOrderReq(t *testing.T) {
	// AddTPSLOrder
	// Add Take Profit And Stop Loss Order
	// /api/v1/st-orders

	builder := order.NewAddTPSLOrderReqBuilder()
	builder.SetClientOid(uuid.NewString()).SetSide("buy").SetSymbol("XBTUSDTM").SetMarginMode("CROSS").
		SetLeverage(1.0).SetType("limit").SetRemark("test_order").SetStopPriceType("TP").
		SetPrice("10000").SetSize(1).SetTriggerStopUpPrice("8000").SetTriggerStopDownPrice("12000")
	req := builder.Build()

	resp, err := orderApi.AddTPSLOrder(req, context.TODO())
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(resp.ToMap())
	if err != nil {
		panic(err)
	}
	fmt.Println("code:", resp.CommonResponse.Code)
	fmt.Println("message:", resp.CommonResponse.Message)
	fmt.Println("data:", string(data))
}

func TestOrderCancelAllStopOrdersReq(t *testing.T) {
	// CancelAllStopOrders
	// Cancel All Stop orders
	// /api/v1/stopOrders

	builder := order.NewCancelAllStopOrdersReqBuilder()
	builder.SetSymbol("XBTUSDTM")
	req := builder.Build()

	resp, err := orderApi.CancelAllStopOrders(req, context.TODO())
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(resp.ToMap())
	if err != nil {
		panic(err)
	}
	fmt.Println("code:", resp.CommonResponse.Code)
	fmt.Println("message:", resp.CommonResponse.Message)
	fmt.Println("data:", string(data))
}

func TestOrderGetStopOrderListReq(t *testing.T) {
	// GetStopOrderList
	// Get Stop Order List
	// /api/v1/stopOrders

	builder := order.NewGetStopOrderListReqBuilder()
	builder.SetSymbol("XBTUSDTM")
	req := builder.Build()

	resp, err := orderApi.GetStopOrderList(req, context.TODO())
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(resp.ToMap())
	if err != nil {
		panic(err)
	}
	fmt.Println("code:", resp.CommonResponse.Code)
	fmt.Println("message:", resp.CommonResponse.Message)
	fmt.Println("data:", string(data))
}

func TestOrderCancelAllOrdersV3Req(t *testing.T) {
	// CancelAllOrders
	// Cancel All Orders
	// /api/v1/orders

	builder := order.NewCancelAllOrdersV3ReqBuilder()
	builder.SetSymbol("XBTUSDTM")
	req := builder.Build()

	resp, err := orderApi.CancelAllOrdersV3(req, context.TODO())
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(resp.ToMap())
	if err != nil {
		panic(err)
	}
	fmt.Println("code:", resp.CommonResponse.Code)
	fmt.Println("message:", resp.CommonResponse.Message)
	fmt.Println("data:", string(data))
}
