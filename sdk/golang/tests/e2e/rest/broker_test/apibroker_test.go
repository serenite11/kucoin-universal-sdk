package broker_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/extension/interceptor"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/api"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/common/logger"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/generate/broker/apibroker"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/types"
)

var apibrokerApi apibroker.APIBrokerAPI

func init() {
	key := os.Getenv("API_KEY")
	secret := os.Getenv("API_SECRET")
	passphrase := os.Getenv("API_PASSPHRASE")

	brokerName := os.Getenv("BROKER_NAME")
	brokerKey := os.Getenv("BROKER_KEY")
	brokerPartner := os.Getenv("BROKER_PARTNER")

	defaultLogger := logger.NewDefaultLogger()
	logger.SetLogger(defaultLogger)

	httpOptionBuilder := types.NewTransportOptionBuilder().
		SetKeepAlive(true).
		AddInterceptors(interceptor.NewLoggingInterceptor(false, defaultLogger))

	option := types.NewClientOptionBuilder().
		WithKey(key).
		WithSecret(secret).
		WithPassphrase(passphrase).
		WithBrokerName(brokerName).
		WithBrokerPartner(brokerPartner).
		WithBrokerKey(brokerKey).
		WithSpotEndpoint(types.GlobalApiEndpoint).
		WithFuturesEndpoint(types.GlobalFuturesApiEndpoint).
		WithBrokerEndpoint(types.GlobalBrokerApiEndpoint).
		WithTransportOption(httpOptionBuilder.Build()).
		Build()

	client := api.NewClient(option)

	apibrokerApi = client.RestService().GetBrokerService().GetAPIBrokerAPI()
}

// TODO no permission
func TestAPIBrokerGetRebaseReq(t *testing.T) {
	// GetRebase
	// Get Broker Rebate
	// /api/v1/broker/api/rebase/download

	builder := apibroker.NewGetRebaseReqBuilder()
	builder.SetBegin("20240610").SetEnd("20241010").SetTradeType("1")
	req := builder.Build()

	resp, err := apibrokerApi.GetRebase(req, context.TODO())
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
