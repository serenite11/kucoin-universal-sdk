package spot_test

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"testing"

	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/api"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/common/logger"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/generate/margin/marginprivate"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/types"
	"github.com/stretchr/testify/assert"
)

var marginPrivateApi marginprivate.MarginPrivateWS

func init() {
	logger.SetLogger(
		&logger.DefaultLogger{
			Logger:     log.New(os.Stdout, "", 0),
			Level:      logger.TRACE,
			Formatter:  "%s [%s] [%s+%d] [%s] %s",
			TimeFormat: "2006-01-02 15:04:05.000",
		})

	// get api key from env
	key := os.Getenv("API_KEY")
	secret := os.Getenv("API_SECRET")
	passphrase := os.Getenv("API_PASSPHRASE")

	// set callback worker number
	wsOption := types.NewWebSocketClientOptionBuilder().
		Build()

	// create a client with transport options
	option := types.NewClientOptionBuilder().
		WithKey(key).
		WithSecret(secret).
		WithPassphrase(passphrase).
		WithSpotEndpoint(types.GlobalApiEndpoint).
		WithFuturesEndpoint(types.GlobalFuturesApiEndpoint).
		WithBrokerEndpoint(types.GlobalBrokerApiEndpoint).
		WithWebSocketClientOption(wsOption).
		Build()

	client := api.NewClient(option)

	marginPrivateApi = client.WsService().NewMarginPrivateWS()
	err := marginPrivateApi.Start()
	if err != nil {
		panic(err)
	}
}

func TestPrivateCrossMarginPosition(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	i := 0
	id, err := marginPrivateApi.CrossMarginPosition(func(topic string, subject string, data *marginprivate.CrossMarginPositionEvent) error {
		assert.NotNil(t, data)
		i++
		if i == 1 {
			wg.Done()
		}
		str, _ := json.Marshal(data)
		fmt.Println(topic, subject, string(str))
		return nil
	})
	assert.Nil(t, err)
	wg.Wait()

	err = marginPrivateApi.UnSubscribe(id)
	assert.Nil(t, err)
}

func TestPrivateIsolatedMarginPosition(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	i := 0
	id, err := marginPrivateApi.IsolatedMarginPosition("DOGE-USDT", func(topic string, subject string, data *marginprivate.IsolatedMarginPositionEvent) error {
		assert.NotNil(t, data)
		i++
		if i == 1 {
			wg.Done()
		}
		str, _ := json.Marshal(data)
		fmt.Println(topic, subject, string(str))
		return nil
	})
	assert.Nil(t, err)
	wg.Wait()

	err = marginPrivateApi.UnSubscribe(id)
	assert.Nil(t, err)
}
