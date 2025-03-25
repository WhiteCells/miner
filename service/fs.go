package service

// import (
// 	"context"
// 	"errors"
// 	"miner/common/dto"
// 	"miner/dao/redis"
// 	"miner/model/info"
// 	"miner/utils"
// )

// type FsService struct {
// 	fsRDB *redis.FsRDB
// }

// func NewFsService() *FsService {
// 	return &FsService{
// 		fsRDB: redis.NewFsRDB(),
// 	}
// }

// // CreateFs 创建飞行表
// func (s *FsService) CreateFs(ctx context.Context, req *dto.CreateFsReq) (*info.Fs, error) {
// 	userID, exists := ctx.Value("user_id").(string)
// 	if !exists {
// 		return nil, errors.New("invalid user_id in context")
// 	}

// 	id, err := utils.GenerateUID()
// 	if err != nil {
// 		return nil, err
// 	}
// 	fs := &info.Fs{
// 		ID:       id,
// 		Name:     req.Name,
// 		Coin:     req.Coin,
// 		WalletID: req.WalletID,
// 		Pool:     req.Pool,
// 		Soft:     req.Soft,
// 	}

// 	if err := s.fsRDB.Set(ctx, userID, fs); err != nil {
// 		return nil, errors.New("create flightsheet failed")
// 	}

// 	return fs, nil
// }

// // DeleteFs 删除飞行表
// func (s *FsService) DeleteFs(ctx context.Context, req *dto.DeleteFsReq) error {
// 	userID, exists := ctx.Value("user_id").(string)
// 	if !exists {
// 		return errors.New("invalid user_id in context")
// 	}

// 	if err := s.fsRDB.Del(ctx, userID, req.FsID); err != nil {
// 		return errors.New("delete flightsheet failed")
// 	}

// 	return nil
// }

// // UpdateFs 更新飞行表
// func (s *FsService) UpdateFs(ctx context.Context, req *dto.UpdateFsReq) error {
// 	userID, exists := ctx.Value("user_id").(string)
// 	if !exists {
// 		return errors.New("invalid user_id in context")
// 	}

// 	// 查找飞行表
// 	fs, err := s.fsRDB.GetByID(ctx, userID, req.FsID)
// 	if err != nil {
// 		return errors.New("flightsheet not found")
// 	}

// 	for key, value := range req.UpdateInfo {
// 		switch key {
// 		case "name":
// 			fs.Name = value.(string)
// 		case "coin":
// 			fs.Coin = value.(string)
// 		case "wallet_id":
// 			fs.WalletID = value.(string)
// 		case "pool":
// 			fs.Pool = value.(string)
// 		case "soft":
// 			fs.Soft = value.(string)
// 		}
// 	}

// 	if err := s.fsRDB.Set(ctx, userID, fs); err != nil {
// 		return errors.New("update flightsheet failed")
// 	}

// 	return nil
// }

// // GetAllFs 获取用户的所有飞行表
// func (s *FsService) GetAllFs(ctx context.Context) ([]info.Fs, error) {
// 	userID, exists := ctx.Value("user_id").(string)
// 	if !exists {
// 		return nil, errors.New("invalid user_id in context")
// 	}
// 	fss, err := s.fsRDB.GetAll(ctx, userID)
// 	if err != nil {
// 		return nil, errors.New("get flightsheet failed")
// 	}
// 	return fss, err
// }

// // GetFsByID
// func (s *FsService) GetFsByID(ctx context.Context, fsID string) (*info.Fs, error) {
// 	userID, exists := ctx.Value("user_id").(string)
// 	if !exists {
// 		return nil, errors.New("invalid user_id in context")
// 	}
// 	fs, err := s.fsRDB.GetByID(ctx, userID, fsID)
// 	if err != nil {
// 		return nil, errors.New("get flightsheet failed")
// 	}
// 	return fs, err
// }

// // ApplyWallet 飞行表应用钱包
// // func (s *FsService) ApplyWallet(ctx context.Context, req *dto.ApplyWalletReq) error {
// // 	userID, exists := ctx.Value("user_id").(string)
// // 	if !exists {
// // 		return errors.New("invalid user_id in context")
// // 	}
// // 	if err := s.fsRDB.ApplyWallet(ctx, userID, req.FsID, req.WaleltID); err != nil {
// // 		return err
// // 	}
// // 	return nil
// // }
