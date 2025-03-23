package dto

type CreateFsReq struct {
	Name     string `json:"name" binding:"required,min=1,max=20"`
	FssubIDs []int  `json:"fssub_i_ds" binding:"required"`
}

type DelFsReq struct {
	FsID int `json:"fs_id" binding:"required"`
}

type DelSubReq struct {
	FssubIDs []int `json:"fssub_i_ds" binding:"required"`
}

type UpdateFsReq struct {
	FsID       int            `json:"fs_id" binding:"required"`
	UpdateInfo map[string]any `json:"update_info" binding:"required"`
}
