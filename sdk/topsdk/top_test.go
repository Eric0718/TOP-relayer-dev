package topsdk

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func newtopsdk() (*TopSdk, error) {
	url := "http://192.168.50.204:19086"
	return NewTopSdk(url)
}

func TestGetTopElectBlockHeadByHeight(t *testing.T) {
	sdk, err := newtopsdk()
	if err != nil {
		t.Fatalf("NewSDK failed,error:%v", err)
	}

	/* h, err := sdk.GetLatestTopElectBlockHeight()
	if err != nil {
		t.Fatalf("GetLatestTopElectBlockHeight failed,error:%v", err)
	} */

	result, err := sdk.GetTopElectBlockHeadByHeight(1)
	if err != nil {
		t.Fatalf("GetTopElectBlockHeadByHeight failed,error:%v", err)
	}
	t.Logf("GetTopElectBlockHeadByHeight ok,result:%v", common.Bytes2Hex(result))
}

func TestGetLatestTopElectBlockHeight(t *testing.T) {
	sdk, err := newtopsdk()
	if err != nil {
		t.Fatalf("NewSDK failed,error:%v", err)
	}

	h, err := sdk.GetLatestTopElectBlockHeight()
	if err != nil {
		t.Fatalf("GetLatestTopElectBlockHeight failed,error:%v", err)
	}
	t.Logf("GetLatestTopElectBlockHeight:%v", h)
}
