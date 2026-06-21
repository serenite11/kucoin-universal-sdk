package affiliate

import (
	"encoding/json"
	"testing"

	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestAffiliateGetAccountReqModel(t *testing.T) {
	// GetAccount
	// Get Account
	// /api/v2/affiliate/inviter/statistics

}

func TestAffiliateGetAccountRespModel(t *testing.T) {
	// GetAccount
	// Get Account
	// /api/v2/affiliate/inviter/statistics

	data := "{\n    \"code\": \"200000\",\n    \"data\": {\n        \"parentUid\": \"1000000\",\n        \"orders\": [\n            {\n                \"orderId\": \"1668458892612980737\",\n                \"currency\": \"USDT\",\n                \"principal\": \"100\",\n                \"interest\": \"0\"\n            }\n        ],\n        \"ltv\": {\n            \"transferLtv\": \"0.6000\",\n            \"onlyClosePosLtv\": \"0.7500\",\n            \"delayedLiquidationLtv\": \"0.9000\",\n            \"instantLiquidationLtv\": \"0.9500\",\n            \"currentLtv\": \"0.0854\"\n        },\n        \"totalMarginAmount\": \"1170.36181573\",\n        \"transferMarginAmount\": \"166.66666666\",\n        \"margins\": [\n            {\n                \"marginCcy\": \"USDT\",\n                \"marginQty\": \"1170.36181573\",\n                \"marginFactor\": \"1.000000000000000000\"\n            }\n        ]\n    }\n}"
	commonResp := &types.RestResponse{}
	err := json.Unmarshal([]byte(data), commonResp)
	assert.Nil(t, err)
	assert.NotNil(t, commonResp.Data)
	resp := &GetAccountResp{}
	err = json.Unmarshal([]byte(commonResp.Data), resp)
	resp.ToMap()
	assert.Nil(t, err)
}
