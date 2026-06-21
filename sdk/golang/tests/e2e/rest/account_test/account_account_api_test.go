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
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/generate/account/account"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/types"
)

var accountApi account.AccountAPI

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

	accountApi = client.RestService().GetAccountService().GetAccountAPI()
}

func TestAccountGetFuturesAccountReq(t *testing.T) {
	// GetFuturesAccount
	// Get Account - Futures
	// /api/v1/account-overview

	builder := account.NewGetFuturesAccountReqBuilder()
	builder.SetCurrency("XBT")
	req := builder.Build()

	resp, err := accountApi.GetFuturesAccount(req, context.TODO())
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

func TestAccountGetSpotAccountDetailReq(t *testing.T) {
	// GetSpotAccountDetail
	// Get Account Detail - Spot
	// /api/v1/accounts/{accountId}

	builder := account.NewGetSpotAccountDetailReqBuilder()
	builder.SetAccountId("66ab0716b9ab290007f8285d")
	req := builder.Build()

	resp, err := accountApi.GetSpotAccountDetail(req, context.TODO())
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

func TestAccountGetSpotAccountListReq(t *testing.T) {
	// GetSpotAccountList
	// Get Account List - Spot
	// /api/v1/accounts

	builder := account.NewGetSpotAccountListReqBuilder()
	builder.SetType("main")
	req := builder.Build()

	resp, err := accountApi.GetSpotAccountList(req, context.TODO())
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

func TestAccountGetSpotLedgerReq(t *testing.T) {
	// GetSpotLedger
	// Get Account Ledgers - Spot/Margin
	// /api/v1/accounts/ledgers

	builder := account.NewGetSpotLedgerReqBuilder()
	builder.SetCurrency("USDT").
		SetStartAt(1732032000000).
		SetEndAt(1732118400000).
		SetCurrentPage(1).
		SetPageSize(100)
	req := builder.Build()

	resp, err := accountApi.GetSpotLedger(req, context.TODO())
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

func TestAccountGetSpotHFLedgerReq(t *testing.T) {
	// GetSpotHFLedger
	// Get Account Ledgers - Trade_hf
	// /api/v1/hf/accounts/ledgers

	builder := account.NewGetSpotHFLedgerReqBuilder()
	builder.SetCurrency("USDT").SetDirection("out").SetBizType("TRANSFER")
	req := builder.Build()

	resp, err := accountApi.GetSpotHFLedger(req, context.TODO())
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

func TestAccountGetSpotAccountTypeReq(t *testing.T) {
	// GetSpotAccountType
	// Get Account Type - Spot
	// /api/v1/hf/accounts/opened

	resp, err := accountApi.GetSpotAccountType(context.TODO())
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

func TestAccountGetIsolatedMarginAccountDetailV1Req(t *testing.T) {
	// GetIsolatedMarginAccountDetailV1
	// Get Account Detail - Isolated Margin V1
	// /api/v1/isolated/account/{symbol}

	builder := account.NewGetIsolatedMarginAccountDetailV1ReqBuilder()
	builder.SetSymbol("BTC-USDT")
	req := builder.Build()

	resp, err := accountApi.GetIsolatedMarginAccountDetailV1(req, context.TODO())
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

func TestAccountGetIsolatedMarginAccountListV1Req(t *testing.T) {
	// GetIsolatedMarginAccountListV1
	// Get Account List - Isolated Margin V1
	// /api/v1/isolated/accounts

	builder := account.NewGetIsolatedMarginAccountListV1ReqBuilder()
	builder.SetBalanceCurrency("BTC")
	req := builder.Build()

	resp, err := accountApi.GetIsolatedMarginAccountListV1(req, context.TODO())
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

func TestAccountGetMarginAccountDetailReq(t *testing.T) {
	// GetMarginAccountDetail
	// Get Account Detail - Margin
	// /api/v1/margin/account

	resp, err := accountApi.GetMarginAccountDetail(context.TODO())
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

func TestAccountGetFuturesLedgerReq(t *testing.T) {
	// GetFuturesLedger
	// Get Account Ledgers - Futures
	// /api/v1/transaction-history

	builder := account.NewGetFuturesLedgerReqBuilder()
	builder.SetCurrency("DOT").SetType("RealisedPNL")
	req := builder.Build()

	resp, err := accountApi.GetFuturesLedger(req, context.TODO())
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

func TestAccountGetApikeyInfoReq(t *testing.T) {
	// GetApikeyInfo
	// Get API Key Information
	// /api/v1/user/api-key

	resp, err := accountApi.GetApikeyInfo(context.TODO())
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

func TestAccountGetAccountInfoReq(t *testing.T) {
	// GetAccountInfo
	// Get Account Summary Info
	// /api/v2/user-info

	resp, err := accountApi.GetAccountInfo(context.TODO())
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

func TestAccountGetMarginHFLedgerReq(t *testing.T) {
	// GetMarginHFLedger
	// Get Account Ledgers - Margin_hf
	// /api/v3/hf/margin/account/ledgers

	builder := account.NewGetMarginHFLedgerReqBuilder()
	builder.SetCurrency("USDT")
	req := builder.Build()

	resp, err := accountApi.GetMarginHFLedger(req, context.TODO())
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

func TestAccountGetIsolatedMarginAccountReq(t *testing.T) {
	// GetIsolatedMarginAccount
	// Get Account - Isolated Margin
	// /api/v3/isolated/accounts

	builder := account.NewGetIsolatedMarginAccountReqBuilder()
	builder.SetSymbol("BTC-USDT").SetQuoteCurrency("USDT").SetQueryType("ISOLATED")
	req := builder.Build()

	resp, err := accountApi.GetIsolatedMarginAccount(req, context.TODO())
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

func TestAccountGetCrossMarginAccountReq(t *testing.T) {
	// GetCrossMarginAccount
	// Get Account - Cross Margin
	// /api/v3/margin/accounts

	builder := account.NewGetCrossMarginAccountReqBuilder()
	builder.SetQuoteCurrency("USDT").SetQueryType("MARGIN")
	req := builder.Build()

	resp, err := accountApi.GetCrossMarginAccount(req, context.TODO())
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
