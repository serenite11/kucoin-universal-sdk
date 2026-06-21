package marginprivate

import (
	"encoding/json"
	"testing"

	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestMarginPrivateCrossMarginPositionRespModel(t *testing.T) {
	// CrossMarginPosition
	// Get Cross Margin Position change

	data := "{\"topic\":\"/margin/position\",\"subject\":\"debt.ratio\",\"type\":\"message\",\"userId\":\"633559791e1cbc0001f319bc\",\"channelType\":\"private\",\"data\":{\"debtRatio\":0,\"totalAsset\":0.00052431772284080000000,\"marginCoefficientTotalAsset\":\"0.0005243177228408\",\"totalDebt\":\"0\",\"assetList\":{\"BTC\":{\"total\":\"0.00002\",\"available\":\"0\",\"hold\":\"0.00002\"},\"USDT\":{\"total\":\"33.68855864\",\"available\":\"15.01916691\",\"hold\":\"18.66939173\"}},\"debtList\":{\"BTC\":\"0\",\"USDT\":\"0\"},\"timestamp\":1729912435657}}"

	commonResp := &types.WsMessage{}
	err := json.Unmarshal([]byte(data), commonResp)
	assert.Nil(t, err)
	assert.NotNil(t, commonResp.RawData)
	obj := &CrossMarginPositionEvent{}
	err = json.Unmarshal(commonResp.RawData, obj)
	assert.Nil(t, err)
}

func TestMarginPrivateIsolatedMarginPositionRespModel(t *testing.T) {
	// IsolatedMarginPosition
	// Get Isolated Margin Position change

	data := "{\"topic\":\"/margin/isolatedPosition:BTC-USDT\",\"subject\":\"positionChange\",\"type\":\"message\",\"userId\":\"633559791e1cbc0001f319bc\",\"channelType\":\"private\",\"data\":{\"tag\":\"BTC-USDT\",\"status\":\"DEBT\",\"statusBizType\":\"DEFAULT_DEBT\",\"accumulatedPrincipal\":\"5.01\",\"changeAssets\":{\"BTC\":{\"total\":\"0.00043478\",\"hold\":\"0\",\"liabilityPrincipal\":\"0\",\"liabilityInterest\":\"0\"},\"USDT\":{\"total\":\"0.98092004\",\"hold\":\"0\",\"liabilityPrincipal\":\"26\",\"liabilityInterest\":\"0.00025644\"}},\"timestamp\":1730121097742}}"

	commonResp := &types.WsMessage{}
	err := json.Unmarshal([]byte(data), commonResp)
	assert.Nil(t, err)
	assert.NotNil(t, commonResp.RawData)
	obj := &IsolatedMarginPositionEvent{}
	err = json.Unmarshal(commonResp.RawData, obj)
	assert.Nil(t, err)
}
