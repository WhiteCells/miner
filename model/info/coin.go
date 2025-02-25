package info

type Coin struct {
	Name string `json:"name" binding:"required,min=1,max=20"`
}
