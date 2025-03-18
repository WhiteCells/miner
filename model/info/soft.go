package info

// 软件的参数很多都是不同的
// 此为 Custom 配置
type Soft struct {
	Coin             string `json:"coin" binding:"required,min=1,max=20"`
	MinerName        string `json:"minerName" binding:"required"` // Miner Name
	CustomInstallUrl string `json:"customInstallUrl"`             // installation URL
	CustomAlgo       string `json:"customAlgo"`                   // Hash algorithm
	CustomTemplate   string `json:"customTemplate"`               // Wallet and worker template
	CustomUrl        string `json:"customUrl"`                    // Pool URL
	CustomPass       string `json:"customPass"`                   // Pass
	CustomUserConfig string `json:"customUserConfig"`             // Extra config arguments
	CustomTls        string `json:"customTls"`                    // Tls
}
