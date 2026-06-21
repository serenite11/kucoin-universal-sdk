package margin_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/extension/interceptor"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/api"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/common/logger"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/generate/margin/credit"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/types"
)

var creditApi credit.CreditAPI

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

	creditApi = client.RestService().GetMarginService().GetCreditAPI()
}

func TestCreditModifyPurchaseReq(t *testing.T) {
	// ModifyPurchase
	// Modify Purchase
	// /api/v3/lend/purchase/update

	builder := credit.NewModifyPurchaseReqBuilder()
	builder.SetCurrency("DOGE").SetInterestRate("0.02").SetPurchaseOrderNo("6745708dad11d500073f083f")
	req := builder.Build()

	resp, err := creditApi.ModifyPurchase(req, context.TODO())
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

func TestCreditGetLoanMarketReq(t *testing.T) {
	// GetLoanMarket
	// Get Loan Market
	// /api/v3/project/list

	builder := credit.NewGetLoanMarketReqBuilder()
	builder.SetCurrency("DOGE")
	req := builder.Build()

	resp, err := creditApi.GetLoanMarket(req, context.TODO())
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

func TestCreditGetLoanMarketInterestRateReq(t *testing.T) {
	// GetLoanMarketInterestRate
	// Get Loan Market Interest Rate
	// /api/v3/project/marketInterestRate

	builder := credit.NewGetLoanMarketInterestRateReqBuilder()
	builder.SetCurrency("DOGE")
	req := builder.Build()

	resp, err := creditApi.GetLoanMarketInterestRate(req, context.TODO())
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

func TestCreditGetPurchaseOrdersReq(t *testing.T) {
	// GetPurchaseOrders
	// Get Purchase Orders
	// /api/v3/purchase/orders

	builder := credit.NewGetPurchaseOrdersReqBuilder()
	builder.SetCurrency("DOGE").SetPurchaseOrderNo("6745708dad11d500073f083f").SetStatus("PENDING")
	req := builder.Build()

	resp, err := creditApi.GetPurchaseOrders(req, context.TODO())
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

func TestCreditPurchaseReq(t *testing.T) {
	// Purchase
	// Purchase
	// /api/v3/purchase

	builder := credit.NewPurchaseReqBuilder()
	builder.SetCurrency("DOGE").SetSize("10").SetInterestRate("0.01")
	req := builder.Build()

	resp, err := creditApi.Purchase(req, context.TODO())
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

func TestCreditGetRedeemOrdersReq(t *testing.T) {
	// GetRedeemOrders
	// Get Redeem Orders
	// /api/v3/redeem/orders

	builder := credit.NewGetRedeemOrdersReqBuilder()
	builder.SetCurrency("DOGE").SetStatus("DONE")
	req := builder.Build()

	resp, err := creditApi.GetRedeemOrders(req, context.TODO())
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

func TestCreditRedeemReq(t *testing.T) {
	// Redeem
	// Redeem
	// /api/v3/redeem

	builder := credit.NewRedeemReqBuilder()
	builder.SetCurrency("DOGE").SetSize("10").SetPurchaseOrderNo("6745708dad11d500073f083f")
	req := builder.Build()

	resp, err := creditApi.Redeem(req, context.TODO())
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
