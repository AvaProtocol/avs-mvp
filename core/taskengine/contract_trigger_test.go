package taskengine

import (
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/expr-lang/expr"
)

type answer struct {
	RoundId         big.Int
	Answer          big.Int
	StartedAt       uint64
	UpdatedAt       uint64
	AnsweredInRound big.Int
}

func TestQueryContract(t *testing.T) {
	conn, _ := ethclient.Dial(os.Getenv("RPC_URL"))

	r, err := QueryContract(
		conn,
		// https://docs.chain.link/data-feeds/price-feeds/addresses?network=ethereum&page=1&search=et#sepolia-testnet
		// ETH-USD pair on sepolia
		common.HexToAddress("0x694AA1769357215DE4FAC081bf1f309aDC325306"),
		chainlinkABI,
		"latestRoundData",
	)

	if err != nil {
		t.Errorf("contract query error: %v", err)
	}

	t.Logf("contract query result: %v", r)
}

func TestExpression(t *testing.T) {
	SetRpc(os.Getenv("RPC_URL"))

	p, e := CompileExpression(`priceChainlink("0x694AA1769357215DE4FAC081bf1f309aDC325306")`)
	if e != nil {
		t.Errorf("Compile expression error: %v", e)
	}

	r, e := expr.Run(p, exprEnv)
	if e != nil {
		t.Errorf("Run expr error: %v %v", e, r)
	}

	if r.(*big.Int).Cmp(big.NewInt(10)) <= 0 {
		t.Errorf("Invalid result data: %v", r)
	}

	t.Logf("Exp Run Result: %v", r.(*big.Int))

	match, e := RunExpressionQuery(`
		bigCmp(
		  priceChainlink("0x694AA1769357215DE4FAC081bf1f309aDC325306"),
		  toBigInt("2000")
		) > 0
	`)
	if e != nil {
		t.Errorf("Run expr error: %v %v", e, r)
	}
	if !match {
		t.Error("Evaluate error. Expected: true, received: false")
	}

	match, e = RunExpressionQuery(`
		bigCmp(
		  priceChainlink("0x694AA1769357215DE4FAC081bf1f309aDC325306"),
		  toBigInt("9262391230023")
		) > 0
	`)
	if e != nil {
		t.Errorf("Run expr error: %v %v", e, r)
	}
	if match {
		t.Error("Evaluate error. Expected: false, got: true")
	}
}

func TestExpressionPanicWonCrash(t *testing.T) {
	rpcConn = nil
	p, e := CompileExpression(`priceChainlink("0x694AA1769357215DE4FAC081bf1f309aDC325306")`)
	if e != nil {
		t.Errorf("Compile expression error: %v", e)
	}

	r, e := expr.Run(p, exprEnv)
	if e == nil || r != nil {
		t.Errorf("Evaluate wrong. Expected: nil, got: %v", r)
	}

	t.Logf("Succesfully recovered from VM crash")
}
