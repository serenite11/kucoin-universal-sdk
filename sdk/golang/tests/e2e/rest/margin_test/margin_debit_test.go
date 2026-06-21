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
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/generate/margin/debit"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/types"
)

var debitApi debit.DebitAPI

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

	debitApi = client.RestService().GetMarginService().GetDebitAPI()
}

func TestDebitGetBorrowHistoryReq(t *testing.T) {
	// GetBorrowHistory
	// Get Borrow History
	// /api/v3/margin/borrow

	builder := debit.NewGetBorrowHistoryReqBuilder()
	builder.SetCurrency("USDT").SetIsIsolated(true).SetSymbol("BTC-USDT")
	req := builder.Build()

	resp, err := debitApi.GetBorrowHistory(req, context.TODO())
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

func TestDebitBorrowReq(t *testing.T) {
	// Borrow
	// Borrow
	// /api/v3/margin/borrow

	builder := debit.NewBorrowReqBuilder()
	builder.SetCurrency("USDT").SetSize(10.0).SetTimeInForce("IOC").
		SetSymbol("BTC-USDT").SetIsIsolated(true).SetIsHf(true)
	req := builder.Build()

	resp, err := debitApi.Borrow(req, context.TODO())
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

func TestDebitGetInterestHistoryReq(t *testing.T) {
	// GetInterestHistory
	// Get Interest History
	// /api/v3/margin/interest

	builder := debit.NewGetInterestHistoryReqBuilder()
	builder.SetSymbol("BTC-USDT").SetIsIsolated(true)
	req := builder.Build()

	resp, err := debitApi.GetInterestHistory(req, context.TODO())
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

func TestDebitGetRepayHistoryReq(t *testing.T) {
	// GetRepayHistory
	// Get Repay History
	// /api/v3/margin/repay

	builder := debit.NewGetRepayHistoryReqBuilder()
	builder.SetCurrency("USDT").SetIsIsolated(true).SetSymbol("BTC-USDT")
	req := builder.Build()

	resp, err := debitApi.GetRepayHistory(req, context.TODO())
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

func TestDebitRepayReq(t *testing.T) {
	// Repay
	// Repay
	// /api/v3/margin/repay

	builder := debit.NewRepayReqBuilder()
	builder.SetCurrency("USDT").SetSize(10.0).SetSymbol("BTC-USDT").SetIsIsolated(true).SetIsHf(true)
	req := builder.Build()

	resp, err := debitApi.Repay(req, context.TODO())
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

func TestDebitModifyLeverageReq(t *testing.T) {
	// ModifyLeverage
	// Modify Leverage
	// /api/v3/position/update-user-leverage

	builder := debit.NewModifyLeverageReqBuilder()
	builder.SetSymbol("BTC-USDT").SetIsIsolated(true).SetLeverage("3.1")
	req := builder.Build()

	resp, err := debitApi.ModifyLeverage(req, context.TODO())
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
