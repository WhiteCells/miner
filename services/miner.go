package services

import (
	"context"
	"errors"
	"miner/common/dto"
	"miner/common/perm"
	"miner/dao/mysql"
	"miner/dao/mysql/relationdao"
	"miner/dao/redis"
	"miner/model"
	"miner/model/info"
	"miner/utils"
	"slices"
)

type MinerService struct {
	minerDAO    *mysql.MinerDAO
	userFarmDAO *relationdao.UserFarmDAO
	minerFsDAO  *relationdao.MinerFsDAO
	minerRDB    *redis.MinerRDB
}

func NewMinerService() *MinerService {
	return &MinerService{
		minerDAO:    mysql.NewMinerDAO(),
		userFarmDAO: relationdao.NewUserFarmDAO(),
		minerFsDAO:  relationdao.NewMinerFsDAO(),
		minerRDB:    redis.NewMinerRDB(),
	}
}

func (m *MinerService) CreateMiner(ctx context.Context, userID, farmID int, req *dto.CreateMinerReq) (*model.Miner, error) {
	if err := m.validPerm(ctx, userID, farmID, []perm.FarmPerm{perm.FarmOwner, perm.FarmManager}); err != nil {
		return nil, err
	}

	pass, err := utils.GenerateRigPass(8)
	if err != nil {
		return nil, err
	}

	miner := &model.Miner{
		Name: req.Name,
		Pass: pass,
	}
	if err := m.minerDAO.CreateMiner(ctx, farmID, miner); err != nil {
		return nil, err
	}

	minerInfo := &info.Miner{
		HiveOsConfig: utils.HiveOsConfig{
			HiveOsUrl:     utils.GenerateHiveOsUrl(),
			ApiHiveOsUrls: utils.GenerateHiveOsUrl(),
			WorkerName:    req.Name,
			FarmID:        farmID,
			RigID:         miner.ID,
			RigPasswd:     pass,
		},
	}

	// redis 缓存 miner 配置
	if err := m.minerRDB.CreateMinerByRigID(ctx, miner.ID, minerInfo); err != nil {
		if err := m.minerDAO.DelMiner(ctx, userID, miner.ID); err != nil {
			return nil, errors.New("set cached failed and del failed")
		}
		return nil, errors.New("set cached failed")
	}

	return miner, nil
}

func (m *MinerService) DelMiner(ctx context.Context, userID, farmID, minerID int) error {
	if err := m.validPerm(ctx, userID, farmID, []perm.FarmPerm{perm.FarmOwner, perm.FarmManager}); err != nil {
		return err
	}
	miner, err := m.minerDAO.GetMinerByID(ctx, minerID)
	if err != nil {
		return err
	}
	if err := m.minerRDB.DelMinerByRigID(ctx, miner.ID); err != nil {
		return err
	}
	if err := m.minerDAO.DelMiner(ctx, userID, minerID); err != nil {
		return err
	}
	return nil
}

func (m *MinerService) UpdateMiner(ctx context.Context, userID, farmID, minerID int, updateInfo map[string]any) error {
	if err := m.validPerm(ctx, userID, farmID, []perm.FarmPerm{perm.FarmOwner, perm.FarmManager}); err != nil {
		return err
	}
	allow := model.GetMinerAllowChangeField()
	updates := make(map[string]any)
	for key, val := range updateInfo {
		if allow[key] {
			updates[key] = val
		}
	}
	return m.minerDAO.UpdateMiner(ctx, userID, minerID, updates)
}

func (m *MinerService) UpdateMinerWatchdog(ctx context.Context, userID, farmID, minerID int, req *dto.UpdateMinerWatchdogReq) error {
	if err := m.validPerm(ctx, userID, farmID, []perm.FarmPerm{perm.FarmOwner, perm.FarmManager}); err != nil {
		return err
	}

	miner, err := m.minerDAO.GetMinerByID(ctx, minerID)
	if err != nil {
		return err
	}

	minerInfo, err := m.minerRDB.GetMinerByRigID(ctx, miner.ID)
	if err != nil {
		return err
	}

	minerInfo.HiveOsConfig.Watchdog = req.Watchdog

	if err := m.minerRDB.UpdateMinerByRigID(ctx, miner.ID, minerInfo); err != nil {
		return err
	}

	return nil
}

func (m *MinerService) UpdateMinerOptions(ctx context.Context, userID int, req *dto.UpdateMinerOptionsReq) error {
	if err := m.validPerm(ctx, userID, req.FarmID, []perm.FarmPerm{perm.FarmOwner, perm.FarmManager}); err != nil {
		return err
	}

	miner, err := m.minerDAO.GetMinerByID(ctx, req.MinerID)
	if err != nil {
		return err
	}

	minerInfo, err := m.minerRDB.GetMinerByRigID(ctx, miner.ID)
	if err != nil {
		return err
	}

	minerInfo.HiveOsConfig.Options = req.Options

	if err := m.minerRDB.UpdateMinerByRigID(ctx, miner.ID, minerInfo); err != nil {
		return err
	}

	return nil
}

func (m *MinerService) UpdateMinerAutofan(ctx context.Context, userID int, req *dto.UpdateMinerAutofanReq) error {
	if err := m.validPerm(ctx, userID, req.FarmID, []perm.FarmPerm{perm.FarmOwner, perm.FarmManager}); err != nil {
		return err
	}

	miner, err := m.minerDAO.GetMinerByID(ctx, req.MinerID)
	if err != nil {
		return err
	}

	minerInfo, err := m.minerRDB.GetMinerByRigID(ctx, miner.ID)
	if err != nil {
		return err
	}

	if err := m.minerRDB.UpdateMinerByRigID(ctx, miner.ID, minerInfo); err != nil {
		return err
	}

	return nil
}

func (m *MinerService) UpdateMinerWallet(ctx context.Context, userID int, req *dto.UpdateMinerWalletReq) error {
	if err := m.validPerm(ctx, userID, req.FarmID, []perm.FarmPerm{perm.FarmOwner, perm.FarmManager}); err != nil {
		return err
	}

	miner, err := m.minerDAO.GetMinerByID(ctx, req.MinerID)
	if err != nil {
		return err
	}

	minerInfo, err := m.minerRDB.GetMinerByRigID(ctx, miner.ID)
	if err != nil {
		return err
	}

	minerInfo.HiveOsWallet = req.Wallet

	return m.minerRDB.UpdateMinerByRigID(ctx, miner.ID, minerInfo)
}

func (m *MinerService) SetWatchdog(ctx context.Context, userID int, req *dto.SetWatchdogReq) error {
	if err := m.validPerm(ctx, userID, req.FarmID, []perm.FarmPerm{perm.FarmOwner, perm.FarmManager}); err != nil {
		return err
	}

	miner, err := m.minerDAO.GetMinerByID(ctx, req.MinerID)
	if err != nil {
		return err
	}

	minerInfo, err := m.minerRDB.GetMinerByRigID(ctx, miner.ID)
	if err != nil {
		return err
	}

	minerInfo.HiveOsConfig.Watchdog = req.Watchdog

	return m.minerRDB.UpdateMinerByRigID(ctx, miner.ID, minerInfo)
}

func (m *MinerService) GetWatchdog(ctx context.Context, userID, farmID, minerID int) (*utils.Watchdog, error) {
	miner, err := m.minerDAO.GetMinerByID(ctx, minerID)
	if err != nil {
		return nil, err
	}
	minerInfo, err := m.minerRDB.GetMinerByRigID(ctx, miner.ID)
	if err != nil {
		return nil, err
	}
	return &minerInfo.HiveOsConfig.Watchdog, nil
}

func (m *MinerService) SetAutoFan(ctx context.Context, userID int, req *dto.SetAutoFanReq) error {
	if err := m.validPerm(ctx, userID, req.FarmID, []perm.FarmPerm{perm.FarmOwner, perm.FarmManager}); err != nil {
		return err
	}
	miner, err := m.minerDAO.GetMinerByID(ctx, req.MinerID)
	if err != nil {
		return err
	}
	minerInfo, err := m.minerRDB.GetMinerByRigID(ctx, miner.ID)
	if err != nil {
		return err
	}
	minerInfo.HiveOsAutoFan = req.AutoFan
	return m.minerRDB.UpdateMinerByRigID(ctx, miner.ID, minerInfo)
}

func (m *MinerService) GetAutoFan(ctx context.Context, userID, farmID, minerID int) (*utils.HiveOsAutoFan, error) {
	miner, err := m.minerDAO.GetMinerByID(ctx, minerID)
	if err != nil {
		return nil, err
	}
	minerInfo, err := m.minerRDB.GetMinerByRigID(ctx, miner.ID)
	if err != nil {
		return nil, err
	}
	return &minerInfo.HiveOsAutoFan, nil
}

func (m *MinerService) SetOptions(ctx context.Context, userID int, req *dto.SetOptionsReq) error {
	if err := m.validPerm(ctx, userID, req.FarmID, []perm.FarmPerm{perm.FarmOwner, perm.FarmManager}); err != nil {
		return err
	}
	miner, err := m.minerDAO.GetMinerByID(ctx, req.MinerID)
	if err != nil {
		return err
	}
	minerInfo, err := m.minerRDB.GetMinerByRigID(ctx, miner.ID)
	if err != nil {
		return err
	}
	minerInfo.HiveOsConfig.Options = req.Options
	return m.minerRDB.UpdateMinerByRigID(ctx, miner.ID, minerInfo)
}

func (m *MinerService) GetOptions(ctx context.Context, userID, farmID, minerID int) (*utils.Options, error) {
	miner, err := m.minerDAO.GetMinerByID(ctx, minerID)
	if err != nil {
		return nil, err
	}
	minerInfo, err := m.minerRDB.GetMinerByRigID(ctx, miner.ID)
	if err != nil {
		return nil, err
	}
	return &minerInfo.HiveOsConfig.Options, nil
}

func (m *MinerService) GetMinerByMinerID(ctx context.Context, userID, minerID int) (*model.Miner, error) {
	return m.minerDAO.GetMinerByID(ctx, minerID)
}

func (m *MinerService) GetMinersByFarmID(ctx context.Context, farmID int, query map[string]any) ([]model.Miner, int64, error) {
	return m.minerDAO.GetMinersByFarmID(ctx, farmID, query)
}

func (m *MinerService) GetMiners(ctx context.Context, query map[string]any) ([]model.Miner, int64, error) {
	return m.minerDAO.GetMiners(ctx, query)
}

func (m *MinerService) ApplyFs(ctx context.Context, userID, farmID, minerID, fsID int) error {
	if err := m.validPerm(ctx, userID, farmID, []perm.FarmPerm{perm.FarmOwner, perm.FarmManager}); err != nil {
		return err
	}
	// 更新 miner cache
	return m.minerFsDAO.BindFsToMiner(ctx, fsID, minerID)
}

func (m *MinerService) UnApplyFs(ctx context.Context, userID, farmID, minerID, fsID int) error {
	if err := m.validPerm(ctx, userID, farmID, []perm.FarmPerm{perm.FarmOwner, perm.FarmManager}); err != nil {
		return err
	}
	return m.minerFsDAO.UnBindFsFromMiner(ctx, fsID, minerID)
}

func (m *MinerService) GetApplyFs(ctx context.Context, farmID, minerID, fsID int) (int, error) {
	return m.minerFsDAO.GetFsIDFromMiner(ctx, minerID)
}

func (m *MinerService) Transfer(ctx context.Context, userID, farmID, minerID int, toFarmHash string) error {
	if err := m.validPerm(ctx, userID, farmID, []perm.FarmPerm{perm.FarmOwner, perm.FarmManager}); err != nil {
		return err
	}
	return m.minerDAO.Transfer(ctx, farmID, minerID, toFarmHash)
}

func (m *MinerService) validPerm(ctx context.Context, userID, farmID int, allowedPerms []perm.FarmPerm) error {
	perm, err := m.userFarmDAO.GetPerm(ctx, userID, farmID)
	if err != nil {
		return err
	}
	if !slices.Contains(allowedPerms, perm) {
		return errors.New("permission denied")
	}
	return nil
}

// 弃用方法
// func (m *MinerService) generateRigID(ctx context.Context, length int) (string, error) {
// 	if length < 8 {
// 		return "", errors.New("invalid argument")
// 	}
// 	const charset = "123456789" // 以 0 开头时会导致转为字符串与实际 ID 不符合
// 	id := make([]byte, length)
// 	for {
// 		for i := range id {
// 			num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
// 			if err != nil {
// 				return "", err
// 			}
// 			id[i] = charset[num.Int64()]
// 		}
// 		uid := string(id)
// 		if !m.minerDAO.ExistsRigID(ctx, uid) {
// 			return uid, nil
// 		}
// 	}
// }

// func (m *MinerService) generateRigPass(length int) (string, error) {
// 	if length < 8 {
// 		return "", errors.New("invalid argument")
// 	}
// 	const charset = "abcdefghijklmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ123456789"
// 	str := make([]byte, length)
// 	for i := range str {
// 		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
// 		if err != nil {
// 			return "", err
// 		}
// 		str[i] = charset[num.Int64()]
// 	}
// 	return string(str), nil
// }
