package dto

// // Hiveos 轮询
// type HelloReq struct {
// 	Method  string          `json:"method"`
// 	Jsonrpc string          `json:"jsonrpc"`
// 	ID      int             `json:"id"`
// 	Params  HelloReq_Params `json:"params"`
// }

// type HelloReq_Params struct {
// 	FarmHash          string                          `json:"farm_hash"`
// 	RigID             string                          `json:"rig_id"`
// 	Passwd            string                          `json:"passwd"`
// 	ServerUrl         string                          `json:"server_url"`
// 	UID               string                          `json:"uid"`
// 	RefId             string                          `json:"ref_id"`
// 	BootTime          string                          `json:"boot_time"`
// 	BootEvent         string                          `json:"boot_event"`
// 	Ip                []string                        `json:"ip"`
// 	NetInterfaces     []HelloReq_Params_NetInterfaces `json:"net_interfaces"`
// 	Openvpn           string                          `json:"openvpn"`
// 	LanConfig         HelloReq_Params_LanConfig       `json:"lan_config"`
// 	Gpu               []HelloReq_Params_Gpu           `json:"gpu"`
// 	GpuCountAmd       string                          `json:"gpu_count_amd"`
// 	GpuCountNvidia    string                          `json:"gpu_count_nvidia"`
// 	GpuCountIntel     string                          `json:"gpu_count_intel"`
// 	Mb                HelloReq_Params_Mb              `json:"mb"`
// 	Cpu               Cpu                             `json:"cpu"`
// 	DiskModel         string                          `json:"disk_model"`
// 	ImageVersion      string                          `json:"image_version"`
// 	Kernel            string                          `json:"kernel"`
// 	AmdVersion        string                          `json:"amd_version"`
// 	NvidiaVersion     string                          `json:"nvidia_version"`
// 	IntelVersion      string                          `json:"intel_version"`
// 	Version           string                          `json:"version"`
// 	ShellinaboxEnable bool                            `json:"shellinabox_enable"`
// 	SshEnable         bool                            `json:"ssh_enable"`
// 	SshPasswordEnable bool                            `json:"ssh_password_enable"`
// }

// type HelloReq_Params_NetInterfaces struct {
// 	Iface string `json:"iface"`
// 	Mac   string `json:"mac"`
// }

// type HelloReq_Params_LanConfig struct {
// 	Dhcp    int    `json:"dhcp"`
// 	Address string `json:"address"`
// 	Gateway string `json:"gateway"`
// 	Dns     string `json:"dns"`
// }

// type HelloReq_Params_Gpu struct {
// 	Busid     string `json:"busid"`
// 	Name      string `json:"name"`
// 	Brand     string `json:"brand"`
// 	Subvendor string `json:"subvendor"`
// }

// type HelloReq_Params_Mb struct {
// 	Manufacturer string `json:"manufacturer"`
// 	Product      string `json:"product"`
// 	SystemUuid   string `json:"system_uuid"`
// 	Bios         string `json:"bios"`
// }

// type Cpu struct {
// 	Model string `json:"model"`
// 	Cores string `json:"cores"`
// 	Aes   string `json:"aes"`
// 	CpuID string `json:"cpu_id"`
// }

// // HiveOsReq
// type HiveOsReq struct {
// 	Method string           `json:"method"` // 请求方法 hello、stats 或 message
// 	Params HiveOsReq_Params `json:"params"`
// }

// type HiveOsReq_Params struct {
// 	V          int                         `json:"v"`         //
// 	RigID      string                      `json:"rig_id"`    // 矿机 ID，由系统生成
// 	Passwd     string                      `json:"passwd"`    // 矿机密码，由系统生成
// 	Meta       HiveOsReq_Params_Meta       `json:"meta"`      //
// 	Temp       []int                       `json:"temp"`      // 温度
// 	Fan        []int                       `json:"fan"`       // 风扇
// 	Power      []int                       `json:"power"`     // 电源
// 	Df         string                      `json:"df"`        // 磁盘
// 	Mem        []int                       `json:"mem"`       // 内存
// 	Cputemp    []int                       `json:"cputemp"`   // cpu 温度
// 	Cpuavg     []float32                   `json:"cpuavg"`    // 平均负载
// 	Miner      string                      `json:"miner"`     // 软件
// 	TotalKhs   int                         `json:"total_khs"` // 总算力
// 	MinerStats HiveOsReq_Params_MinerStats `json:"miner_stats"`
// }

// // HiveosReq_Params
// // - HiveosReq_Params_Meta
// type HiveOsReq_Params_Meta struct {
// 	FsID   int                          `json:"fs_id"`  // 矿机使用的飞行表 ID
// 	Custom HiveOsReq_Params_Meta_Custom `json:"custom"` //
// }

// // HiveosReq_Params
// // - HiveOsReq_Params_Meta_Custom
// type HiveOsReq_Params_Meta_Custom struct {
// 	Coin string `json:"coin"` // 矿机使用代币
// }

// // HiveosReq_Params
// // - HiveOsReq_Params_MinerStats
// type HiveOsReq_Params_MinerStats struct {
// 	Status     string    `json:"status"`      // 状态
// 	Khs        string    `json:"khs"`         // 挖矿软件算力
// 	Hs         []float32 `json:"hs"`          // 每张卡算力
// 	HsUnits    string    `json:"hs_units"`    // 算力单位
// 	Temp       []int     `json:"temp"`        // 每张卡温度
// 	Fan        []int     `json:"fan"`         // 每张卡风扇转速
// 	Uptime     int       `json:"uptime"`      // 运行时长
// 	Ver        string    `json:"ver"`         // 挖矿软件版本
// 	Ar         []int     `json:"ar"`          // 矿池接受的提交次数
// 	Algo       string    `json:"algo"`        // 算法
// 	BusNumbers []int     `json:"bus_numbers"` // 每张显卡的 PCI 总线号
// }

// // message
// // HiveosResReq
// type HiveosResReq struct {
// 	Method  string              `json:"method"`
// 	Jsonrpc string              `json:"jsonrpc"`
// 	ID      int                 `json:"id"`
// 	Params  HiveosResReq_Params `json:"params"`
// }

// // HiveosResReq
// //  - Params
// type HiveosResReq_Params struct {
// 	RigID   string `json:"rig_id"`
// 	Passwd  string `json:"passwd"`
// 	Type    string `json:"type"`
// 	Data    string `json:"data"`
// 	ID      string `json:"id"`
// 	Payload string `json:"payload"`
// }

// type ServerRsp struct {
// 	ID      int              `json:"id"`
// 	Jsonrpc string           `json:"jsonrpc"`
// 	Result  ServerRsp_Result `json:"result"`
// }

// // ServerRsp
// //  - Result
// type ServerRsp_Result struct {
// 	ID        int    `json:"id"`
// 	Config    string `json:"config"`
// 	Wallet    string `json:"wallet"`
// 	Autofan   string `json:"autofan"`
// 	Justwrite int    `json:"justwrite"`
// 	Command   string `json:"command"`
// 	Exec      string `json:"exec"`
// 	Confseq   int    `json:"confseq"`
// }
