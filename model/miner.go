package model

type Miner struct {
	ID          int    `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement"`
	Name        string `json:"name" gorm:"column:name;type:varchar(255)"`
	GPUInfo     string `json:"gpu_info" gorm:"column:gpu_info;type:varchar(255)"`
	Status      int    `json:"status" gorm:"column:status;type:int"`
	IPAddress   string `json:"ip_address" gorm:"column:ip_address;type:varchar(255)"`
	SSHPort     int    `json:"ssh_port" gorm:"column:ssh_port;type:int"`
	SSHUser     string `json:"ssh_user" gorm:"column:ssh_user;type:varchar(255)"`
	SSHPassword string `json:"ssh_password" gorm:"column:ssh_password;type:varchar(255)"`
}

func (Miner) TableName() string {
	return "miner"
}
