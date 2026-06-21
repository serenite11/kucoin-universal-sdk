package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/extension/interceptor"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/api"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/common/logger"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestTimeout(t *testing.T) {
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
	marketApi := client.RestService().GetSpotService().GetMarketAPI()

	ctx, cFunc := context.WithTimeout(context.Background(), time.Microsecond)
	defer cFunc()

	_, err := marketApi.GetServerTime(ctx)
	assert.NotNil(t, err)
	fmt.Println(err.Error())
}

func TestTimeoutRetry(t *testing.T) {
	key := os.Getenv("API_KEY")
	secret := os.Getenv("API_SECRET")
	passphrase := os.Getenv("API_PASSPHRASE")

	defaultLogger := logger.NewDefaultLogger()
	logger.SetLogger(defaultLogger)

	httpOptionBuilder := types.NewTransportOptionBuilder().
		SetTimeout(time.Microsecond).
		SetKeepAlive(true).
		SetRetryDelay(time.Second).
		SetMaxRetries(2).
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
	marketApi := client.RestService().GetSpotService().GetMarketAPI()

	start := time.Now()
	_, err := marketApi.GetServerTime(context.Background())
	end := time.Now()
	assert.NotNil(t, err)
	fmt.Println(err.Error())
	cost := end.Sub(start)
	assert.True(t, cost > 2*time.Second && cost < 3*time.Second)
}

func TestProxy(t *testing.T) {
	key := os.Getenv("API_KEY")
	secret := os.Getenv("API_SECRET")
	passphrase := os.Getenv("API_PASSPHRASE")

	proxyURL, err := url.Parse("http://192.168.1.1:8080")
	if err != nil {
		log.Fatalf("Failed to parse proxy URL: %v", err)
	}

	proxyURL.String()

	defaultLogger := logger.NewDefaultLogger()
	logger.SetLogger(defaultLogger)

	httpOptionBuilder := types.NewTransportOptionBuilder().
		SetTimeout(time.Second * 5).
		SetKeepAlive(true).
		SetRetryDelay(time.Second).
		SetMaxRetries(2).
		SetProxy(http.ProxyURL(proxyURL)).
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
	marketApi := client.RestService().GetSpotService().GetMarketAPI()

	_, err = marketApi.GetServerTime(context.Background())
	fmt.Println(err)
}
