package fee

import (
	"encoding/json"
	"testing"

	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestFeeGetBasicFeeReqModel(t *testing.T) {
	// GetBasicFee
	// Get Basic Fee - Spot/Margin
	// /api/v1/base-fee

	data := "{\"currencyType\": 1}"
	req := &GetBasicFeeReq{}
	err := json.Unmarshal([]byte(data), req)
	req.ToMap()
	assert.Nil(t, err)
}

func TestFeeGetBasicFeeRespModel(t *testing.T) {
	// GetBasicFee
	// Get Basic Fee - Spot/Margin
	// /api/v1/base-fee

	data := "{\n    \"code\": \"200000\",\n    \"data\": {\n        \"takerFeeRate\": \"0.001\",\n        \"makerFeeRate\": \"0.001\"\n    }\n}"
	commonResp := &types.RestResponse{}
	err := json.Unmarshal([]byte(data), commonResp)
	assert.Nil(t, err)
	assert.NotNil(t, commonResp.Data)
	resp := &GetBasicFeeResp{}
	err = json.Unmarshal([]byte(commonResp.Data), resp)
	resp.ToMap()
	assert.Nil(t, err)
}

func TestFeeGetSpotActualFeeReqModel(t *testing.T) {
	// GetSpotActualFee
	// Get Actual Fee - Spot/Margin
	// /api/v1/trade-fees

	data := "{\"symbols\": \"BTC-USDT,ETH-USDT\"}"
	req := &GetSpotActualFeeReq{}
	err := json.Unmarshal([]byte(data), req)
	req.ToMap()
	assert.Nil(t, err)
}

func TestFeeGetSpotActualFeeRespModel(t *testing.T) {
	// GetSpotActualFee
	// Get Actual Fee - Spot/Margin
	// /api/v1/trade-fees

	data := "{\"code\":\"200000\",\"data\":[{\"symbol\":\"BTC-USDT\",\"takerFeeRate\":\"0.001\",\"makerFeeRate\":\"0.001\"},{\"symbol\":\"ETH-USDT\",\"takerFeeRate\":\"0.001\",\"makerFeeRate\":\"0.001\"}]}"
	commonResp := &types.RestResponse{}
	err := json.Unmarshal([]byte(data), commonResp)
	assert.Nil(t, err)
	assert.NotNil(t, commonResp.Data)
	resp := &GetSpotActualFeeResp{}
	err = json.Unmarshal([]byte(commonResp.Data), resp)
	resp.ToMap()
	assert.Nil(t, err)
}

func TestFeeGetFuturesActualFeeReqModel(t *testing.T) {
	// GetFuturesActualFee
	// Get Actual Fee - Futures
	// /api/v1/trade-fees

	data := "{\"symbol\": \"XBTUSDTM\"}"
	req := &GetFuturesActualFeeReq{}
	err := json.Unmarshal([]byte(data), req)
	req.ToMap()
	assert.Nil(t, err)
}

func TestFeeGetFuturesActualFeeRespModel(t *testing.T) {
	// GetFuturesActualFee
	// Get Actual Fee - Futures
	// /api/v1/trade-fees

	data := "{\"code\":\"200000\",\"data\":{\"symbol\":\"XBTUSDTM\",\"takerFeeRate\":\"0.0006\",\"makerFeeRate\":\"0.0002\"}}"
	commonResp := &types.RestResponse{}
	err := json.Unmarshal([]byte(data), commonResp)
	assert.Nil(t, err)
	assert.NotNil(t, commonResp.Data)
	resp := &GetFuturesActualFeeResp{}
	err = json.Unmarshal([]byte(commonResp.Data), resp)
	resp.ToMap()
	assert.Nil(t, err)
}
