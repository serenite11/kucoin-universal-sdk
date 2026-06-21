package infra

import (
	"context"
	"time"

	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/internal/interfaces"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/types"
)

const (
	PathPrivate = "/api/v1/bullet-private"
	PathPublic  = "/api/v1/bullet-public"
)

type DefaultWsTokenProvider struct {
	transport interfaces.Transport
	domain    string
	private   bool
}

func NewDefaultWsTokenProvider(transport interfaces.Transport, domain string, private bool) *DefaultWsTokenProvider {
	return &DefaultWsTokenProvider{
		transport: transport,
		domain:    domain,
		private:   private,
	}
}

func (p *DefaultWsTokenProvider) GetToken() (error, []*interfaces.WsToken) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	path := PathPublic
	if p.private {
		path = PathPrivate
	}

	resp := &tokenResponse{}
	err := p.transport.Call(ctx, p.domain, false, "Post", path, nil, resp, false)
	if err != nil {
		return err, nil
	}

	for _, sr := range resp.Servers {
		sr.Token = resp.Token
	}

	return nil, resp.Servers

}

func (p *DefaultWsTokenProvider) Close() error {
	p.transport.Close()
	return nil
}

type tokenResponse struct {
	commonResponse *types.RestResponse
	Token          string                `json:"token"`
	Servers        []*interfaces.WsToken `json:"instanceServers"`
}

func (m *tokenResponse) SetCommonResponse(response *types.RestResponse) {
	m.commonResponse = response
}
