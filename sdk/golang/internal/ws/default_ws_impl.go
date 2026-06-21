package ws

import (
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/internal/infra"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/common/logger"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/generate"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/generate/futures/futuresprivate"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/generate/futures/futurespublic"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/generate/margin/marginprivate"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/generate/margin/marginpublic"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/generate/spot/spotprivate"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/generate/spot/spotpublic"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/types"
)

type KuCoinDefaultWsImpl struct {
	option *types.ClientOption
}

func NewKuCoinDefaultWsImpl(op *types.ClientOption) *KuCoinDefaultWsImpl {
	if op == nil || op.WebSocketClientOption == nil {
		logger.GetLogger().Warnf("no websocket option provided")
		return nil
	}

	return &KuCoinDefaultWsImpl{
		op,
	}
}

func (impl *KuCoinDefaultWsImpl) NewSpotPublicWS() spotpublic.SpotPublicWS {
	return spotpublic.NewSpotPublicWSImp(infra.NewDefaultWsService(impl.option, types.DomainTypeSpot, false, generate.SdkVersion))
}

func (impl *KuCoinDefaultWsImpl) NewSpotPrivateWS() spotprivate.SpotPrivateWS {
	return spotprivate.NewSpotPrivateWSImp(infra.NewDefaultWsService(impl.option, types.DomainTypeSpot, true, generate.SdkVersion))
}

func (impl *KuCoinDefaultWsImpl) NewFuturesPublicWS() futurespublic.FuturesPublicWS {
	return futurespublic.NewFuturesPublicWSImp(infra.NewDefaultWsService(impl.option, types.DomainTypeFutures, false, generate.SdkVersion))
}

func (impl *KuCoinDefaultWsImpl) NewFuturesPrivateWS() futuresprivate.FuturesPrivateWS {
	return futuresprivate.NewFuturesPrivateWSImp(infra.NewDefaultWsService(impl.option, types.DomainTypeFutures, true, generate.SdkVersion))
}

func (impl *KuCoinDefaultWsImpl) NewMarginPublicWS() marginpublic.MarginPublicWS {
	return marginpublic.NewMarginPublicWSImp(infra.NewDefaultWsService(impl.option, types.DomainTypeSpot, false, generate.SdkVersion))

}

func (impl *KuCoinDefaultWsImpl) NewMarginPrivateWS() marginprivate.MarginPrivateWS {
	return marginprivate.NewMarginPrivateWSImp(infra.NewDefaultWsService(impl.option, types.DomainTypeSpot, true, generate.SdkVersion))
}
