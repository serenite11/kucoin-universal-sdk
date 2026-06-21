package infra

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/internal/interfaces"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/internal/util"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/common/logger"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/types"
)

type DefaultWsService struct {
	client       interfaces.WebsocketTransport
	topicManager *TopicManager
	option       *types.WebSocketClientOption
	stopChan     chan struct{}
	private      bool
}

func NewDefaultWsService(option *types.ClientOption, domain string, private bool, sdkVersion string) interfaces.WebSocketService {
	tokenTransport := NewDefaultTransport(option, sdkVersion)
	wsOption := option.WebSocketClientOption

	client := NewWebSocketClient(NewDefaultWsTokenProvider(tokenTransport, domain, private), wsOption)

	return &DefaultWsService{
		client:       client,
		topicManager: NewTopicManager(),
		option:       wsOption,
		stopChan:     make(chan struct{}),
		private:      private,
	}
}

func (ws *DefaultWsService) recovery() {
	go func() {
		for {
			select {
			case <-ws.stopChan:
				return
			case <-ws.client.Reconnected():
				logger.GetLogger().Infof("websocket client reconnected, resubscribe...")
				oldTopicManager := ws.topicManager
				ws.topicManager = NewTopicManager()
				oldTopicManager.Range(func(key, value any) bool {
					cm := value.(*CallbackManager)
					subInfo := cm.GetSubInfo()
					for _, sub := range subInfo {
						subId, err := ws.Subscribe(sub.Prefix, sub.Args, sub.Callback)
						if err != nil {
							ws.notifyEvent(types.EventReSubscribeError, fmt.Sprintf("id: %s, err: %s", subId, err.Error()))
						} else {
							ws.notifyEvent(types.EventReSubscribeOK, subId)
						}
					}
					return true
				})
			}
		}
	}()
}

func (ws *DefaultWsService) Start() error {
	err := ws.client.Start()
	if err != nil {
		return err
	}

	ws.run()
	ws.recovery()
	return nil
}

func (c *DefaultWsService) notifyEvent(event types.WebSocketEvent, msg string) {
	defer func() {
		if r := recover(); r != nil {
			logger.GetLogger().Errorf("recovered in notifyEvent: %v", r)
		}
	}()

	if c.option.EventCallback != nil {
		c.option.EventCallback(event, msg)
	}
}

func (ws *DefaultWsService) run() {
	go func() {
		for {
			select {
			case <-ws.stopChan:
				return
			case msg, ok := <-ws.client.Read():
				if !ok {
					continue
				}

				if msg.Type != types.Message {
					continue
				}

				callbackManager := ws.topicManager.GetCallbackManager(msg.Topic)
				cb := callbackManager.Get(msg.Topic)
				if cb == nil {
					logger.GetLogger().Errorf("can not find callback manager, topic: %s", msg.Topic)
					continue
				}

				// process user callback
				func() {
					defer func() {
						if r := recover(); r != nil {
							logger.GetLogger().Errorf("recovered in callback: %v", r)
						}
					}()
					err := cb.OnMessage(msg)
					if err != nil {
						ws.notifyEvent(types.EventCallbackError, err.Error())
						logger.GetLogger().Errorf("callback error: %s", err.Error())
					}
				}()

			}
		}
	}()
}

func (ws *DefaultWsService) Stop() error {
	defer close(ws.stopChan)
	logger.GetLogger().Infof("closing websocket client")
	return ws.client.Stop()
}

func (ws *DefaultWsService) Subscribe(prefix string, args []string, callback interfaces.WebSocketMessageCallback) (id string, err error) {
	defer func() {
		logger.GetLogger().Infof("subscribe prefix:[%s], args:%s, private:[%v], id:[%s], err:[%v]", prefix, args, ws.private, id, err)
	}()

	if args == nil {
		args = []string{}
	}

	subInfo := &util.SubInfo{Prefix: prefix, Args: args, Callback: callback}
	id = subInfo.ToId()

	callbackManager := ws.topicManager.GetCallbackManager(prefix)
	created := callbackManager.Add(subInfo)
	if !created {
		// already subscribed
		return id, fmt.Errorf("already subscribed")
	}

	defer func() {
		if err != nil {
			callbackManager.Remove(id)
		}
	}()

	subEvent := types.WsMessage{}
	subEvent.Id = id
	subEvent.Type = types.SubscribeMessage
	subEvent.Topic = subInfo.SubTopic()
	subEvent.PrivateChannel = ws.private
	subEvent.Response = true

	ctx, cancel := context.WithTimeout(context.Background(), ws.option.WriteTimeout)
	defer cancel()

	writeSignal := ws.client.Write(ctx, &subEvent)

	// wait ack or timeout
	select {
	case err = <-writeSignal:
		return id, err
	case <-ctx.Done():
		return id, ctx.Err()
	}
}

func (ws *DefaultWsService) Unsubscribe(id string) (err error) {
	defer func() {
		logger.GetLogger().Infof("unsubscribe id:[%s], err:[%v]", id, err)
	}()

	subInfo := &util.SubInfo{}
	err = subInfo.FromId(id)
	if err != nil {
		return err
	}
	callbackManager := ws.topicManager.GetCallbackManager(subInfo.SubTopic())

	subEvent := types.WsMessage{}
	subEvent.Id = uuid.NewString()
	subEvent.Type = types.UnsubscribeMessage
	subEvent.Topic = subInfo.SubTopic()
	subEvent.Response = true
	subEvent.PrivateChannel = ws.private

	ctx, cancel := context.WithTimeout(context.Background(), ws.option.WriteTimeout)
	defer cancel()

	writeSignal := ws.client.Write(ctx, &subEvent)

	// wait ack or timeout
	select {
	case err := <-writeSignal:
		{
			if err == nil {
				// write success, remove topicCallbackMapping
				callbackManager.Remove(id)
			}
			return err
		}
	case <-ctx.Done():
		return ctx.Err()
	}

}
