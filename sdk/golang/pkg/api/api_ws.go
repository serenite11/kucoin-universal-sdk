package api

import (
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/generate/futures/futuresprivate"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/generate/futures/futurespublic"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/generate/margin/marginprivate"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/generate/margin/marginpublic"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/generate/spot/spotprivate"
	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/generate/spot/spotpublic"
)

type KucoinWSService interface {
	// NewSpotPublicWS returns the interface to interact with
	// the Spot Trading websocket(public channel) API of Kucoin.
	NewSpotPublicWS() spotpublic.SpotPublicWS

	// NewSpotPrivateWS returns the interface to interact with
	// the Spot Trading websocket(private channel) API of Kucoin.
	NewSpotPrivateWS() spotprivate.SpotPrivateWS

	// NewMarginPublicWS returns the interface to interact with
	// the Margin Trading websocket(public channel) API of Kucoin.
	NewMarginPublicWS() marginpublic.MarginPublicWS

	// NewMarginPrivateWS returns the interface to interact with
	// the Margin Trading websocket(private channel) API of Kucoin.
	NewMarginPrivateWS() marginprivate.MarginPrivateWS

	// NewFuturesPublicWS returns the interface to interact with
	// the Futures Trading websocket(public channel) API of Kucoin.
	NewFuturesPublicWS() futurespublic.FuturesPublicWS

	// NewFuturesPrivateWS returns the interface to interact with
	// the Futures Trading websocket(private channel) API of Kucoin.
	NewFuturesPrivateWS() futuresprivate.FuturesPrivateWS
}
