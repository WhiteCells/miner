package info

type MinerStats struct {
	FsID       int       `json:"fs_id"`
	Coin       string    `json:"coin"`
	Temp       []int     `json:"temp"`
	Fan        []int     `json:"fan"`
	Power      []int     `json:"power"`
	Df         string    `json:"df"`
	Mem        []int     `json:"mem"`
	Cputemp    []int     `json:"cputemp"`
	Cpuavg     []float32 `json:"cpuavg"`
	Miner      string    `json:"miner"`
	TotalKhs   int       `json:"total_khs"`
	Status     string    `json:"status"`
	Khs        string    `json:"khs"`
	Hs         []float32 `json:"hs"`
	HsUnits    string    `json:"hs_units"`
	Algo       string    `json:"algo"`
	BusNumbers []int     `json:"bus_numbers"`
}

type MinerInfo struct {
	RigID         string   `json:"rig_id"`
	Passwd        string   `json:"passwd"`
	ServerUrl     string   `json:"server_url"`
	UID           string   `json:"uid"`
	RefId         string   `json:"ref_id"`
	BootTime      string   `json:"boot_time"`
	BootEvent     string   `json:"boot_event"`
	Ip            []string `json:"ip"`
	NetInterfaces []struct {
		Iface string `json:"iface"`
		Mac   string `json:"mac"`
	} `json:"net_interfaces"`
	Openvpn   string `json:"openvpn"`
	LanConfig struct {
		Dhcp    int    `json:"dhcp"`
		Address string `json:"address"`
		Gateway string `json:"gateway"`
		Dns     string `json:"dns"`
	} `json:"lan_config"`
	Gpu []struct {
		Busid     string `json:"busid"`
		Name      string `json:"name"`
		Brand     string `json:"brand"`
		Subvendor string `json:"subvendor"`
	} `json:"gpu"`
	GpuCountAmd    string `json:"gpu_count_amd"`
	GpuCountNvidia string `json:"gpu_count_nvidia"`
	GpuCountIntel  string `json:"gpu_count_intel"`
	Mb             struct {
		Manufacturer string `json:"manufacturer"`
		Product      string `json:"product"`
		SystemUuid   string `json:"system_uuid"`
		Bios         string `json:"bios"`
	} `json:"mb"`
	Cpu struct {
		Model string `json:"model"`
		Cores string `json:"cores"`
		Aes   string `json:"aes"`
		CpuID string `json:"cpu_id"`
	}
	DiskModel         string `json:"disk_model"`
	ImageVersion      string `json:"image_version"`
	Kernel            string `json:"kernel"`
	AmdVersion        string `json:"amd_version"`
	NvidiaVersion     string `json:"nvidia_version"`
	IntelVersion      string `json:"intel_version"`
	Version           string `json:"version"`
	ShellinaboxEnable bool   `json:"shellinabox_enable"`
	SshEnable         bool   `json:"ssh_enable"`
	SshPasswordEnable bool   `json:"ssh_password_enable"`
}
