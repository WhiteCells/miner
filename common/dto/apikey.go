package dto

type GetBalanceReq struct {
	Address string `json:"address"`
}

type BscApiRspBody struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}
