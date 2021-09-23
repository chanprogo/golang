package execcmd

import (
	"encoding/json"
	"math/big"
)

type BalanceResult struct {
	ID      int    `json:"id"`
	JSONRPC string `json:"jsonrpc"`
	Result  string `json:"result"`
}

func GetBalance(address string) (*big.Int, error) {

	method := "eth_getBalance"
	params := `["` + address + `","latest"]`

	jsrp, err := CallGeth(method, params)
	if err != nil {
		return nil, err
	}

	r := BalanceResult{}
	err = json.Unmarshal([]byte(jsrp), &r)
	if err != nil {
		return nil, err
	}
	balance := big.NewInt(0)
	balance.SetString(r.Result, 0)

	return balance, nil
}
