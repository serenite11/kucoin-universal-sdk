package apibroker

import (
	"encoding/json"
	"testing"

	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestAPIBrokerGetRebaseReqModel(t *testing.T) {
	// GetRebase
	// Get Broker Rebate
	// /api/v1/broker/api/rebase/download

	data := "{\"begin\": \"20240610\", \"end\": \"20241010\", \"tradeType\": \"1\"}"
	req := &GetRebaseReq{}
	err := json.Unmarshal([]byte(data), req)
	req.ToMap()
	assert.Nil(t, err)
}

func TestAPIBrokerGetRebaseRespModel(t *testing.T) {
	// GetRebase
	// Get Broker Rebate
	// /api/v1/broker/api/rebase/download

	data := "{\n    \"code\": \"200000\",\n    \"data\": {\n        \"url\": \"https://kc-v2-promotion.s3.ap-northeast-1.amazonaws.com/broker/671aec522593f600019766d0_file.csv?X-Amz-Security-Token=IQo*********2cd90f14efb\"\n    }\n}"
	commonResp := &types.RestResponse{}
	err := json.Unmarshal([]byte(data), commonResp)
	assert.Nil(t, err)
	assert.NotNil(t, commonResp.Data)
	resp := &GetRebaseResp{}
	err = json.Unmarshal([]byte(commonResp.Data), resp)
	resp.ToMap()
	assert.Nil(t, err)
}
