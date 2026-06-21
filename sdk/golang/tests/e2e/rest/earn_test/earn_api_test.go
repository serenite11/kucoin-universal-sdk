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
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/generate/earn/earn"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/types"
)

var earnApi earn.EarnAPI

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

	earnApi = client.RestService().GetEarnService().GetEarnAPI()
}

func TestEarnGetETHStakingProductsReq(t *testing.T) {
	// GetETHStakingProducts
	// Get ETH Staking Products
	// /api/v1/earn/eth-staking/products

	builder := earn.NewGetETHStakingProductsReqBuilder()
	builder.SetCurrency("ETH")
	req := builder.Build()

	resp, err := earnApi.GetETHStakingProducts(req, context.TODO())
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
	return
}

func TestEarnGetAccountHoldingReq(t *testing.T) {
	// GetAccountHolding
	// Get Account Holding
	// /api/v1/earn/hold-assets

	builder := earn.NewGetAccountHoldingReqBuilder()
	builder.SetCurrency("USDT").SetProductId("2152").SetProductCategory("DEMAND")
	req := builder.Build()

	resp, err := earnApi.GetAccountHolding(req, context.TODO())
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

func TestEarnGetKcsStakingProductsReq(t *testing.T) {
	// GetKcsStakingProducts
	// Get KCS Staking Products
	// /api/v1/earn/kcs-staking/products

	builder := earn.NewGetKcsStakingProductsReqBuilder()
	builder.SetCurrency("KCS")
	req := builder.Build()

	resp, err := earnApi.GetKcsStakingProducts(req, context.TODO())
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

func TestEarnRedeemReq(t *testing.T) {
	// Redeem
	// Redeem
	// /api/v1/earn/orders

	builder := earn.NewRedeemReqBuilder()
	builder.SetOrderId("2849600").SetAmount("10.0").SetFromAccountType("MAIN").SetConfirmPunishRedeem("1")
	req := builder.Build()

	resp, err := earnApi.Redeem(req, context.TODO())
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

func TestEarnPurchaseReq(t *testing.T) {
	// Purchase
	// purchase
	// /api/v1/earn/orders

	builder := earn.NewPurchaseReqBuilder()
	builder.SetProductId("2152").SetAmount("10.0").SetAccountType("MAIN")
	req := builder.Build()

	resp, err := earnApi.Purchase(req, context.TODO())
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

func TestEarnGetPromotionProductsReq(t *testing.T) {
	// GetPromotionProducts
	// Get Promotion Products
	// /api/v1/earn/promotion/products

	builder := earn.NewGetPromotionProductsReqBuilder()
	builder.SetCurrency("USDT")
	req := builder.Build()

	resp, err := earnApi.GetPromotionProducts(req, context.TODO())
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

func TestEarnGetRedeemPreviewReq(t *testing.T) {
	// GetRedeemPreview
	// Get Redeem Preview
	// /api/v1/earn/redeem-preview

	builder := earn.NewGetRedeemPreviewReqBuilder()
	builder.SetOrderId("2849600").SetFromAccountType("MAIN")
	req := builder.Build()

	resp, err := earnApi.GetRedeemPreview(req, context.TODO())
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

func TestEarnGetSavingsProductsReq(t *testing.T) {
	// GetSavingsProducts
	// Get Savings Products
	// /api/v1/earn/saving/products

	builder := earn.NewGetSavingsProductsReqBuilder()
	builder.SetCurrency("USDT")
	req := builder.Build()

	resp, err := earnApi.GetSavingsProducts(req, context.TODO())
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

func TestEarnGetStakingProductsReq(t *testing.T) {
	// GetStakingProducts
	// Get Staking Products
	// /api/v1/earn/staking/products

	builder := earn.NewGetStakingProductsReqBuilder()
	builder.SetCurrency("ATOM")
	req := builder.Build()

	resp, err := earnApi.GetStakingProducts(req, context.TODO())
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
