package dto

// Hiveos 轮询
// hello
// stats
// Hiveos ---> Server
type HiveosReq struct {
	Method string `json:"method"` // 请求方法 hello、stats 或 message
	Params struct {
		V      int    `json:"v"`      //
		RigID  string `json:"rig_id"` // 矿机 ID，由系统生成
		Passwd string `json:"passwd"` // 矿机密码，由系统生成
		Meta   struct {
			FsID   int `json:"fs_id"` // 矿机使用的飞行表 ID
			Custom struct {
				Coin string `json:"coin"` // 矿机使用代币
			}
		} `json:"meta"`
		Temp       []int     `json:"temp"`      // 温度
		Fan        []int     `json:"fan"`       // 风扇
		Power      []int     `json:"power"`     // 电源
		Df         string    `json:"df"`        // 磁盘
		Mem        []int     `json:"mem"`       // 内存
		Cputemp    []int     `json:"cputemp"`   // cpu 温度
		Cpuavg     []float32 `json:"cpuavg"`    // 平均负载
		Miner      string    `json:"miner"`     // 软件
		TotalKhs   int       `json:"total_khs"` // 总算力
		MinerStats struct {
			Status     string    `json:"status"`      // 状态
			Khs        string    `json:"khs"`         // 挖矿软件算力
			Hs         []float32 `json:"hs"`          // 每张卡算力
			HsUnits    string    `json:"hs_units"`    // 算力单位
			Temp       []int     `json:"temp"`        // 每张卡温度
			Fan        []int     `json:"fan"`         // 每张卡风扇转速
			Uptime     int       `json:"uptime"`      // 运行时长
			Ver        string    `json:"ver"`         // 挖矿软件版本
			Ar         []int     `json:"ar"`          // 矿池接受的提交次数
			Algo       string    `json:"algo"`        // 算法
			BusNumbers []int     `json:"bus_numbers"` // 每张显卡的 PCI 总线号
		} `json:"miner_stats"`
	} `json:"params"`
}

// message
// Hiveos ---> Server
type HiveosResReq struct {
	Method  string `json:"method"`
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Params  struct {
		RigID   string `json:"rig_id"`
		Passwd  string `json:"passwd"`
		Type    string `json:"type"`
		Data    string `json:"data"`
		ID      string `json:"id"`
		Payload string `json:"payload"`
	} `json:"params"`
}

type ServerRsp struct {
	ID      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		ID        int    `json:"id"`
		Config    string `json:"config"`
		Wallet    string `json:"wallet"`
		Autofan   string `json:"autofan"`
		Justwrite int    `json:"justwrite"`
		Command   string `json:"command"`
		Exec      string `json:"exec"`
		Confseq   int    `json:"confseq"`
	} `json:"result"`
}
