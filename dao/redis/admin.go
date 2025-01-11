package redis

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"miner/common/status"
	"miner/model/info"
	"miner/utils"
	"strconv"
)

type AdminRDB struct {
	userRDB     *UserRDB
	farmRDB     *FarmRDB
	minerRDB    *MinerRDB
	SystemRDB   *SystemRDB
	minepoolRDB *MinepoolRDB
}

func NewAdminRDB() *AdminRDB {
	return &AdminRDB{
		userRDB:     NewUserRDB(),
		farmRDB:     NewFarmRDB(),
		minerRDB:    NewMinerRDB(),
		SystemRDB:   NewSystemRDB(),
		minepoolRDB: NewMinpoolRDB(),
	}
}

// 获取所有用户信息
func (c *AdminRDB) GetAllUsers(ctx context.Context) (*[]info.User, error) {
	idInfo, err := utils.RDB.HGetAll(ctx, UserField)
	if err != nil {
		return nil, err
	}
	var users []info.User
	for userID := range idInfo {
		user, err := c.userRDB.GetByID(ctx, userID)
		if err != nil {
			return nil, err
		}
		users = append(users, *user)
	}
	return &users, nil
}

// 获取用户操作日志
// func (c *AdminRDB) GetUserOperLogs(ctx context.Context) (*)

// 获取用户登陆日志
// func (c *AdminRDB) GetUserLoginLogs(ctx context.Context) (*)

// 获取用户积分记录
// func (c *AdminRDB) GetUserPointsRecords(ctx context.Context) (*)

// 获取指定用户的所有矿场
func (c *AdminRDB) GetUserFarms(ctx context.Context, userID string) (*[]info.Farm, error) {
	return c.farmRDB.GetAll(ctx, userID)
}

// 获取指定用户的所有矿机
func (c *AdminRDB) GetUserMiners(ctx context.Context, farmID string) (*[]info.Miner, error) {
	return c.minerRDB.GetAll(ctx, farmID)
}

// +-----------------------+------+
// | key                   | val  |
// +-----------------------+------+
// | admin_reward_invite   | 111  |
// +-----------------------+------+
// | admin_reward_recharge | 111  |
// +-----------------------+------+
// | admin_switch_register | 1    |
// +-----------------------+------+

// 修改注册开关
func (c *AdminRDB) SetSwitchRegister(ctx context.Context, status status.RegisterStatus) error {
	return utils.RDB.Set(ctx, AdminSwitchRegisterField, status)
}

// 获取注册开关
func (c *AdminRDB) GetSwitchRegister(ctx context.Context) (string, error) {
	return utils.RDB.Get(ctx, AdminSwitchRegisterField)
}

// 修改邀请积分奖励数量
func (c *AdminRDB) SetInviteReward(ctx context.Context, reward int) error {
	return utils.RDB.Set(ctx, AdminInviteRewardField, reward)
}

// 获取邀请积分奖励数量
func (c *AdminRDB) GetInviteReward(ctx context.Context) (int, error) {
	rewardStr, err := utils.RDB.Get(ctx, AdminInviteRewardField)
	if err != nil {
		return 0, err
	}
	reward, err := strconv.Atoi(rewardStr)
	if err != nil {
		return 0, err
	}
	return reward, nil
}

// 修改充值积分奖励比例
func (c *AdminRDB) SetRechargeRatio(ctx context.Context, ratio float64) error {
	return utils.RDB.Set(ctx, AdminRechargeRatioField, ratio)
}

// 获取充值积分奖励比例
func (c *AdminRDB) GetRechargeRatio(ctx context.Context) (float64, error) {
	ratioStr, err := utils.RDB.Get(ctx, AdminRechargeRatioField)
	if err != nil {
		return 0, err
	}
	ratio, err := strconv.ParseFloat(ratioStr, 64)
	if err != nil {
		return 0, err
	}
	return ratio, nil
}

// 全局飞行表
// +-----------+---------+------+
// | field     | key     | val  |
// +-----------+---------+------+
// | admin_gfs | <fs_id> | info |
// +-----------+---------+------+

// 设置全局飞行表
func (c *AdminRDB) SetGlobalFs(ctx context.Context, fs *info.Fs) error {
	fsJSON, err := json.Marshal(fs)
	if err != nil {
		return err
	}
	return utils.RDB.HSet(ctx, AdminGfsField, fs.ID, string(fsJSON))
}

// 设置用户状态
func (c *AdminRDB) SetUserStatus(ctx context.Context, userID string, s status.UserStatus) error {
	user, err := c.userRDB.GetByID(ctx, userID)
	if err != nil {
		return err
	}
	user.Status = s
	return c.userRDB.Set(ctx, user)
}

// 设置矿池的消耗
func (c *AdminRDB) SetMinepoolCost(ctx context.Context, mpID string, cost float64) error {
	mp, err := c.minepoolRDB.GetByID(ctx, mpID)
	if err != nil {
		return err
	}
	mp.Cost = cost
	return c.minepoolRDB.Set(ctx, mp)
}

// 助记词
// string
// +-------------------+----------------+
// | key               | val            |
// +-------------------+----------------+
// | mnemonic:active   | <mnemonic_str> |
// +-------------------+----------------+
// list
// +----------------+------------------+
// | key            | val              |
// +----------------+------------------+
// | mnemonic:non   | <mnemonic_str>   |
// +----------------+------------------+

// 设置助记词
// 如果一存在活跃助记词
// 则将该活跃的助记词加入不活跃列表
// 将设置的助记词设置为活跃
func (c *AdminRDB) SetMnemonic(ctx context.Context, mnemonic string) error {
	// 检查助记词
	if !utils.ValidMnemonic(mnemonic) {
		return errors.New("invalid mnemonic")
	}

	activeKey := MakeKey(Mnemonic, Active)
	nonKey := MakeKey(Mnemonic, Non)

	// 助记词加密
	encryptMnemonic, err := c.encryptMnemonic(mnemonic, utils.Config.Mnemonic.Key)
	if err != nil {
		return err
	}

	// 判断是否存在已活跃助记词
	prevMnemonic, err := c.getEncryptMnemonic(ctx)

	// 不存在活跃的助记词
	if err != nil {
		return utils.RDB.Set(ctx, activeKey, encryptMnemonic)
	}

	// 存在活跃的助记词
	pipe := utils.RDB.Client.TxPipeline()

	pipe.RPush(ctx, nonKey, prevMnemonic)
	pipe.Set(ctx, activeKey, encryptMnemonic, 0)

	_, err = pipe.Exec(ctx)

	// 更新
	if err == nil {
		err = utils.UpdateTxWallet(mnemonic)
	}

	return err
}

// 获取加密助记词
func (c *AdminRDB) getEncryptMnemonic(ctx context.Context) (string, error) {
	key := MakeField(Mnemonic, Active)
	return utils.RDB.Get(ctx, key)
}

// 获取活跃助记词
func (c *AdminRDB) GetMnemonic(ctx context.Context) (string, error) {
	key := MakeKey(Mnemonic, Active)
	mnemonic, err := utils.RDB.Get(ctx, key)
	if err != nil {
		return "", err
	}
	// 解密
	mnemonic, err = c.decryptMnemonic(mnemonic, utils.Config.Mnemonic.Key)
	if err != nil {
		return "", err
	}
	return mnemonic, nil
}

// 获取所有助记词
func (c *AdminRDB) GetAllMnemonic(ctx context.Context) (*[]string, error) {
	var mnemonics []string

	// active
	activeKey := MakeKey(Mnemonic, Active)
	activeMnemonic, err := utils.RDB.Get(ctx, activeKey)
	if err != nil {
		return nil, err
	}
	mnemonics = append(mnemonics, activeMnemonic)

	// non
	nonKey := MakeKey(Mnemonic, Non)
	nonMnemonics, err := utils.RDB.LRange(ctx, nonKey)
	if err != nil {
		return nil, err
	}
	mnemonics = append(mnemonics, nonMnemonics...)

	for i := range mnemonics {
		mnemonics[i], err = c.decryptMnemonic(mnemonics[i], utils.Config.Mnemonic.Key)
		if err != nil {
			return nil, err
		}
	}

	return &mnemonics, nil
}

// 加密助记词
func (c *AdminRDB) encryptMnemonic(plaintext, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// 解密助记词
func (c *AdminRDB) decryptMnemonic(ciphertext, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	decoded, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	iv := decoded[:aes.BlockSize]
	decoded = decoded[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(decoded, decoded)
	return string(decoded), nil
}
