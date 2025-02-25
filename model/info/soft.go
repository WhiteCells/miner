package info

// 软件的参数很多都是不同的
// 此为 Custom 配置
type Soft struct {
	MinerName        string `json:"custom_name"`        // Miner Name
	CustomInstallUrl string `json:"custom_install_url"` // installation URL
	CustomAlgo       string `json:"custom_algo"`        // Hash algorithm
	CustomTemplate   string `json:"custom_template"`    // Wallet and worker template
	CustomUrl        string `json:"custom_url"`         // Pool URL
	CustomPass       string `json:"custom_pass"`        // Pass
	CustomUserConfig string `json:"custom_user_config"` // Extra config arguments
	CustomTls        string `json:"custom_tls"`         // todo 未知
}
