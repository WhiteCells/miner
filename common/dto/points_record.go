package dto

type GetUserPointsRecordsReq struct {
	PageNum  int `json:"page_num" binding:"required"`
	PageSize int `json:"page_size" binding:"required"`
}
