package account_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/extension/interceptor"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/api"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/common/logger"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/generate/account/fee"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/types"
)

var feeApi fee.FeeAPI

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

	feeApi = client.RestService().GetAccountService().GetFeeAPI()
}

func TestFeeGetBasicFeeReq(t *testing.T) {
	// GetBasicFee
	// Get Basic Fee - Spot/Margin
	// /api/v1/base-fee

	builder := fee.NewGetBasicFeeReqBuilder()
	builder.SetCurrencyType(0)
	req := builder.Build()

	resp, err := feeApi.GetBasicFee(req, context.TODO())
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

func TestFeeGetSpotActualFeeReq(t *testing.T) {
	// GetSpotActualFee
	// Get Actual Fee - Spot/Margin
	// /api/v1/trade-fees

	builder := fee.NewGetSpotActualFeeReqBuilder()
	builder.SetSymbols("BTC-USDT")
	req := builder.Build()

	resp, err := feeApi.GetSpotActualFee(req, context.TODO())
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

func TestFeeGetFuturesActualFeeReq(t *testing.T) {
	// GetFuturesActualFee
	// Get Actual Fee - Futures
	// /api/v1/trade-fees

	builder := fee.NewGetFuturesActualFeeReqBuilder()
	builder.SetSymbol("XBTUSDM")
	req := builder.Build()

	resp, err := feeApi.GetFuturesActualFee(req, context.TODO())
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
