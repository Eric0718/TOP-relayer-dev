package eth2top

import (
	"context"
	"math/big"
	"sync"
	"testing"
	"toprelayer/base"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const SUBMITTERURL string = "http://192.168.50.204:19086"

//const LISTENURL string = "http://192.168.50.235:8545"
const LISTENURL string = "https://rinkeby.infura.io/v3/a3564d02d1bc4df58b7079a06b59a1cc"

var DEFAULTPATH = "../../.relayer/wallet/top"
var CONTRACT common.Address = common.HexToAddress("0xa3e165d80c949833C5c82550D697Ab31Fd3BB446")
var abipath string = "../../contract/topbridge/topbridge.abi"

func TestSubmitHeader(t *testing.T) {
	sub := &Eth2TopRelayer{}
	err := sub.Init(SUBMITTERURL, LISTENURL, DEFAULTPATH, "", abipath, base.TOP, CONTRACT, 90, 0, false)
	if err != nil {
		t.Fatal(err)
	}
	var batchHeaders []*types.Header

	/* 	currH, err := sub.getTopBridgeCurrentHeight()
	   	if err != nil {
	   		t.Fatal(err)
	   	}
	   	t.Log("bridge contract current height:", currH) */

	currH := uint64(10781005)

	header0, err := sub.ethsdk.HeaderByNumber(context.Background(), big.NewInt(0).SetUint64(currH))
	if err != nil {
		t.Fatal(err)
	}

	data1, _ := base.EncodeHeader(header0)
	t.Log("hex data1:", common.Bytes2Hex(data1))

	dd, _ := header0.MarshalJSON()
	t.Log("header data1:", string(dd))

	header1, err := sub.ethsdk.HeaderByNumber(context.Background(), big.NewInt(0).SetUint64(currH+1))
	if err != nil {
		t.Fatal(err)
	}
	batchHeaders = append(batchHeaders, header1)
	data2, _ := base.EncodeHeader(batchHeaders)
	t.Log("hex data2:", common.Bytes2Hex(data2))
	dd2, _ := header1.MarshalJSON()
	t.Log("header data2:", string(dd2))

	return

	batchHeaders = append(batchHeaders, header1)
	header2, err := sub.ethsdk.HeaderByNumber(context.Background(), big.NewInt(0).SetUint64(currH+2))
	if err != nil {
		t.Fatal(err)
	}

	batchHeaders = append(batchHeaders, header2)

	data, err := base.EncodeHeaders(batchHeaders)
	if err != nil {
		t.Fatal("EncodeToBytes:", err)
	}

	t.Log("header data:", data)

	nonce, err := sub.wallet.GetNonce(sub.wallet.CurrentAccount().Address)
	if err != nil {
		t.Fatal("GasPrice:", err)
	}

	tx, err := sub.submitEthHeader(data, nonce)
	if err != nil {
		t.Fatal("submitEthHeader:", err)
	}
	t.Log("hash:", tx.Hash())

	/* 	hashes, err := sub.signAndSendTransactions(1, 10)
	   	if err != nil {
	   		t.Fatal("signAndSendTransactions error:", err)
	   	}
	   	t.Log("hashes:", hashes) */

	/* nonce, err := sub.wallet.GetNonce(sub.wallet.CurrentAccount().Address)
	if err != nil {
		t.Fatal("GetNonce error:", err)
	}
	balance, err := sub.wallet.GetBalance(sub.wallet.CurrentAccount().Address)
	if err != nil {
		t.Fatal("GetBalance error:", err)
	}
	t.Log("balance:", balance, "nonce:", nonce)

	var headers []*types.Header
	for i := 1; i <= 2; i++ {
		headers = append(headers, &types.Header{Number: big.NewInt(int64(i))})
	}
	hash, err := sub.batch(headers, nonce)
	if err != nil {
		t.Fatal("batch error:", err)
	}
	t.Log("stx hash:", hash) */

	/* data, err := base.EncodeHeaders(&headers)
	if err != nil {
		t.Fatal("EncodeToBytes:", err)
	}

	if sub.wallet == nil {
		t.Fatal("nil wallet!!!")
	}

	stx, err := sub.submitEthHeader(data, nonce)
	if err != nil {
		t.Fatal("SubmitHeader error:", err)
	}
	t.Log("stx hash:", stx.Hash(), "type:", stx.Type())

	byt, err := stx.MarshalBinary()
	if err != nil {
		t.Fatal("MarshalBinary error:", err)
	}
	t.Log("rawtx:", hexutil.Encode(byt)) */
}

func TestEstimateGas(t *testing.T) {
	sub := &Eth2TopRelayer{}
	err := sub.Init(SUBMITTERURL, LISTENURL, DEFAULTPATH, "", abipath, base.TOP, CONTRACT, 90, 0, false)
	if err != nil {
		t.Fatal(err)
	}

	header := &types.Header{Number: big.NewInt(int64(1))}
	data, err := base.EncodeHeaders(header)
	if err != nil {
		t.Fatal("EncodeToBytes:", err)
	}

	pric, err := sub.wallet.GasPrice(context.Background())
	if err != nil {
		t.Fatal("GasPrice:", err)
	}

	gaslimit, err := sub.estimateGas(pric, data)
	if err != nil {
		t.Fatal("estimateGas:", err)
	}
	t.Log("gasprice", pric, "gaslimit:", gaslimit)
}

func TestGetTopBridgeState(t *testing.T) {
	sub := &Eth2TopRelayer{}
	err := sub.Init(SUBMITTERURL, LISTENURL, DEFAULTPATH, "", abipath, base.TOP, CONTRACT, 90, 0, false)
	if err != nil {
		t.Fatal(err)
	}

	curr, err := sub.getTopBridgeCurrentHeight()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("current height:", curr)
}

func TestStartRelayer(t *testing.T) {
	sub := &Eth2TopRelayer{}
	err := sub.Init(SUBMITTERURL, LISTENURL, DEFAULTPATH, "", abipath, base.TOP, CONTRACT, 90, 0, false)
	if err != nil {
		t.Fatal(err)
	}

	wg := &sync.WaitGroup{}
	err = sub.StartRelayer(wg)
	if err != nil {
		t.Fatal(err)
	}
}
