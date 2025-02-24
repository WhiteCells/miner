package info

type Pool struct {
	Name string `json:"name"`
	Urls []Url  `json:"urls"`
}

type Url struct {
	Name string `json:"name"`
	Host string `json:"host"`
}
