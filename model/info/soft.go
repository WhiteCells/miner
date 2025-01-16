package info

// 软件的参数很多都是不同的
// 此为 Custom 配置
type Soft struct {
	Name        string `json:"name"`         // 软件名称
	InstallUrl  string `json:"install_url"`  // 安装链接
	EncryptAlgo string `json:"encrypt_algo"` // 加密算法
	Temp        string `json:"temp"`         // 钱包与矿机模板
	PoolUrl     string `json:"pool_url"`     // 矿池地址
	Pass        string `json:"pass"`         // 密码
	Other       string `json:"other"`        // 其他配置参数
}
