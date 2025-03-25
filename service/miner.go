package service

// import (
// 	"context"
// 	"crypto/rand"
// 	"errors"
// 	"math/big"
// 	"miner/common/dto"
// 	"miner/common/perm"
// 	"miner/dao/redis"
// 	"miner/model/info"
// 	"miner/utils"
// 	"os"
// 	"slices"
// 	"strings"
// )

// type MinerService struct {
// 	minerRDB   *redis.MinerRDB
// 	farmRDB    *redis.FarmRDB
// 	hiveosRDB  *redis.HiveOsRDB
// 	softAllRDB *redis.SoftAllRDB
// }

// func NewMinerService() *MinerService {
// 	return &MinerService{
// 		minerRDB:   redis.NewMinerRDB(),
// 		farmRDB:    redis.NewFarmRDB(),
// 		hiveosRDB:  redis.NewHiveOsRDB(),
// 		softAllRDB: redis.NewSoftAllRDB(),
// 	}
// }

// // CreateMiner 创建矿机
// func (s *MinerService) CreateMiner(ctx context.Context, req *dto.CreateMinerReq) (*info.Miner, error) {
// 	userID, exists := ctx.Value("user_id").(string)
// 	if !exists {
// 		return nil, errors.New("invalid user_id in context")
// 	}
// 	// 检查用户对矿场的权限
// 	if !s.validFarmPerm(ctx, userID, req.FarmID, []perm.FarmPerm{perm.FarmOwner, perm.FarmManager}) {
// 		return nil, errors.New("permission denied")
// 	}

// 	uid, err := utils.GenerateUID()
// 	if err != nil {
// 		return nil, err
// 	}

// 	rigID, err := s.generateRigID(ctx, 8)
// 	if err != nil {
// 		return nil, err
// 	}

// 	pass, err := utils.GeneratePass(8)
// 	if err != nil {
// 		return nil, err
// 	}

// 	hiveOsUrl := utils.GenerateHiveOsUrl()

// 	// 创建矿机
// 	miner := &info.Miner{
// 		ID:    uid,
// 		Name:  req.Name,
// 		RigID: rigID,
// 		Pass:  pass,
// 		Perm:  perm.MinerOwner,
// 		HiveOsConfig: utils.HiveOsConfig{
// 			HiveOsUrl:     hiveOsUrl,
// 			ApiHiveOsUrls: hiveOsUrl,
// 			WorkerName:    req.Name,
// 			FarmID:        req.FarmID,
// 			RigID:         rigID,
// 			RigPasswd:     pass,
// 		},
// 	}

// 	// 创建矿机
// 	if err = s.minerRDB.Set(ctx, req.FarmID, miner); err != nil {
// 		return nil, err
// 	}

// 	// 建立 rigID->{userID:farmID:minerID} 映射
// 	if err = s.hiveosRDB.SetRigMapping(ctx, userID, rigID, req.FarmID, miner.ID); err != nil {
// 		defer func(minerRDB *redis.MinerRDB, ctx context.Context, farmID string, minerID string) {
// 			err := minerRDB.Del(ctx, farmID, minerID)
// 			if err != nil {

// 			}
// 		}(s.minerRDB, ctx, req.FarmID, miner.ID)
// 		return nil, err
// 	}

// 	return miner, err
// }

// // CreateMiner 创建矿机
// func (s *MinerService) CreateMinerByUserID(ctx context.Context, userID string, farmID string, minerName string) (*info.Miner, error) {
// 	// 检查用户对矿场的权限
// 	if !s.validFarmPerm(ctx, userID, farmID, []perm.FarmPerm{perm.FarmOwner, perm.FarmManager}) {
// 		return nil, errors.New("permission denied")
// 	}

// 	uid, err := utils.GenerateUID()
// 	if err != nil {
// 		return nil, err
// 	}

// 	rigID, err := s.generateRigID(ctx, 8)
// 	if err != nil {
// 		return nil, err
// 	}

// 	pass, err := utils.GeneratePass(8)
// 	if err != nil {
// 		return nil, err
// 	}

// 	hiveOsUrl := utils.GenerateHiveOsUrl()

// 	// 创建矿机
// 	miner := &info.Miner{
// 		ID:    uid,
// 		Name:  minerName,
// 		RigID: rigID,
// 		Pass:  pass,
// 		Perm:  perm.MinerOwner,
// 		HiveOsConfig: utils.HiveOsConfig{
// 			HiveOsUrl:     hiveOsUrl,
// 			ApiHiveOsUrls: hiveOsUrl,
// 			WorkerName:    minerName,
// 			FarmID:        farmID,
// 			RigID:         rigID,
// 			RigPasswd:     pass,
// 		},
// 	}

// 	// 创建矿机
// 	if err = s.minerRDB.Set(ctx, farmID, miner); err != nil {
// 		return nil, err
// 	}

// 	// 建立 rigID->{userID:farmID:minerID} 映射
// 	if err = s.hiveosRDB.SetRigMapping(ctx, userID, rigID, farmID, miner.ID); err != nil {
// 		defer s.minerRDB.Del(ctx, farmID, miner.ID)
// 		return nil, err
// 	}

// 	return miner, err
// }

// // DeleteMiner 删除矿机
// func (s *MinerService) DeleteMiner(ctx context.Context, req *dto.DeleteMinerReq) error {
// 	// 检查用户对 Miner 的权限
// 	if !s.validPerm(ctx, req.FarmID, req.MinerID, []perm.MinerPerm{perm.MinerOwner}) {
// 		return errors.New("permission denied")
// 	}
// 	// 删除矿机
// 	if err := s.minerRDB.Del(ctx, req.FarmID, req.MinerID); err != nil {
// 		return errors.New("delete miner failed")
// 	}
// 	return nil
// }

// // UpdateMiner 更新矿机信息
// func (s *MinerService) UpdateMiner(ctx context.Context, req *dto.UpdateMinerReq) error {
// 	if !s.validPerm(ctx, req.FarmID, req.MinerID, []perm.MinerPerm{perm.MinerOwner, perm.MinerManager}) {
// 		return errors.New("permission denied")
// 	}

// 	miner, err := s.minerRDB.GetByID(ctx, req.FarmID, req.MinerID)
// 	if err != nil {
// 		return errors.New("miner not found")
// 	}

// 	// 更新矿机信息
// 	utils.UpdateStructObjFromMap(miner, req.UpdateInfo)

// 	// 保存更新
// 	if err := s.minerRDB.Set(ctx, req.FarmID, miner); err != nil {
// 		return err
// 	}

// 	return nil
// }

// // UpdateMinerWatchdog
// func (s *MinerService) UpdateMinerWatchdog(ctx context.Context, req *dto.UpdateMinerWatchdogReq) error {
// 	if !s.validPerm(ctx, req.FarmID, req.MinerID, []perm.MinerPerm{perm.MinerOwner, perm.MinerManager}) {
// 		return errors.New("permission denied")
// 	}

// 	miner, err := s.minerRDB.GetByID(ctx, req.FarmID, req.MinerID)
// 	if err != nil {
// 		return errors.New("miner not found")
// 	}

// 	utils.UpdateStructObjFromMap(&miner.HiveOsConfig.Watchdog, req.Watchdog)

// 	if err := s.minerRDB.Set(ctx, req.FarmID, miner); err != nil {
// 		return err
// 	}

// 	return nil
// }

// // UpdateMinerOptions
// func (s *MinerService) UpdateMinerOptions(ctx context.Context, req *dto.UpdateMinerOptionsReq) error {
// 	if !s.validPerm(ctx, req.FarmID, req.MinerID, []perm.MinerPerm{perm.MinerOwner, perm.MinerManager}) {
// 		return errors.New("permission denied")
// 	}

// 	miner, err := s.minerRDB.GetByID(ctx, req.FarmID, req.MinerID)
// 	if err != nil {
// 		return errors.New("miner not found")
// 	}

// 	utils.UpdateStructObjFromMap(&miner.HiveOsConfig.Options, req.Options)

// 	if err := s.minerRDB.Set(ctx, req.FarmID, miner); err != nil {
// 		return err
// 	}

// 	return nil
// }

// // UpdateMinerAutofan
// func (s *MinerService) UpdateMinerAutofan(ctx context.Context, req *dto.UpdateMinerAutofanReq) error {
// 	if !s.validPerm(ctx, req.FarmID, req.MinerID, []perm.MinerPerm{perm.MinerOwner, perm.MinerManager}) {
// 		return errors.New("permission denied")
// 	}

// 	miner, err := s.minerRDB.GetByID(ctx, req.FarmID, req.MinerID)
// 	if err != nil {
// 		return errors.New("miner not found")
// 	}

// 	utils.UpdateStructObjFromMap(&miner.HiveOsAutoFan, req.Autofan)

// 	if err := s.minerRDB.Set(ctx, req.FarmID, miner); err != nil {
// 		return err
// 	}

// 	return nil
// }

// // UpdateMinerWallet
// func (s *MinerService) UpdateMinerWallet(ctx context.Context, req *dto.UpdateMinerWalletReq) error {
// 	if !s.validPerm(ctx, req.FarmID, req.MinerID, []perm.MinerPerm{perm.MinerOwner, perm.MinerManager}) {
// 		return errors.New("permission denied")
// 	}

// 	miner, err := s.minerRDB.GetByID(ctx, req.FarmID, req.MinerID)
// 	if err != nil {
// 		return errors.New("miner not found")
// 	}

// 	utils.UpdateStructObjFromMap(&miner.HiveOsWallet, req.Wallet)

// 	if err := s.minerRDB.Set(ctx, req.FarmID, miner); err != nil {
// 		return err
// 	}

// 	return nil
// }

// // GetMinerByID 获取矿机信息
// func (s *MinerService) GetMinerByID(ctx context.Context, farmID string, minerID string) (*info.Miner, error) {
// 	miner, err := s.minerRDB.GetByID(ctx, farmID, minerID)
// 	return miner, err
// }

// // GetMiner 获取用户在矿场的所有矿机
// func (s *MinerService) GetFarmAllMiner(ctx context.Context, farmID string) ([]info.Miner, error) {
// 	miners, err := s.minerRDB.GetAll(ctx, farmID)
// 	return miners, err
// }

// // Transfer 转移矿机
// func (s *MinerService) Transfer(ctx context.Context, req *dto.TransferMinerReq) error {
// 	userID, exists := ctx.Value("user_id").(string)
// 	if !exists {
// 		return errors.New("invalid user_id in context")
// 	}
// 	// 权限检查
// 	if !s.validPerm(ctx, req.FromFarmID, req.MinerID, []perm.MinerPerm{perm.MinerOwner}) {
// 		return errors.New("permission denied")
// 	}
// 	// 转移
// 	if err := s.minerRDB.Transfer(ctx, userID, req.FromFarmID, req.MinerID, req.ToUserID, req.ToFarmID); err != nil {
// 		return errors.New("transfer miner failed")
// 	}
// 	return nil
// }

// // 获取 rig.conf
// func (s *MinerService) GetRigConf(ctx context.Context, farmID string, minerID string) (string, error) {
// 	userID, exists := ctx.Value("user_id").(string)
// 	if !exists {
// 		return "", errors.New("invalid user_id in context")
// 	}
// 	// 权限检查
// 	if !s.validPerm(ctx, farmID, minerID, []perm.MinerPerm{perm.MinerOwner}) {
// 		return "", errors.New("permission denied")
// 	}
// 	// 从 utils/rig.conf 文件中读取内容
// 	templateBytes, err := os.ReadFile("utils/rig.conf")
// 	if err != nil {
// 		return "", err
// 	}
// 	temp := string(templateBytes)

// 	hive_host := utils.GenerateHiveOsUrl()

// 	farm, err := s.farmRDB.GetByID(ctx, userID, farmID)
// 	if err != nil {
// 		return "", err
// 	}

// 	miner, err := s.minerRDB.GetByID(ctx, farmID, minerID)
// 	if err != nil {
// 		return "", err
// 	}

// 	kv := map[string]string{
// 		"HIVE_HOST_URL": hive_host,
// 		"API_HOST_URLs": hive_host,
// 		"RIG_ID":        miner.RigID,
// 		"RIG_PASSWD":    miner.Pass,
// 		"WORKER_NAME":   miner.Name,
// 		"FARM_ID":       farmID,
// 		"TIMEZONE":      farm.TimeZone,
// 	}

// 	for key, val := range kv {
// 		placeholder := "${" + key + "}"
// 		temp = strings.ReplaceAll(temp, placeholder, val)
// 	}

// 	return temp, nil
// }

// // SetWatchdog 设置 watchdog 选项
// func (s *MinerService) SetWatchdog(ctx context.Context, req *dto.SetWatchdogReq) error {
// 	miner, err := s.minerRDB.GetByID(ctx, req.FarmID, req.MinerID)
// 	if err != nil {
// 		return err
// 	}
// 	miner.HiveOsConfig.Watchdog = req.Watchdog
// 	return s.minerRDB.Set(ctx, req.FarmID, miner)
// }

// // GetWatchdog 获取 watchdog 选项
// func (s *MinerService) GetWatchdog(ctx context.Context, farmID string, minerID string) (*utils.Watchdog, error) {
// 	miner, err := s.minerRDB.GetByID(ctx, farmID, minerID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &miner.HiveOsConfig.Watchdog, nil
// }

// // SetAutoFan 设置 fan 选项
// func (s *MinerService) SetAutoFan(ctx context.Context, req *dto.SetAutoFanReq) error {
// 	miner, err := s.minerRDB.GetByID(ctx, req.FarmID, req.MinerID)
// 	if err != nil {
// 		return err
// 	}
// 	miner.HiveOsAutoFan = req.AutoFan
// 	return s.minerRDB.Set(ctx, req.FarmID, miner)
// }

// // GetAutoFan 获取 fan 选项
// func (s *MinerService) GetAutoFan(ctx context.Context, farmID string, minerID string) (*utils.HiveOsAutoFan, error) {
// 	miner, err := s.minerRDB.GetByID(ctx, farmID, minerID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &miner.HiveOsAutoFan, nil
// }

// // SetOptions 设置 worker 选项
// func (s *MinerService) SetOptions(ctx context.Context, req *dto.SetOptionsReq) error {
// 	miner, err := s.minerRDB.GetByID(ctx, req.FarmID, req.MinerID)
// 	if err != nil {
// 		return err
// 	}
// 	miner.HiveOsConfig.Options = req.Options
// 	return s.minerRDB.Set(ctx, req.FarmID, miner)
// }

// // GetOptions 获取 worker 选项
// func (s *MinerService) GetOptions(ctx context.Context, farmID, minerID string) (*utils.Options, error) {
// 	miner, err := s.minerRDB.GetByID(ctx, farmID, minerID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &miner.HiveOsConfig.Options, nil
// }

// // 转移矿机到其他矿场
// // func (s *MinerService) TransferMiner(ctx context.Context, userID, minerID, fromFarmID, toFarmID int) error {
// // 	// 检查源矿场权限
// // 	if !s.checkMinerPermission(userID, fromFarmID, minerID, []perm.MinerPerm{perm.MinerOwner}) {
// // 		return errors.New("permission denied")
// // 	}

// // 	// 检查目标矿场权限
// // 	if !s.farmService.checkFarmPermission(userID, toFarmID, []perm.FarmPerm{perm.FarmOwner, perm.FarmManager}) {
// // 		return errors.New("permission denied for target farm")
// // 	}

// // 	// 更新矿场-矿机关联

// // 	// 清除缓存
// // 	// 更新缓存
// // 	return nil
// // }

// // ApplyFs 矿机应用飞行表
// func (s *MinerService) ApplyFs(ctx context.Context, req *dto.ApplyMinerFsReq) error {
// 	if !s.validPerm(ctx, req.FarmID, req.MinerID, []perm.MinerPerm{perm.MinerOwner, perm.MinerManager}) {
// 		return errors.New("permission denied")
// 	}
// 	softInfo, err := s.softAllRDB.Get(ctx, req.SoftName)
// 	if err != nil {
// 		return errors.New("get soft info err")
// 	}
// 	return s.minerRDB.ApplyFs(ctx, req.FarmID, req.MinerID, req.FsID, softInfo)
// }

// func (s *MinerService) validPerm(ctx context.Context, farmID string, minerID string, allowedPerms []perm.MinerPerm) bool {
// 	farm, err := s.minerRDB.GetByID(ctx, farmID, minerID)
// 	if err != nil {
// 		return false
// 	}
// 	return slices.Contains(allowedPerms, farm.Perm)
// }

// func (s *MinerService) validFarmPerm(ctx context.Context, userID string, farmID string, allowedPerms []perm.FarmPerm) bool {
// 	farm, err := s.farmRDB.GetByID(ctx, userID, farmID)
// 	if err != nil {
// 		return false
// 	}
// 	return slices.Contains(allowedPerms, farm.Perm)
// }

// func (s *MinerService) generateRigID(ctx context.Context, length int) (string, error) {
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
// 		// rigIDMutex.Lock()
// 		if !s.hiveosRDB.ExistsRigID(ctx, uid) {
// 			return uid, nil
// 		}
// 		// rigIDMutex.Unlock()
// 	}
// }
