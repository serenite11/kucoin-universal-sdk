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
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/generate/futures/futuresprivate"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/types"
	"github.com/stretchr/testify/assert"
)

var futuresPrivateApi futuresprivate.FuturesPrivateWS

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

	futuresPrivateApi = client.WsService().NewFuturesPrivateWS()
	err := futuresPrivateApi.Start()
	if err != nil {
		panic(err)
	}
}

func TestPrivateAllOrder(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	i := 0
	id, err := futuresPrivateApi.AllOrder(func(topic string, subject string, data *futuresprivate.AllOrderEvent) error {
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

	err = futuresPrivateApi.UnSubscribe(id)
	assert.Nil(t, err)
}

func TestPrivateAllPosition(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	i := 0
	id, err := futuresPrivateApi.AllPosition(func(topic string, subject string, data *futuresprivate.AllPositionEvent) error {
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

	err = futuresPrivateApi.UnSubscribe(id)
	assert.Nil(t, err)
}

func TestPrivateBalance(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	i := 0
	id, err := futuresPrivateApi.Balance(func(topic string, subject string, data *futuresprivate.BalanceEvent) error {
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

	err = futuresPrivateApi.UnSubscribe(id)
	assert.Nil(t, err)
}

func TestPrivateCrossLeverage(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	i := 0
	id, err := futuresPrivateApi.CrossLeverage(func(topic string, subject string, data *futuresprivate.CrossLeverageEvent) error {
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

	err = futuresPrivateApi.UnSubscribe(id)
	assert.Nil(t, err)
}

func TestPrivateMarginMode(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	i := 0
	id, err := futuresPrivateApi.MarginMode(func(topic string, subject string, data *futuresprivate.MarginModeEvent) error {
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

	err = futuresPrivateApi.UnSubscribe(id)
	assert.Nil(t, err)
}

func TestPrivateOrder(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	i := 0
	id, err := futuresPrivateApi.Order("XBTUSDTM", func(topic string, subject string, data *futuresprivate.OrderEvent) error {
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

	err = futuresPrivateApi.UnSubscribe(id)
	assert.Nil(t, err)
}

func TestPrivatePosition(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	i := 0
	id, err := futuresPrivateApi.Position("DOGEUSDTM", func(topic string, subject string, data *futuresprivate.PositionEvent) error {
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

	err = futuresPrivateApi.UnSubscribe(id)
	assert.Nil(t, err)
}

func TestPrivateStopOrders(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	i := 0
	id, err := futuresPrivateApi.StopOrders(func(topic string, subject string, data *futuresprivate.StopOrdersEvent) error {
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

	err = futuresPrivateApi.UnSubscribe(id)
	assert.Nil(t, err)
}
