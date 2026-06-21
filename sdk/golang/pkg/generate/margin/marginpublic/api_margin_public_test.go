package marginpublic

import (
	"encoding/json"
	"testing"

	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestMarginPublicIndexPriceRespModel(t *testing.T) {
	// IndexPrice
	// Index Price

	data := "{\"id\":\"5c24c5da03aa673885cd67a0\",\"type\":\"message\",\"topic\":\"/indicator/index:USDT-BTC\",\"subject\":\"tick\",\"data\":{\"symbol\":\"USDT-BTC\",\"granularity\":5000,\"timestamp\":1551770400000,\"value\":0.0001092}}"

	commonResp := &types.WsMessage{}
	err := json.Unmarshal([]byte(data), commonResp)
	assert.Nil(t, err)
	assert.NotNil(t, commonResp.RawData)
	obj := &IndexPriceEvent{}
	err = json.Unmarshal(commonResp.RawData, obj)
	assert.Nil(t, err)
}

func TestMarginPublicMarkPriceRespModel(t *testing.T) {
	// MarkPrice
	// Mark Price

	data := "{\"id\":\"5c24c5da03aa673885cd67aa\",\"type\":\"message\",\"topic\":\"/indicator/markPrice:USDT-BTC\",\"subject\":\"tick\",\"data\":{\"symbol\":\"USDT-BTC\",\"granularity\":5000,\"timestamp\":1551770400000,\"value\":0.0001093}}"

	commonResp := &types.WsMessage{}
	err := json.Unmarshal([]byte(data), commonResp)
	assert.Nil(t, err)
	assert.NotNil(t, commonResp.RawData)
	obj := &MarkPriceEvent{}
	err = json.Unmarshal(commonResp.RawData, obj)
	assert.Nil(t, err)
}
