package futurespublic

import (
	"encoding/json"
	"testing"

	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestFuturesPublicAnnouncementRespModel(t *testing.T) {
	// Announcement
	// announcement

	data := "{\"topic\":\"/contract/announcement\",\"subject\":\"funding.begin\",\"data\":{\"symbol\":\"XBTUSDTM\",\"fundingTime\":1551770400000,\"fundingRate\":-0.002966,\"timestamp\":1551770400000}}"

	commonResp := &types.WsMessage{}
	err := json.Unmarshal([]byte(data), commonResp)
	assert.Nil(t, err)
	assert.NotNil(t, commonResp.RawData)
	obj := &AnnouncementEvent{}
	err = json.Unmarshal(commonResp.RawData, obj)
	assert.Nil(t, err)
}

func TestFuturesPublicExecutionRespModel(t *testing.T) {
	// Execution
	// Match execution data.

	data := "{\"topic\":\"/contractMarket/execution:XBTUSDTM\",\"type\":\"message\",\"subject\":\"match\",\"sn\":1794100537695,\"data\":{\"symbol\":\"XBTUSDTM\",\"sequence\":1794100537695,\"side\":\"buy\",\"size\":2,\"price\":\"90503.9\",\"takerOrderId\":\"247822202957807616\",\"makerOrderId\":\"247822167163555840\",\"tradeId\":\"1794100537695\",\"ts\":1731898619520000000}}"

	commonResp := &types.WsMessage{}
	err := json.Unmarshal([]byte(data), commonResp)
	assert.Nil(t, err)
	assert.NotNil(t, commonResp.RawData)
	obj := &ExecutionEvent{}
	err = json.Unmarshal(commonResp.RawData, obj)
	assert.Nil(t, err)
}

func TestFuturesPublicInstrumentRespModel(t *testing.T) {
	// Instrument
	// instrument

	data := "{\"topic\":\"/contract/instrument:XBTUSDTM\",\"type\":\"message\",\"subject\":\"mark.index.price\",\"data\":{\"markPrice\":90445.02,\"indexPrice\":90445.02,\"granularity\":1000,\"timestamp\":1731899129000}}"

	commonResp := &types.WsMessage{}
	err := json.Unmarshal([]byte(data), commonResp)
	assert.Nil(t, err)
	assert.NotNil(t, commonResp.RawData)
	obj := &InstrumentEvent{}
	err = json.Unmarshal(commonResp.RawData, obj)
	assert.Nil(t, err)
}

func TestFuturesPublicKlinesRespModel(t *testing.T) {
	// Klines
	// Klines

	data := "{\"topic\":\"/contractMarket/limitCandle:XBTUSDTM_1min\",\"type\":\"message\",\"data\":{\"symbol\":\"XBTUSDTM\",\"candles\":[\"1731898200\",\"90638.6\",\"90638.6\",\"90638.6\",\"90638.6\",\"21.0\",\"21\"],\"time\":1731898208357},\"subject\":\"candle.stick\"}"

	commonResp := &types.WsMessage{}
	err := json.Unmarshal([]byte(data), commonResp)
	assert.Nil(t, err)
	assert.NotNil(t, commonResp.RawData)
	obj := &KlinesEvent{}
	err = json.Unmarshal(commonResp.RawData, obj)
	assert.Nil(t, err)
}

func TestFuturesPublicOrderbookIncrementRespModel(t *testing.T) {
	// OrderbookIncrement
	// Orderbook - Increment

	data := "{\"topic\":\"/contractMarket/level2:XBTUSDTM\",\"type\":\"message\",\"subject\":\"level2\",\"sn\":1709400450243,\"data\":{\"sequence\":1709400450243,\"change\":\"90631.2,sell,2\",\"timestamp\":1731897467182}}"

	commonResp := &types.WsMessage{}
	err := json.Unmarshal([]byte(data), commonResp)
	assert.Nil(t, err)
	assert.NotNil(t, commonResp.RawData)
	obj := &OrderbookIncrementEvent{}
	err = json.Unmarshal(commonResp.RawData, obj)
	assert.Nil(t, err)
}

func TestFuturesPublicOrderbookLevel50RespModel(t *testing.T) {
	// OrderbookLevel50
	// Orderbook - Level50

	data := "{\"topic\":\"/contractMarket/level2Depth50:XBTUSDTM\",\"type\":\"message\",\"subject\":\"level2\",\"sn\":1731680249700,\"data\":{\"bids\":[[\"89778.6\",1534],[\"89778.2\",54]],\"sequence\":1709294490099,\"timestamp\":1731680249700,\"ts\":1731680249700,\"asks\":[[\"89778.7\",854],[\"89779.2\",4]]}}"

	commonResp := &types.WsMessage{}
	err := json.Unmarshal([]byte(data), commonResp)
	assert.Nil(t, err)
	assert.NotNil(t, commonResp.RawData)
	obj := &OrderbookLevel50Event{}
	err = json.Unmarshal(commonResp.RawData, obj)
	assert.Nil(t, err)
}

func TestFuturesPublicOrderbookLevel5RespModel(t *testing.T) {
	// OrderbookLevel5
	// Orderbook - Level5

	data := "{\"topic\":\"/contractMarket/level2Depth5:XBTUSDTM\",\"type\":\"message\",\"subject\":\"level2\",\"sn\":1731680019100,\"data\":{\"bids\":[[\"89720.9\",513],[\"89720.8\",12],[\"89718.6\",113],[\"89718.4\",19],[\"89718.3\",7]],\"sequence\":1709294294670,\"timestamp\":1731680019100,\"ts\":1731680019100,\"asks\":[[\"89721\",906],[\"89721.1\",203],[\"89721.4\",113],[\"89723.2\",113],[\"89725.4\",113]]}}"

	commonResp := &types.WsMessage{}
	err := json.Unmarshal([]byte(data), commonResp)
	assert.Nil(t, err)
	assert.NotNil(t, commonResp.RawData)
	obj := &OrderbookLevel5Event{}
	err = json.Unmarshal(commonResp.RawData, obj)
	assert.Nil(t, err)
}

func TestFuturesPublicSymbolSnapshotRespModel(t *testing.T) {
	// SymbolSnapshot
	// Symbol Snapshot

	data := "{\"topic\":\"/contractMarket/snapshot:XBTUSDTM\",\"type\":\"message\",\"subject\":\"snapshot.24h\",\"id\":\"673ab3fff4088b0001664f41\",\"data\":{\"highPrice\":91512.8,\"lastPrice\":90326.7,\"lowPrice\":88747.8,\"price24HoursBefore\":89880.4,\"priceChg\":446.3,\"priceChgPct\":0.0049,\"symbol\":\"XBTUSDTM\",\"ts\":1731900415023929239,\"turnover\":526928331.0482177734,\"volume\":5834.46}}"

	commonResp := &types.WsMessage{}
	err := json.Unmarshal([]byte(data), commonResp)
	assert.Nil(t, err)
	assert.NotNil(t, commonResp.RawData)
	obj := &SymbolSnapshotEvent{}
	err = json.Unmarshal(commonResp.RawData, obj)
	assert.Nil(t, err)
}

func TestFuturesPublicTickerV1RespModel(t *testing.T) {
	// TickerV1
	// Get Ticker(not recommended)

	data := "{\"topic\":\"/contractMarket/ticker:XBTUSDTM\",\"type\":\"message\",\"subject\":\"ticker\",\"sn\":1793133570931,\"data\":{\"symbol\":\"XBTUSDTM\",\"sequence\":1793133570931,\"side\":\"sell\",\"size\":1,\"price\":\"90147.7\",\"bestBidSize\":2186,\"bestBidPrice\":\"90147.7\",\"bestAskPrice\":\"90147.8\",\"tradeId\":\"1793133570931\",\"bestAskSize\":275,\"ts\":1731679215637000000}}"

	commonResp := &types.WsMessage{}
	err := json.Unmarshal([]byte(data), commonResp)
	assert.Nil(t, err)
	assert.NotNil(t, commonResp.RawData)
	obj := &TickerV1Event{}
	err = json.Unmarshal(commonResp.RawData, obj)
	assert.Nil(t, err)
}

func TestFuturesPublicTickerV2RespModel(t *testing.T) {
	// TickerV2
	// Get Ticker V2

	data := "{\"topic\":\"/contractMarket/tickerV2:XBTUSDTM\",\"type\":\"message\",\"subject\":\"tickerV2\",\"sn\":1709284589209,\"data\":{\"symbol\":\"XBTUSDTM\",\"sequence\":1709284589209,\"bestBidSize\":713,\"bestBidPrice\":\"88987.4\",\"bestAskPrice\":\"88987.5\",\"bestAskSize\":1037,\"ts\":1731665526461000000}}"

	commonResp := &types.WsMessage{}
	err := json.Unmarshal([]byte(data), commonResp)
	assert.Nil(t, err)
	assert.NotNil(t, commonResp.RawData)
	obj := &TickerV2Event{}
	err = json.Unmarshal(commonResp.RawData, obj)
	assert.Nil(t, err)
}
