package info

type Pool struct {
	Name string `json:"name" binding:"required,min=1,max=20"`
	Urls []Url  `json:"urls" binding:"required"`
}

type Url struct {
	Name string `json:"name" binding:"required,min=1,max=20"`
	Host string `json:"host" binding:"required,min=1,max=20"`
}
