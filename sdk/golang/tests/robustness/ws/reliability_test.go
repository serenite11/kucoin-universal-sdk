package rest

import (
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/extension/interceptor"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/api"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/common/logger"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/generate/spot/spotpublic"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestResubscribe(t *testing.T) {
	key := os.Getenv("API_KEY")
	secret := os.Getenv("API_SECRET")
	passphrase := os.Getenv("API_PASSPHRASE")

	defaultLogger := logger.NewDefaultLogger()
	logger.SetLogger(defaultLogger)

	httpOptionBuilder := types.NewTransportOptionBuilder().
		SetKeepAlive(true).
		AddInterceptors(interceptor.NewLoggingInterceptor(false, defaultLogger))

	websocketOptions := types.NewWebSocketClientOptionBuilder().Build()

	option := types.NewClientOptionBuilder().
		WithKey(key).
		WithSecret(secret).
		WithPassphrase(passphrase).
		WithSpotEndpoint(types.GlobalApiEndpoint).
		WithFuturesEndpoint(types.GlobalFuturesApiEndpoint).
		WithBrokerEndpoint(types.GlobalBrokerApiEndpoint).
		WithTransportOption(httpOptionBuilder.Build()).
		WithWebSocketClientOption(websocketOptions).
		Build()

	client := api.NewClient(option)
	spotWsApi := client.WsService().NewSpotPublicWS()
	err := spotWsApi.Start()
	assert.Nil(t, err)
	defer spotWsApi.Stop()

	btc := atomic.Bool{}
	eth := atomic.Bool{}
	doge := atomic.Bool{}

	cb := func(topic string, subject string, data *spotpublic.TickerEvent) error {
		if topic == "/market/ticker:BTC-USDT" {
			btc.Store(true)
		}

		if topic == "/market/ticker:ETH-USDT" {
			eth.Store(true)
		}

		if topic == "/market/ticker:DOGE-USDT" {
			doge.Store(true)
		}

		fmt.Println(topic)
		return nil
	}

	subid, err := spotWsApi.Ticker([]string{"BTC-USDT", "ETH-USDT"}, cb)
	assert.Nil(t, err)
	subid2, err := spotWsApi.Ticker([]string{"ETH-USDT", "DOGE-USDT"}, cb)
	assert.Nil(t, err)

	time.Sleep(2 * time.Second)

	assert.True(t, btc.Load())
	assert.True(t, eth.Load())
	assert.True(t, doge.Load())

	err = spotWsApi.UnSubscribe(subid2)
	assert.Nil(t, err)
	btc.Store(false)
	eth.Store(false)
	doge.Store(false)

	time.Sleep(2 * time.Second)

	assert.True(t, btc.Load())
	assert.False(t, eth.Load())
	assert.False(t, doge.Load())

	err = spotWsApi.UnSubscribe(subid)
	assert.Nil(t, err)
	btc.Store(false)
	eth.Store(false)
	doge.Store(false)
	time.Sleep(2 * time.Second)

	assert.False(t, btc.Load())
	assert.False(t, eth.Load())
	assert.False(t, doge.Load())
}

func TestResubscribe2(t *testing.T) {
	key := os.Getenv("API_KEY")
	secret := os.Getenv("API_SECRET")
	passphrase := os.Getenv("API_PASSPHRASE")

	defaultLogger := logger.NewDefaultLogger()
	logger.SetLogger(defaultLogger)

	httpOptionBuilder := types.NewTransportOptionBuilder().
		SetKeepAlive(true).
		AddInterceptors(interceptor.NewLoggingInterceptor(false, defaultLogger))

	websocketOptions := types.NewWebSocketClientOptionBuilder().Build()

	option := types.NewClientOptionBuilder().
		WithKey(key).
		WithSecret(secret).
		WithPassphrase(passphrase).
		WithSpotEndpoint(types.GlobalApiEndpoint).
		WithFuturesEndpoint(types.GlobalFuturesApiEndpoint).
		WithBrokerEndpoint(types.GlobalBrokerApiEndpoint).
		WithTransportOption(httpOptionBuilder.Build()).
		WithWebSocketClientOption(websocketOptions).
		Build()

	client := api.NewClient(option)
	spotWsApi := client.WsService().NewSpotPublicWS()
	err := spotWsApi.Start()
	assert.Nil(t, err)
	defer spotWsApi.Stop()

	cb := func(topic string, subject string, data *spotpublic.TickerEvent) error {
		return nil
	}

	_, err = spotWsApi.Ticker([]string{"BTC-USDT", "ETH-USDT"}, cb)
	assert.Nil(t, err)
	_, err = spotWsApi.Ticker([]string{"BTC-USDT", "ETH-USDT"}, cb)
	assert.NotNil(t, err)
}

func TestStartStop(t *testing.T) {
	key := os.Getenv("API_KEY")
	secret := os.Getenv("API_SECRET")
	passphrase := os.Getenv("API_PASSPHRASE")

	defaultLogger := logger.NewDefaultLogger()
	logger.SetLogger(defaultLogger)

	httpOptionBuilder := types.NewTransportOptionBuilder().
		SetKeepAlive(true).
		AddInterceptors(interceptor.NewLoggingInterceptor(false, defaultLogger))

	websocketOptions := types.NewWebSocketClientOptionBuilder().Build()

	option := types.NewClientOptionBuilder().
		WithKey(key).
		WithSecret(secret).
		WithPassphrase(passphrase).
		WithSpotEndpoint(types.GlobalApiEndpoint).
		WithFuturesEndpoint(types.GlobalFuturesApiEndpoint).
		WithBrokerEndpoint(types.GlobalBrokerApiEndpoint).
		WithTransportOption(httpOptionBuilder.Build()).
		WithWebSocketClientOption(websocketOptions).
		Build()

	client := api.NewClient(option)
	spotWsApi := client.WsService().NewSpotPublicWS()
	err := spotWsApi.Start()
	assert.Nil(t, err)

	cb := func(topic string, subject string, data *spotpublic.TickerEvent) error {
		fmt.Println(data.ToMap())
		return nil
	}

	_, err = spotWsApi.Ticker([]string{"BTC-USDT", "ETH-USDT"}, cb)
	assert.Nil(t, err)

	time.Sleep(2 * time.Second)
	spotWsApi.Stop()
	time.Sleep(2 * time.Second)
}

func TestConcurrentSubscribe(t *testing.T) {
	key := os.Getenv("API_KEY")
	secret := os.Getenv("API_SECRET")
	passphrase := os.Getenv("API_PASSPHRASE")

	defaultLogger := logger.NewDefaultLogger()
	logger.SetLogger(defaultLogger)

	httpOptionBuilder := types.NewTransportOptionBuilder().
		SetKeepAlive(true).
		AddInterceptors(interceptor.NewLoggingInterceptor(false, defaultLogger))

	websocketOptions := types.NewWebSocketClientOptionBuilder().Build()

	option := types.NewClientOptionBuilder().
		WithKey(key).
		WithSecret(secret).
		WithPassphrase(passphrase).
		WithSpotEndpoint(types.GlobalApiEndpoint).
		WithFuturesEndpoint(types.GlobalFuturesApiEndpoint).
		WithBrokerEndpoint(types.GlobalBrokerApiEndpoint).
		WithTransportOption(httpOptionBuilder.Build()).
		WithWebSocketClientOption(websocketOptions).
		Build()

	client := api.NewClient(option)
	spotWsApi := client.WsService().NewSpotPublicWS()
	err := spotWsApi.Start()
	assert.Nil(t, err)
	defer spotWsApi.Stop()

	symbol := []string{"BTC-USDT", "ETH-USDT", "DOGE-USDT"}

	t1 := atomic.Int32{}
	t2 := atomic.Int32{}
	t3 := atomic.Int32{}
	t4 := atomic.Int32{}
	t5 := atomic.Int32{}

	go func() {
		_, err = spotWsApi.Ticker(symbol, func(topic string, subject string, data *spotpublic.TickerEvent) error {
			t1.Add(1)
			return nil
		})
		assert.Nil(t, err)
	}()
	go func() {
		_, err = spotWsApi.AllTickers(func(topic string, subject string, data *spotpublic.AllTickersEvent) error {
			t2.Add(1)
			return nil
		})
		assert.Nil(t, err)
	}()

	go func() {
		_, err = spotWsApi.OrderbookLevel1(symbol, func(topic string, subject string, data *spotpublic.OrderbookLevel1Event) error {
			t3.Add(1)
			return nil
		})
		assert.Nil(t, err)
	}()

	go func() {
		_, err = spotWsApi.OrderbookLevel5(symbol, func(topic string, subject string, data *spotpublic.OrderbookLevel5Event) error {
			t4.Add(1)
			return nil
		})
		assert.Nil(t, err)
	}()

	go func() {
		_, err = spotWsApi.OrderbookLevel50(symbol, func(topic string, subject string, data *spotpublic.OrderbookLevel50Event) error {
			t5.Add(1)
			return nil
		})
		assert.Nil(t, err)
	}()

	time.Sleep(2 * time.Second)

	assert.True(t, t1.Load() > 0)
	assert.True(t, t2.Load() > 0)
	assert.True(t, t3.Load() > 0)
	assert.True(t, t4.Load() > 0)
	assert.True(t, t5.Load() > 0)
}

func TestReconnect(t *testing.T) {
	key := os.Getenv("API_KEY")
	secret := os.Getenv("API_SECRET")
	passphrase := os.Getenv("API_PASSPHRASE")

	defaultLogger := logger.NewDefaultLogger()
	logger.SetLogger(defaultLogger)

	httpOptionBuilder := types.NewTransportOptionBuilder().
		SetKeepAlive(true).
		AddInterceptors(interceptor.NewLoggingInterceptor(false, defaultLogger))

	websocketOptions := types.NewWebSocketClientOptionBuilder().WithReconnect(true).WithEventCallback(func(event types.WebSocketEvent, msg string) {
		fmt.Printf("event callback, event:%s, msg: %s", event.String(), msg)
	}).Build()

	option := types.NewClientOptionBuilder().
		WithKey(key).
		WithSecret(secret).
		WithPassphrase(passphrase).
		WithSpotEndpoint(types.GlobalApiEndpoint).
		WithFuturesEndpoint(types.GlobalFuturesApiEndpoint).
		WithBrokerEndpoint(types.GlobalBrokerApiEndpoint).
		WithTransportOption(httpOptionBuilder.Build()).
		WithWebSocketClientOption(websocketOptions).
		Build()

	client := api.NewClient(option)
	spotWsApi := client.WsService().NewSpotPublicWS()
	err := spotWsApi.Start()
	assert.Nil(t, err)
	defer spotWsApi.Stop()

	_, err = spotWsApi.Trade([]string{"BTC-USDT", "ETH-USDT"}, func(topic string, subject string, data *spotpublic.TradeEvent) error {
		fmt.Println(topic)
		return nil
	})
	assert.Nil(t, err)
	_, err = spotWsApi.Ticker([]string{"BTC-USDT", "ETH-USDT"}, func(topic string, subject string, data *spotpublic.TickerEvent) error {
		fmt.Println(topic)
		return nil
	})
	assert.Nil(t, err)

	time.Sleep(600 * time.Second)
	// disconnect network manually
}

func TestReadBufferFull(t *testing.T) {
	key := os.Getenv("API_KEY")
	secret := os.Getenv("API_SECRET")
	passphrase := os.Getenv("API_PASSPHRASE")

	defaultLogger := logger.NewDefaultLogger()
	logger.SetLogger(defaultLogger)

	httpOptionBuilder := types.NewTransportOptionBuilder().
		SetKeepAlive(true).
		AddInterceptors(interceptor.NewLoggingInterceptor(false, defaultLogger))

	set := atomic.Bool{}
	wg := sync.WaitGroup{}
	wg.Add(1)
	websocketOptions := types.NewWebSocketClientOptionBuilder().WithReconnect(true).WithEventCallback(func(event types.WebSocketEvent, msg string) {
		fmt.Printf("event callback, event:%s, msg: %s\n", event.String(), msg)

		if event == types.EventReadBufferFull {
			if set.CompareAndSwap(false, true) {
				wg.Done()
			}
		}

	}).WithReadMessageBuffer(1).Build()

	option := types.NewClientOptionBuilder().
		WithKey(key).
		WithSecret(secret).
		WithPassphrase(passphrase).
		WithSpotEndpoint(types.GlobalApiEndpoint).
		WithFuturesEndpoint(types.GlobalFuturesApiEndpoint).
		WithBrokerEndpoint(types.GlobalBrokerApiEndpoint).
		WithTransportOption(httpOptionBuilder.Build()).
		WithWebSocketClientOption(websocketOptions).
		Build()

	client := api.NewClient(option)
	spotWsApi := client.WsService().NewSpotPublicWS()
	err := spotWsApi.Start()
	assert.Nil(t, err)
	defer spotWsApi.Stop()

	_, err = spotWsApi.Ticker([]string{"BTC-USDT", "ETH-USDT"}, func(topic string, subject string, data *spotpublic.TickerEvent) error {
		time.Sleep(2 * time.Second)
		return nil
	})
	assert.Nil(t, err)

	wg.Wait()
}
