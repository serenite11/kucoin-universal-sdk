package interfaces

import (
	"context"

	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/types"
)

// WsToken holds the token and API endpoint for a WebSocket connection
type WsToken struct {
	Token        string `json:"token"`
	PingInterval int64  `json:"pingInterval"`
	Endpoint     string `json:"endpoint"`
	Protocol     string `json:"protocol"`
	Encrypt      bool   `json:"encrypt"`
	PingTimeout  int64  `json:"pingTimeout"`
}

// WsTokenProvider defines a method for retrieving a WebSocket token
type WsTokenProvider interface {
	GetToken() (error, []*WsToken)

	Close() error
}

// WebSocketMessageCallback defines a method for handling WebSocket messages
type WebSocketMessageCallback interface {
	// OnMessage handles the incoming WebSocket messages
	OnMessage(message *types.WsMessage) error
}

// WebSocketService defines methods for subscribing to
// and unsubscribing from topics in a WebSocket connection
type WebSocketService interface {
	// Start starts the WebSocket service and handles incoming WebSocket messages.
	Start() error
	// Stop stops the WebSocket service
	Stop() error
	// Subscribe to a topic with a provided message callback
	Subscribe(topic string, args []string, callback WebSocketMessageCallback) (string, error)
	// Unsubscribe from a topic
	Unsubscribe(id string) error
}

// WebsocketTransport defines methods required for managing a WebSocket connection.
// This includes connecting to the WebSocket, closing the connection,
// writing messages, and reading from the connection.
type WebsocketTransport interface {
	Start() error

	Stop() error

	Write(context.Context, *types.WsMessage) <-chan error

	Read() <-chan *types.WsMessage

	Reconnected() <-chan struct{}
}
