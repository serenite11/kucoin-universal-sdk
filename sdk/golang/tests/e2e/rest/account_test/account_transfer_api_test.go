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
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/generate/account/transfer"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/types"
)

var transferApi transfer.TransferAPI

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

	transferApi = client.RestService().GetAccountService().GetTransferAPI()
}

func TestTransferGetTransferQuotasReq(t *testing.T) {
	// GetTransferQuotas
	// Get Transfer Quotas
	// /api/v1/accounts/transferable

	builder := transfer.NewGetTransferQuotasReqBuilder()
	builder.SetCurrency("USDT").SetType("TRADE")
	req := builder.Build()

	resp, err := transferApi.GetTransferQuotas(req, context.TODO())
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

func TestTransferFuturesAccountTransferInReq(t *testing.T) {
	// FuturesAccountTransferIn
	// Futures Account Transfer In
	// /api/v1/transfer-in

	builder := transfer.NewFuturesAccountTransferInReqBuilder()
	builder.SetCurrency("USDT").SetAmount(0.01).SetPayAccountType("MAIN")
	req := builder.Build()

	resp, err := transferApi.FuturesAccountTransferIn(req, context.TODO())
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

func TestTransferGetFuturesAccountTransferOutLedgerReq(t *testing.T) {
	// GetFuturesAccountTransferOutLedger
	// Get Futures Account Transfer Out Ledger
	// /api/v1/transfer-list

	builder := transfer.NewGetFuturesAccountTransferOutLedgerReqBuilder()
	builder.SetCurrency("USDT").SetType("SUCCESS")
	req := builder.Build()

	resp, err := transferApi.GetFuturesAccountTransferOutLedger(req, context.TODO())
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

func TestTransferInnerTransferReq(t *testing.T) {
	// InnerTransfer
	// Inner Transfer
	// /api/v2/accounts/inner-transfer

	builder := transfer.NewInnerTransferReqBuilder()
	builder.SetClientOid("11122334455").
		SetCurrency("USDT").SetAmount("0.01").SetFrom("trade").
		SetTo("isolated").SetToTag("BTC-USDT")
	req := builder.Build()

	resp, err := transferApi.InnerTransfer(req, context.TODO())
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

func TestTransferSubAccountTransferReq(t *testing.T) {
	// SubAccountTransfer
	// SubAccount Transfer
	// /api/v2/accounts/sub-transfer

	builder := transfer.NewSubAccountTransferReqBuilder()
	builder.SetClientOid("11223344").SetCurrency("USDT").SetAmount("10").
		SetDirection("OUT").SetAccountType("TRADE").SetSubAccountType("TRADE").
		SetSubUserId("66ab063aaea134000115b4c2")
	req := builder.Build()

	resp, err := transferApi.SubAccountTransfer(req, context.TODO())
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

func TestTransferFlexTransferReq(t *testing.T) {
	// FlexTransfer
	// Flex Transfer
	// /api/v3/accounts/universal-transfer

	builder := transfer.NewFlexTransferReqBuilder()
	builder.SetClientOid("11223344").SetCurrency("USDT").
		SetAmount("0.01").SetFromUserId("66ab063aaea134000115b4c2").SetFromAccountType("TRADE").
		SetType("INTERNAL").SetToUserId("66ab063aaea134000115b4c2").SetToAccountType("MAIN")
	req := builder.Build()

	resp, err := transferApi.FlexTransfer(req, context.TODO())
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

func TestTransferFuturesAccountTransferOutReq(t *testing.T) {
	// FuturesAccountTransferOut
	// Futures Account Transfer Out
	// /api/v3/transfer-out

	builder := transfer.NewFuturesAccountTransferOutReqBuilder()
	builder.SetCurrency("USDT").SetAmount(0.01).SetRecAccountType("MAIN")
	req := builder.Build()

	resp, err := transferApi.FuturesAccountTransferOut(req, context.TODO())
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
