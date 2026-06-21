package rest

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sync/atomic"
	"testing"
	"time"

	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/extension/interceptor"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/api"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/common/logger"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/generate/spot/spotpublic"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/types"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/tests/robustness/util"
	"github.com/stretchr/testify/assert"
)

func TestResourceLeak(t *testing.T) {
	key := os.Getenv("API_KEY")
	secret := os.Getenv("API_SECRET")
	passphrase := os.Getenv("API_PASSPHRASE")

	defaultLogger := logger.NewDefaultLogger()
	logger.SetLogger(defaultLogger)

	httpOptionBuilder := types.NewTransportOptionBuilder().
		SetKeepAlive(true).
		AddInterceptors(interceptor.NewLoggingInterceptor(false, defaultLogger))

	websocketOptions := types.NewWebSocketClientOptionBuilder().WithReconnect(true).Build()

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

	tracker := &util.MemStatsTracker{}

	c := atomic.Int32{}
	_, err = spotWsApi.AllTickers(func(topic string, subject string, data *spotpublic.AllTickersEvent) error {
		c.Add(1)
		if c.Load()%10 == 0 {
			tracker.Record()
			runtime.GC()
		}
		return nil
	})
	assert.Nil(t, err)

	time.Sleep(10 * time.Second)
	tracker.PrintStats()
}

func GetTcpStat() {
	pid := os.Getpid()
	fmt.Printf("Current PID: %d\n", pid)

	cmd := exec.Command("lsof", "-p", fmt.Sprintf("%d", pid))
	var output bytes.Buffer
	cmd.Stdout = &output

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error running lsof: %v\n", err)
		return
	}

	// 过滤 TCP 连接
	lines := bytes.Split(output.Bytes(), []byte("\n"))
	fmt.Println("TCP Connections:")
	for _, line := range lines {
		if bytes.Contains(line, []byte("TCP")) {
			fmt.Println(string(line))
		}
	}
}

func TestTcpLeak(t *testing.T) {
	key := os.Getenv("API_KEY")
	secret := os.Getenv("API_SECRET")
	passphrase := os.Getenv("API_PASSPHRASE")

	defaultLogger := logger.NewDefaultLogger()
	logger.SetLogger(defaultLogger)

	httpOptionBuilder := types.NewTransportOptionBuilder().
		SetKeepAlive(true).
		AddInterceptors(interceptor.NewLoggingInterceptor(false, defaultLogger))

	websocketOptions := types.NewWebSocketClientOptionBuilder().WithReconnect(true).Build()

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
	futuresWsApi := client.WsService().NewFuturesPrivateWS()
	marginWsApi := client.WsService().NewMarginPublicWS()
	err := spotWsApi.Start()
	assert.Nil(t, err)
	err = futuresWsApi.Start()
	assert.Nil(t, err)
	err = marginWsApi.Start()
	assert.Nil(t, err)

	GetTcpStat()
	time.Sleep(5 * time.Second)
	spotWsApi.Stop()
	futuresWsApi.Stop()
	marginWsApi.Stop()
	time.Sleep(5 * time.Second)
	GetTcpStat()
}
