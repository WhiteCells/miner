package info

type Pool struct {
	Name   string   `json:"name"`
	Server []server `json:"server"`
}

type server struct {
	Name string `json:"name"`
	Host string `json:"host"`
}
