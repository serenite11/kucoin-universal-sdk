package infra

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/internal/interfaces"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/common/logger"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/types"
	"github.com/stretchr/testify/assert"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func mockWebSocketServer(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	err = conn.WriteJSON(&types.WsMessage{
		Id:   IntToString(time.Now().UnixNano()),
		Type: types.WelcomeMessage,
	})
	if err != nil {
		panic(err)
	}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			break
		}

		m := &types.WsMessage{}

		err = json.Unmarshal(message, m)
		if err != nil {
			panic(err)
		}

		switch m.Type {
		case types.PingMessage:
			{
				err = conn.WriteJSON(&types.WsMessage{
					Id:   m.Id,
					Type: types.PongMessage,
				})
				if err != nil {
					panic(err)
				}
			}
		case types.SubscribeMessage:
			{
				logger.GetLogger().Infof("[server]receive subscribe message")
				err = conn.WriteJSON(&types.WsMessage{
					Id:   m.Id,
					Type: types.AckMessage,
				})
				if err != nil {
					panic(err)
				}
				time.Sleep(time.Millisecond * 10)

				for i := 0; i < 10; i++ {
					err = conn.WriteJSON(&types.WsMessage{
						Id:   IntToString(int64(i)),
						Type: types.Message,
					})
					if err != nil {
						panic(err)
					}
					time.Sleep(time.Millisecond * 10)
				}
			}

		}
	}
}

func mockWebSocketServer2(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	err = conn.WriteJSON(&types.WsMessage{
		Id:   IntToString(time.Now().UnixNano()),
		Type: types.WelcomeMessage,
	})
	if err != nil {
		panic(err)
	}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			break
		}

		m := &types.WsMessage{}

		err = json.Unmarshal(message, m)
		if err != nil {
			panic(err)
		}

		switch m.Type {
		case types.PingMessage:
			{
				err = conn.WriteJSON(&types.WsMessage{
					Id:   m.Id,
					Type: types.PongMessage,
				})
				if err != nil {
					panic(err)
				}
			}
		case types.SubscribeMessage:
			{
				logger.GetLogger().Infof("[server]receive subscribe message")
				return
			}

		}
	}
}

type tokenProviderMock struct {
	endpoint     string
	PingInterval int64
}

func (m *tokenProviderMock) GetToken() (error, []*interfaces.WsToken) {
	if m.PingInterval == 0 {
		m.PingInterval = 10
	}
	return nil, []*interfaces.WsToken{&interfaces.WsToken{
		Token:        "token-test",
		PingInterval: m.PingInterval,
		Endpoint:     m.endpoint,
		Protocol:     "",
		Encrypt:      false,
		PingTimeout:  200,
	}}
}

func (m *tokenProviderMock) Close() error {
	return nil
}

func TestWebSocketClient(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(mockWebSocketServer))
	defer server.Close()

	serverURL := "ws" + server.URL[4:]

	tp := &tokenProviderMock{endpoint: serverURL}

	eventCnt := 0
	wsOp := types.NewWebSocketClientOptionBuilder().WithEventCallback(func(event types.WebSocketEvent, msg string) {
		if event == types.EventConnected || event == types.EventDisconnected {
			eventCnt++
		}
	}).Build()

	client := NewWebSocketClient(tp, wsOp)
	err := client.Start()
	assert.Nil(t, err)
	<-time.After(time.Millisecond * 20)
	assert.Equal(t, int64(3), client.metric.goroutines)
	err = client.Start()
	assert.NotNil(t, err)
	assert.Equal(t, int64(3), client.metric.goroutines)

	assert.Equal(t, 1, eventCnt)

	client.Stop()
	client.Stop()
	assert.Equal(t, int64(0), client.metric.goroutines)
	assert.Equal(t, 2, eventCnt)
	assert.Equal(t, 0, len(client.ackEvent))
}

func TestWebSocketClient2(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(mockWebSocketServer))
	defer server.Close()

	serverURL := "ws" + server.URL[4:]

	tp := &tokenProviderMock{endpoint: serverURL}

	eventCnt := 0
	wsOp := types.NewWebSocketClientOptionBuilder().WithEventCallback(func(event types.WebSocketEvent, msg string) {
		if event == types.EventConnected || event == types.EventDisconnected {
			eventCnt++
		}
	}).Build()

	client := NewWebSocketClient(tp, wsOp)

	for i := 0; i < 10; i++ {
		go func() {
			client.Start()
		}()
	}
	<-time.After(time.Millisecond * 20)
	assert.Equal(t, int64(3), client.metric.goroutines)
	client.Stop()
	assert.Equal(t, int64(0), client.metric.goroutines)
	assert.Equal(t, 2, eventCnt)
	assert.Equal(t, 0, len(client.ackEvent))
}

func TestWebSocketClient_ping(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(mockWebSocketServer))
	defer server.Close()

	serverURL := "ws" + server.URL[4:]

	tp := &tokenProviderMock{endpoint: serverURL}

	wsOp := types.NewWebSocketClientOptionBuilder().Build()

	client := NewWebSocketClient(tp, wsOp)
	err := client.Start()
	assert.Nil(t, err)

	<-time.After(time.Millisecond * 200)

	assert.Equal(t, int64(0), client.metric.pingErr)
	assert.True(t, client.metric.pingSuccess > 0)

	assert.Nil(t, client.Stop())
	assert.Equal(t, 0, len(client.ackEvent))
	assert.Equal(t, int64(0), client.metric.goroutines)

}

func TestWebSocketClient_write_read(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(mockWebSocketServer))
	defer server.Close()

	serverURL := "ws" + server.URL[4:]

	tp := &tokenProviderMock{endpoint: serverURL}

	wsOp := types.NewWebSocketClientOptionBuilder().Build()

	client := NewWebSocketClient(tp, wsOp)
	err := client.Start()
	assert.Nil(t, err)

	ctx, fc := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer fc()

	ch := client.Write(ctx, &types.WsMessage{Id: "id123", Type: types.SubscribeMessage})
	assert.Nil(t, <-ch)

	cnt := 0
	go func() {
		for m := range client.Read() {
			fmt.Println(m)
			cnt++
		}
	}()

	<-time.After(time.Millisecond * 200)
	assert.Nil(t, client.Stop())
	assert.Equal(t, 10, cnt)
	assert.Equal(t, 0, len(client.ackEvent))
	assert.Equal(t, int64(0), client.metric.goroutines)

}

func TestWebSocketClient_reconnect(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(mockWebSocketServer2))
	defer server.Close()

	serverURL := "ws" + server.URL[4:]

	tp := &tokenProviderMock{endpoint: serverURL, PingInterval: 100}

	wsOp := types.NewWebSocketClientOptionBuilder().WithEventCallback(func(event types.WebSocketEvent, msg string) {
		fmt.Println(event)
	}).WithReconnectInterval(time.Millisecond * 5).Build()

	client := NewWebSocketClient(tp, wsOp)
	assert.Nil(t, client.Start())

	<-time.After(time.Millisecond * 200)

	ctx, fc := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer fc()
	ch := client.Write(ctx, &types.WsMessage{Id: "id123", Type: types.SubscribeMessage})
	<-ch

	<-time.After(time.Millisecond * 200)

	assert.True(t, client.metric.pingSuccess > 0)
	assert.Equal(t, int64(3), client.metric.goroutines)

	client.Stop()
	assert.Equal(t, 0, len(client.ackEvent))
	assert.Equal(t, int64(0), client.metric.goroutines)

}

func TestWebSocketClient_reconnect2(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(mockWebSocketServer2))
	defer server.Close()

	serverURL := "ws" + server.URL[4:]

	tp := &tokenProviderMock{endpoint: serverURL, PingInterval: 200}

	var ev types.WebSocketEvent
	wsOp := types.NewWebSocketClientOptionBuilder().WithReconnectInterval(time.Millisecond * 5).WithEventCallback(func(event types.WebSocketEvent, msg string) {
		fmt.Println(event)
		ev = event
	}).WithReconnect(false).Build()

	client := NewWebSocketClient(tp, wsOp)
	assert.Nil(t, client.Start())

	<-time.After(time.Millisecond * 200)

	ctx, fc := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer fc()
	ch := client.Write(ctx, &types.WsMessage{Id: "id123", Type: types.SubscribeMessage})
	<-ch

	<-time.After(time.Millisecond * 200)
	assert.Equal(t, int64(0), client.metric.goroutines)
	assert.Equal(t, types.EventClientFail, ev)
}

func TestWebSocketClient_reconnect3(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(mockWebSocketServer2))
	defer server.Close()

	serverURL := "ws" + server.URL[4:]

	tp := &tokenProviderMock{endpoint: serverURL}
	tp.PingInterval = 50

	wsOp := types.NewWebSocketClientOptionBuilder().WithReconnectInterval(time.Millisecond * 5).
		WithEventCallback(func(event types.WebSocketEvent, msg string) {
			fmt.Println(event, msg)
		}).Build()

	client := NewWebSocketClient(tp, wsOp)
	assert.Nil(t, client.Start())

	<-time.After(time.Millisecond * 200)

	ctx, fc := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer fc()
	ch := client.Write(ctx, &types.WsMessage{Id: "id123", Type: types.SubscribeMessage})
	<-ch

	<-time.After(time.Millisecond * 200)

	assert.True(t, client.metric.pingSuccess > 0)
	assert.Equal(t, int64(3), client.metric.goroutines)

	client.Stop()
	assert.Equal(t, 0, len(client.ackEvent))
	assert.Equal(t, int64(0), client.metric.goroutines)
}
