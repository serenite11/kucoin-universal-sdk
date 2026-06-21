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
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/generate/margin/market"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/types"
)

var marketApi market.MarketAPI

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

	marketApi = client.RestService().GetMarginService().GetMarketAPI()
}

func TestMarketGetIsolatedMarginSymbolsReq(t *testing.T) {
	// GetIsolatedMarginSymbols
	// Get Symbols - Isolated Margin
	// /api/v1/isolated/symbols

	resp, err := marketApi.GetIsolatedMarginSymbols(context.TODO())
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

func TestMarketGetMarginConfigReq(t *testing.T) {
	// GetMarginConfig
	// Get Margin Config
	// /api/v1/margin/config

	resp, err := marketApi.GetMarginConfig(context.TODO())
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

func TestMarketGetMarkPriceDetailReq(t *testing.T) {
	// GetMarkPriceDetail
	// Get Mark Price Detail
	// /api/v1/mark-price/{symbol}/current

	builder := market.NewGetMarkPriceDetailReqBuilder()
	builder.SetSymbol("USDT-BTC")
	req := builder.Build()

	resp, err := marketApi.GetMarkPriceDetail(req, context.TODO())
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

func TestMarketGetETFInfoReq(t *testing.T) {
	// GetETFInfo
	// Get ETF Info
	// /api/v3/etf/info

	builder := market.NewGetETFInfoReqBuilder()
	builder.SetCurrency("BTCUP")
	req := builder.Build()

	resp, err := marketApi.GetETFInfo(req, context.TODO())
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

func TestMarketGetCrossMarginSymbolsReq(t *testing.T) {
	// GetCrossMarginSymbols
	// Get Symbols - Cross Margin
	// /api/v3/margin/symbols

	builder := market.NewGetCrossMarginSymbolsReqBuilder()
	builder.SetSymbol("BTC-USDT")
	req := builder.Build()

	resp, err := marketApi.GetCrossMarginSymbols(req, context.TODO())
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

func TestMarketGetMarkPriceListReq(t *testing.T) {
	// GetMarkPriceList
	// Get Mark Price List
	// /api/v3/mark-price/all-symbols

	resp, err := marketApi.GetMarkPriceList(context.TODO())
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
