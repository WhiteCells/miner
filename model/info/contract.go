package info

type Contract struct {
	Net     string `json:"net"`     // 网路（链）https://bsc-dataseed.binance.org/
	Coin    string `json:"coin"`    // 代币 BNB BUSD
	Address string `json:"address"` // 合约地址 0xe9e7cea3dedca5984780bafc599bd69add087d56
}
