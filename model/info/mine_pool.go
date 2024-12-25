package info

type MinePool struct {
	ID     string   `json:"id"`
	Url    []string `json:"url"`
	Status bool     `json:"status"`
}

/*
	{
		"id": 321,
		"url": ["http://", "http://"],
		"status": 1
	}
*/
