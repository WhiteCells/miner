package services

import (
	"context"
	"crypto/rand"
	"errors"
	"math/big"
	"miner/common/dto"
	"miner/common/perm"
	"miner/dao/mysql"
	"miner/dao/mysql/relationdao"
	"miner/dao/redis"
	"miner/model"
	"miner/model/info"
	"miner/utils"
	"slices"
	"strconv"
)

type MinerService struct {
	minerDAO    *mysql.MinerDAO
	minerRDB    *redis.MinerRDB
	userFarmDAO *relationdao.UserFarmDAO
}

func NewMinerService() *MinerService {
	return &MinerService{
		minerDAO:    mysql.NewMinerDAO(),
		minerRDB:    redis.NewMinerRDB(),
		userFarmDAO: relationdao.NewUserFarmDAO(),
	}
}

func (m *MinerService) CreateMiner(ctx context.Context, userID, farmID int, req *dto.CreateMinerReq) error {
	if !m.validPerm(ctx, userID, farmID, []perm.FarmPerm{perm.FarmOwner, perm.FarmManager}) {
		return errors.New("permission denied")
	}

	rigID, err := m.generateRigID(ctx, 8)
	if err != nil {
		return err
	}
	pass, err := m.generateRigPass(ctx, 8)
	if err != nil {
		return err
	}

	miner := &model.Miner{
		Name:  req.Name,
		RigID: rigID,
		Pass:  pass,
	}
	if err := m.minerDAO.CreateMiner(ctx, userID, farmID, miner); err != nil {
		return err
	}

	minerInfo := &info.Miner{
		HiveOsConfig: utils.HiveOsConfig{
			HiveOsUrl:     utils.GenerateHiveOsUrl(),
			ApiHiveOsUrls: utils.GenerateHiveOsUrl(),
			WorkerName:    req.Name,
			FarmID:        strconv.Itoa(farmID),
			RigID:         rigID,
			RigPasswd:     pass,
		},
	}

	// redis 缓存 miner 配置
	if err := m.minerRDB.CreateMinerByRigID(ctx, rigID, minerInfo); err != nil {
		if err := m.minerDAO.DelMiner(ctx, userID, miner.ID); err != nil {
			return errors.New("set cached failed and del failed")
		}
		return errors.New("set cached failed")
	}

	return err
}

func (m *MinerService) DelMiner(ctx context.Context, userID, farmID, minerID int) error {
	if !m.validPerm(ctx, userID, farmID, []perm.FarmPerm{perm.FarmOwner, perm.FarmManager}) {
		return errors.New("permission denied")
	}
	return m.minerDAO.DelMiner(ctx, userID, minerID)
}

func (m *MinerService) UpdateMiner(ctx context.Context, userID, farmID, minerID int, updateInfo map[string]any) error {
	if !m.validPerm(ctx, userID, farmID, []perm.FarmPerm{perm.FarmOwner, perm.FarmManager}) {
		return errors.New("permission denied")
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
	// if !m.validPerm(ctx, userID, req.FarmID, []perm.FarmPerm{perm.FarmOwner, perm.FarmManager}) {
	// 	return errors.New("permission denied")
	// }

	miner, err := m.minerDAO.GetMinerByMinerID(ctx, userID, minerID)
	if err != nil {
		return err
	}

	minerInfo, err := m.minerRDB.GetMinerByRigID(ctx, miner.RigID)
	if err != nil {
		return errors.New("miner not found")
	}

	minerInfo.HiveOsConfig.Watchdog = req.Watchdog

	if err := m.minerRDB.UpdateMinerByRigID(ctx, miner.RigID, minerInfo); err != nil {
		return err
	}

	return nil
}

func (m *MinerService) UpdateMinerOptions(ctx context.Context, userID int, req *dto.UpdateMinerOptionsReq) error {
	if !m.validPerm(ctx, userID, req.FarmID, []perm.FarmPerm{perm.FarmOwner, perm.FarmManager}) {
		return errors.New("permission denied")
	}

	miner, err := m.minerDAO.GetMinerByMinerID(ctx, userID, req.MinerID)
	if err != nil {
		return err
	}

	minerInfo, err := m.minerRDB.GetMinerByRigID(ctx, miner.RigID)
	if err != nil {
		return err
	}

	minerInfo.HiveOsConfig.Options = req.Options

	if err := m.minerRDB.UpdateMinerByRigID(ctx, miner.RigID, minerInfo); err != nil {
		return err
	}

	return nil
}

func (m *MinerService) UpdateMinerAutofan(ctx context.Context, userID int, req *dto.UpdateMinerAutofanReq) error {
	if !m.validPerm(ctx, userID, req.FarmID, []perm.FarmPerm{perm.FarmOwner, perm.FarmManager}) {
		return errors.New("permission denied")
	}

	miner, err := m.minerDAO.GetMinerByMinerID(ctx, userID, req.MinerID)
	if err != nil {
		return errors.New("miner not found")
	}

	minerInfo, err := m.minerRDB.GetMinerByRigID(ctx, miner.RigID)

	if err := m.minerRDB.UpdateMinerByRigID(ctx, miner.RigID, minerInfo); err != nil {
		return err
	}

	return nil
}

func (m *MinerService) UpdateMinerWallet(ctx context.Context, userID int, req *dto.UpdateMinerWalletReq) error {
	if !m.validPerm(ctx, userID, req.FarmID, []perm.FarmPerm{perm.FarmOwner, perm.FarmManager}) {
		return errors.New("permission denied")
	}

	miner, err := m.minerRDB.GetByID(ctx, req.FarmID, req.MinerID)
	if err != nil {
		return errors.New("miner not found")
	}

	if err := m.minerRDB.Set(ctx, req.FarmID, miner); err != nil {
		return err
	}

	return nil
}

func (m *MinerService) SetWatchdog(ctx context.Context, userID int, req *dto.SetWatchdogReq) error {
	miner, err := m.minerDAO.GetMinerByMinerID(ctx, userID, req.MinerID)
	if err != nil {
		return err
	}
	minerInfo, err := m.minerRDB.GetMinerByRigID(ctx, miner.RigID)
	if err != nil {
		return err
	}
	minerInfo.HiveOsConfig.Watchdog = req.Watchdog
	return m.minerRDB.UpdateMinerByRigID(ctx, miner.RigID, minerInfo)
}

func (m *MinerService) GetWatchdog(ctx context.Context, userID, farmID, minerID int) (*utils.Watchdog, error) {
	miner, err := m.minerDAO.GetMinerByMinerID(ctx, userID, minerID)
	if err != nil {
		return nil, err
	}
	minerInfo, err := m.minerRDB.GetMinerByRigID(ctx, miner.RigID)
	if err != nil {
		return nil, err
	}
	return &minerInfo.HiveOsConfig.Watchdog, nil
}

func (s *MinerService) SetAutoFan(ctx context.Context, req *dto.SetAutoFanReq) error {
	miner, err := s.minerRDB.GetByID(ctx, req.FarmID, req.MinerID)
	if err != nil {
		return err
	}
	miner.HiveOsAutoFan = req.AutoFan
	return s.minerRDB.Set(ctx, req.FarmID, miner)
}

func (s *MinerService) GetAutoFan(ctx context.Context, farmID, minerID int) (*utils.HiveOsAutoFan, error) {
	miner, err := s.minerRDB.GetByID(ctx, farmID, minerID)
	if err != nil {
		return nil, err
	}
	return &miner.HiveOsAutoFan, nil
}

func (m *MinerService) SetOptions(ctx context.Context, userID int, req *dto.SetOptionsReq) error {
	miner, err := m.minerDAO.GetMinerByMinerID(ctx, userID, req.MinerID)
	if err != nil {
		return err
	}
	minerInfo, err := m.minerRDB.GetMinerByRigID(ctx, miner.RigID)
	if err != nil {
		return err
	}
	minerInfo.HiveOsConfig.Options = req.Options
	return m.minerRDB.UpdateMinerByRigID(ctx, miner.RigID, minerInfo)
}

func (m *MinerService) GetOptions(ctx context.Context, userID, farmID, minerID int) (*utils.Options, error) {
	miner, err := m.minerDAO.GetMinerByMinerID(ctx, userID, minerID)
	if err != nil {
		return nil, err
	}
	minerInfo, err := m.minerRDB.GetMinerByRigID(ctx, miner.RigID)
	if err != nil {
		return nil, err
	}
	return &minerInfo.HiveOsConfig.Options, nil
}

func (m *MinerService) GetMinerByMinerID(ctx context.Context, userID, minerID int) (*model.Miner, error) {
	return m.minerDAO.GetMinerByMinerID(ctx, userID, minerID)
}

func (m *MinerService) GetMinersByFarmID(ctx context.Context, farmID int, query map[string]any) (*[]model.Miner, int64, error) {
	return m.minerDAO.GetMinersByFarmID(ctx, farmID, query)
}

func (m *MinerService) GetMiners(ctx context.Context, query map[string]any) (*[]model.Miner, int64, error) {
	return m.minerDAO.GetMiners(ctx, query)
}

func (m *MinerService) ApplyFs(ctx context.Context, userID, farmID, minerID, fsID int) error {
	// todo 检查用户对 miner 的权限
	// 更新 miner cache
	return m.minerDAO.ApplyFs(ctx, minerID, fsID)
}

func (m *MinerService) Transfer(ctx context.Context, farmID, minerID int, toFarmHash string) error {
	// todo 检查用户对 miner 的权限
	return m.minerDAO.Transfer(ctx, farmID, minerID, toFarmHash)
}

func (m *MinerService) validPerm(ctx context.Context, userID, farmID int, allowedPerms []perm.FarmPerm) bool {
	perm, err := m.userFarmDAO.GetPerm(ctx, userID, farmID)
	if err != nil {
		return false
	}
	return slices.Contains(allowedPerms, perm)
}

func (m *MinerService) generateRigID(ctx context.Context, length int) (string, error) {
	if length < 8 {
		return "", errors.New("invalid argument")
	}
	const charset = "123456789" // 以 0 开头时会导致转为字符串与实际 ID 不符合
	id := make([]byte, length)
	for {
		for i := range id {
			num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
			if err != nil {
				return "", err
			}
			id[i] = charset[num.Int64()]
		}
		uid := string(id)
		if !m.minerDAO.ExistsRigID(ctx, uid) {
			return uid, nil
		}
	}
}

func (m *MinerService) generateRigPass(ctx context.Context, length int) (string, error) {
	if length < 8 {
		return "", errors.New("invalid argument")
	}
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	str := make([]byte, length)
	for i := range str {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		str[i] = charset[num.Int64()]
	}
	return string(str), nil
}
