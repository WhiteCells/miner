package info

type Fs struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Coin     string `json:"coin"`
	WalletID string `json:"wallet_id"` // 钱包的名字可以重复
	Pool     string `json:"pool"`      // 名字即为 ID
	Soft     string `json:"soft"`      // 名字即为 ID
}
