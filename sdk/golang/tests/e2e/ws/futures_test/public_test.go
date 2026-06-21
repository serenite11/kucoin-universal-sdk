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
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/generate/futures/futurespublic"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/types"
	"github.com/stretchr/testify/assert"
)

var futuresPublicApi futurespublic.FuturesPublicWS

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

	futuresPublicApi = client.WsService().NewFuturesPublicWS()
	err := futuresPublicApi.Start()
	if err != nil {
		panic(err)
	}
}

// TODO
func TestPublicAnnouncement(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	i := 0
	id, err := futuresPublicApi.Announcement("XBTUSDTM", func(topic string, subject string, data *futurespublic.AnnouncementEvent) error {
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

	err = futuresPublicApi.UnSubscribe(id)
	assert.Nil(t, err)
}

func TestPublicExecution(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	i := 0
	id, err := futuresPublicApi.Execution("XBTUSDTM", func(topic string, subject string, data *futurespublic.ExecutionEvent) error {
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

	err = futuresPublicApi.UnSubscribe(id)
	assert.Nil(t, err)
}

func TestPublicInstrument(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	i := 0
	id, err := futuresPublicApi.Instrument("XBTUSDTM", func(topic string, subject string, data *futurespublic.InstrumentEvent) error {
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

	err = futuresPublicApi.UnSubscribe(id)
	assert.Nil(t, err)
}

func TestPublicKlines(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	i := 0
	id, err := futuresPublicApi.Klines("XBTUSDTM", "1min", func(topic string, subject string, data *futurespublic.KlinesEvent) error {
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

	err = futuresPublicApi.UnSubscribe(id)
	assert.Nil(t, err)
}

func TestPublicOrderbookIncrement(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	i := 0
	id, err := futuresPublicApi.OrderbookIncrement("XBTUSDTM", func(topic string, subject string, data *futurespublic.OrderbookIncrementEvent) error {
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

	err = futuresPublicApi.UnSubscribe(id)
	assert.Nil(t, err)
}

func TestPublicOrderbookLevel50(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	i := 0
	id, err := futuresPublicApi.OrderbookLevel50("XBTUSDTM", func(topic string, subject string, data *futurespublic.OrderbookLevel50Event) error {
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

	err = futuresPublicApi.UnSubscribe(id)
	assert.Nil(t, err)
}

func TestPublicOrderbookLevel5(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	i := 0
	id, err := futuresPublicApi.OrderbookLevel5("XBTUSDTM", func(topic string, subject string, data *futurespublic.OrderbookLevel5Event) error {
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

	err = futuresPublicApi.UnSubscribe(id)
	assert.Nil(t, err)
}

func TestPublicSymbolSnapshot(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	i := 0
	id, err := futuresPublicApi.SymbolSnapshot("XBTUSDTM", func(topic string, subject string, data *futurespublic.SymbolSnapshotEvent) error {
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

	err = futuresPublicApi.UnSubscribe(id)
	assert.Nil(t, err)
}

func TestPublicTickerV1(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	i := 0
	id, err := futuresPublicApi.TickerV1("XBTUSDTM", func(topic string, subject string, data *futurespublic.TickerV1Event) error {
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

	err = futuresPublicApi.UnSubscribe(id)
	assert.Nil(t, err)
}

func TestPublicTickerV2(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	i := 0
	id, err := futuresPublicApi.TickerV2("XBTUSDTM", func(topic string, subject string, data *futurespublic.TickerV2Event) error {
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

	err = futuresPublicApi.UnSubscribe(id)
	assert.Nil(t, err)
}
