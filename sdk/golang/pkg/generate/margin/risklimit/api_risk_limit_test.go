package risklimit

import (
	"encoding/json"
	"testing"

	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestRiskLimitGetMarginRiskLimitReqModel(t *testing.T) {
	// GetMarginRiskLimit
	// Get Margin Risk Limit
	// /api/v3/margin/currencies

	data := "{\"isIsolated\": true, \"currency\": \"BTC\", \"symbol\": \"BTC-USDT\"}"
	req := &GetMarginRiskLimitReq{}
	err := json.Unmarshal([]byte(data), req)
	req.ToMap()
	assert.Nil(t, err)
}

func TestRiskLimitGetMarginRiskLimitRespModel(t *testing.T) {
	// GetMarginRiskLimit
	// Get Margin Risk Limit
	// /api/v3/margin/currencies

	data := "{\n    \"code\": \"200000\",\n    \"data\": [\n        {\n            \"timestamp\": 1729678659275,\n            \"currency\": \"BTC\",\n            \"borrowMaxAmount\": \"75.15\",\n            \"buyMaxAmount\": \"217.12\",\n            \"holdMaxAmount\": \"217.12\",\n            \"borrowCoefficient\": \"1\",\n            \"marginCoefficient\": \"1\",\n            \"precision\": 8,\n            \"borrowMinAmount\": \"0.001\",\n            \"borrowMinUnit\": \"0.0001\",\n            \"borrowEnabled\": true\n        }\n    ]\n}\n"
	commonResp := &types.RestResponse{}
	err := json.Unmarshal([]byte(data), commonResp)
	assert.Nil(t, err)
	assert.NotNil(t, commonResp.Data)
	resp := &GetMarginRiskLimitResp{}
	err = json.Unmarshal([]byte(commonResp.Data), resp)
	resp.ToMap()
	assert.Nil(t, err)
}
