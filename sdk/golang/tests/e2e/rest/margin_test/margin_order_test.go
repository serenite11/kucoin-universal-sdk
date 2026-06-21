package margin

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
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/generate/margin/order"
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

	orderApi = client.RestService().GetMarginService().GetOrderAPI()
}

func TestOrderAddOrderV1Req(t *testing.T) {
	// AddOrderV1
	// Add Order(V1)
	// /api/v1/margin/order

	builder := order.NewAddOrderV1ReqBuilder()
	builder.SetClientOid(uuid.NewString()).SetSide("buy").SetSymbol("BTC-USDT").
		SetType("limit").SetPrice("10000").SetSize("0.001").SetAutoBorrow(false).
		SetAutoRepay(false).SetMarginModel("isolated")
	req := builder.Build()

	resp, err := orderApi.AddOrderV1(req, context.TODO())
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

// TODO rate limit
func TestOrderAddOrderTestV1Req(t *testing.T) {
	// AddOrderTestV1
	// Add Order Test(V1)
	// /api/v1/margin/order/test

	builder := order.NewAddOrderTestV1ReqBuilder()
	builder.SetClientOid(uuid.NewString()).SetSide("buy").SetSymbol("BTC-USDT").
		SetType("limit").SetPrice("10000").SetSize("0.001").SetAutoBorrow(true).
		SetAutoRepay(true).SetMarginModel("isolated")
	req := builder.Build()

	resp, err := orderApi.AddOrderTestV1(req, context.TODO())
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

func TestOrderGetTradeHistoryReq(t *testing.T) {
	// GetTradeHistory
	// Get Trade History
	// /api/v3/hf/margin/fills

	builder := order.NewGetTradeHistoryReqBuilder()
	builder.SetSymbol("DOGE-USDT").SetTradeType("MARGIN_TRADE").SetSide("buy")
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

func TestOrderGetSymbolsWithOpenOrderReq(t *testing.T) {
	// GetSymbolsWithOpenOrder
	// Get Symbols With Open Order
	// /api/v3/hf/margin/order/active/symbols

	builder := order.NewGetSymbolsWithOpenOrderReqBuilder()
	builder.SetTradeType("MARGIN_ISOLATED_TRADE")
	req := builder.Build()

	resp, err := orderApi.GetSymbolsWithOpenOrder(req, context.TODO())
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
	// /api/v3/hf/margin/order

	builder := order.NewAddOrderReqBuilder()
	builder.SetClientOid(uuid.NewString()).SetSide("buy").SetSymbol("DOGE-USDT").
		SetType("market").SetFunds("10").SetAutoBorrow(true).
		SetAutoRepay(true).SetIsIsolated(false)
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

// TODO rate limit
func TestOrderAddOrderTestReq(t *testing.T) {
	// AddOrderTest
	// Add Order Test
	// /api/v3/hf/margin/order/test

	builder := order.NewAddOrderTestReqBuilder()
	builder.SetClientOid(uuid.NewString()).SetSide("buy").SetSymbol("BTC-USDT").
		SetType("limit").SetPrice("10000").SetSize("0.004").SetAutoBorrow(true).
		SetAutoRepay(true).SetIsIsolated(true)
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
func TestOrderGetOpenOrdersReq(t *testing.T) {
	// GetOpenOrders
	// Get Open Orders
	// /api/v3/hf/margin/orders/active

	builder := order.NewGetOpenOrdersReqBuilder()
	builder.SetSymbol("BTC-USDT").SetTradeType("MARGIN_ISOLATED_TRADE")
	req := builder.Build()

	resp, err := orderApi.GetOpenOrders(req, context.TODO())
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
	// /api/v3/hf/margin/orders/client-order/{clientOid}

	builder := order.NewCancelOrderByClientOidReqBuilder()
	builder.SetClientOid("be4c9f3e-d070-437d-9f97-48854fa05c28").SetSymbol("BTC-USDT")
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
func TestOrderGetOrderByClientOidReq(t *testing.T) {
	// GetOrderByClientOid
	// Get Order By ClientOid
	// /api/v3/hf/margin/orders/client-order/{clientOid}

	builder := order.NewGetOrderByClientOidReqBuilder()
	builder.SetSymbol("BTC-USDT").SetClientOid("e865a276-0253-42b7-8c2f-a6764760b199")
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

func TestOrderCancelAllOrdersBySymbolReq(t *testing.T) {
	// CancelAllOrdersBySymbol
	// Cancel All Orders By Symbol
	// /api/v3/hf/margin/orders

	builder := order.NewCancelAllOrdersBySymbolReqBuilder()
	builder.SetSymbol("BTC-USDT").SetTradeType("MARGIN_ISOLATED_TRADE")
	req := builder.Build()

	resp, err := orderApi.CancelAllOrdersBySymbol(req, context.TODO())
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

func TestOrderGetClosedOrdersReq(t *testing.T) {
	// GetClosedOrders
	// Get Closed Orders
	// /api/v3/hf/margin/orders/done

	builder := order.NewGetClosedOrdersReqBuilder()
	builder.SetSymbol("BTC-USDT").SetTradeType("MARGIN_ISOLATED_TRADE").SetSide("buy").SetType("limit")
	req := builder.Build()

	resp, err := orderApi.GetClosedOrders(req, context.TODO())
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

func TestOrderCancelOrderByOrderIdReq(t *testing.T) {
	// CancelOrderByOrderId
	// Cancel Order By OrderId
	// /api/v3/hf/margin/orders/{orderId}

	builder := order.NewCancelOrderByOrderIdReqBuilder()
	builder.SetOrderId("67441ae54fe7460007bcf007").SetSymbol("BTC-USDT")
	req := builder.Build()

	resp, err := orderApi.CancelOrderByOrderId(req, context.TODO())
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
	// /api/v3/hf/margin/orders/{orderId}

	builder := order.NewGetOrderByOrderIdReqBuilder()
	builder.SetSymbol("BTC-USDT").SetOrderId("67441ae54fe7460007bcf007")
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
