package info

type Fs struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Coin string `json:"coin"`
	Mine string `json:"mine"`
	Soft string `json:"soft"`
}

/*
	{
		"id": 101,
		"name": "name",
		"coin": "ETC",
		"mine": 321,
		"soft": "custom",
		"wallet": 123
	}
*/
