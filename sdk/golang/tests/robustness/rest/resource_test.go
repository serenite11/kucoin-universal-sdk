package rest

import (
	"context"
	"fmt"
	"net/http/httptrace"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/extension/interceptor"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/api"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/common/logger"
	futuresmarket "github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/generate/futures/market"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/generate/spot/market"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/types"
	"github.com/stretchr/testify/assert"
)

var rSpotMarketApi market.MarketAPI
var rFuturesMarketApi futuresmarket.MarketAPI

func init() {
	key := os.Getenv("API_KEY")
	secret := os.Getenv("API_SECRET")
	passphrase := os.Getenv("API_PASSPHRASE")

	defaultLogger := logger.NewDefaultLogger()
	logger.SetLogger(defaultLogger)

	httpOptionBuilder := types.NewTransportOptionBuilder().
		SetKeepAlive(true).
		SetMaxConnsPerHost(2).
		SetIdleConnTimeout(1 * time.Second).
		SetMaxIdleConnsPerHost(2).
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
	rSpotMarketApi = client.RestService().GetSpotService().GetMarketAPI()
	rFuturesMarketApi = client.RestService().GetFuturesService().GetMarketAPI()
}

func TestConcurrency(t *testing.T) {

	wg := &sync.WaitGroup{}
	wg.Add(10 * 4)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			resp, err := rSpotMarketApi.GetServerTime(context.Background())
			assert.Nil(t, err)
			assert.True(t, resp.Data > 0)
		}()

		go func() {
			defer wg.Done()
			resp, err := rSpotMarketApi.GetMarketList(context.Background())
			assert.Nil(t, err)
			assert.True(t, len(resp.Data) > 0)
		}()

		go func() {
			defer wg.Done()
			resp, err := rFuturesMarketApi.GetServerTime(context.Background())
			assert.Nil(t, err)
			assert.True(t, resp.Data > 0)
		}()

		go func() {
			defer wg.Done()
			resp, err := rFuturesMarketApi.GetAllSymbols(context.Background())
			assert.Nil(t, err)
			assert.True(t, len(resp.Data) > 0)
		}()
	}
	wg.Wait()
}

func printResouceUsage() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	fmt.Printf(`===================
Alloc = %v KB
TotalAlloc = %v KB
Sys = %v KB
NumGC = %v
Goroutines: %d
===================
`, memStats.Alloc/1024, memStats.TotalAlloc/1024, memStats.Sys/1024, memStats.NumGC, runtime.NumGoroutine())
}

func TestResourceLeak(t *testing.T) {

	wg := &sync.WaitGroup{}
	wg.Add(10 * 10 * 4)
	go func() {
		for i := 0; i < 10; i++ {
			printResouceUsage()
			time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		for j := 0; j < 10; j++ {

			go func() {
				defer wg.Done()
				resp, err := rSpotMarketApi.GetServerTime(context.Background())
				assert.Nil(t, err)
				assert.True(t, resp.Data > 0)
			}()

			go func() {
				defer wg.Done()
				resp, err := rSpotMarketApi.GetMarketList(context.Background())
				assert.Nil(t, err)
				assert.True(t, len(resp.Data) > 0)
			}()

			go func() {
				defer wg.Done()
				resp, err := rFuturesMarketApi.GetServerTime(context.Background())
				assert.Nil(t, err)
				assert.True(t, resp.Data > 0)
			}()

			go func() {
				defer wg.Done()
				resp, err := rFuturesMarketApi.GetAllSymbols(context.Background())
				assert.Nil(t, err)
				assert.True(t, len(resp.Data) > 0)
			}()
		}
	}
	wg.Wait()
	printResouceUsage()
}

func TestTcpLeak(t *testing.T) {

	c := int32(0)

	trace := &httptrace.ClientTrace{
		ConnectDone: func(network, addr string, err error) {
			atomic.AddInt32(&c, 1)
			fmt.Printf("connected network:%s, addr:%s", network, addr)
		},
	}

	wg := &sync.WaitGroup{}
	wg.Add(10 * 4)
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)

		go func() {
			defer wg.Done()
			resp, err := rSpotMarketApi.GetServerTime(httptrace.WithClientTrace(context.Background(), trace))
			assert.Nil(t, err)
			assert.True(t, resp.Data > 0)
		}()

		go func() {
			defer wg.Done()
			resp, err := rSpotMarketApi.GetMarketList(httptrace.WithClientTrace(context.Background(), trace))
			assert.Nil(t, err)
			assert.True(t, len(resp.Data) > 0)
		}()

		go func() {
			defer wg.Done()
			resp, err := rFuturesMarketApi.GetServerTime(httptrace.WithClientTrace(context.Background(), trace))
			assert.Nil(t, err)
			assert.True(t, resp.Data > 0)
		}()

		go func() {
			defer wg.Done()
			resp, err := rFuturesMarketApi.GetAllSymbols(httptrace.WithClientTrace(context.Background(), trace))
			assert.Nil(t, err)
			assert.True(t, len(resp.Data) > 0)
		}()
	}
	wg.Wait()
	assert.Equal(t, int32(4), atomic.LoadInt32(&c))
}
