package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"miner/common/dto"
	"miner/common/points"
	"miner/common/role"
	"miner/common/status"
	"miner/dao/mysql"
	"miner/dao/redis"
	"miner/model"
	"miner/model/info"
	"miner/utils"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userDAO         *mysql.UserDAO
	userRDB         *redis.UserRDB
	adminRDB        *redis.AdminRDB
	coinRDB         *redis.CoinRDB
	poolRDB         *redis.PoolRDB
	softRDB         *redis.SoftRDB
	operLogDAO      *mysql.OperLogDAO
	pointsRecordDAO *mysql.PointsRecordDAO
	bscApiKeyRDB    *redis.BscApiKeyRDB
}

func NewUserSerivce() *UserService {
	return &UserService{
		userDAO:         mysql.NewUserDAO(),
		userRDB:         redis.NewUserRDB(),
		adminRDB:        redis.NewAdminRDB(),
		coinRDB:         redis.NewCoinRDB(),
		poolRDB:         redis.NewPoolRDB(),
		softRDB:         redis.NewSoftRDB(),
		operLogDAO:      mysql.NewOperLogDAO(),
		pointsRecordDAO: mysql.NewPointRecordDAO(),
		bscApiKeyRDB:    redis.NewBscApiKeyRDB(),
	}
}

// Register 用户注册
func (s *UserService) Register(ctx *gin.Context, req *dto.RegisterReq) (string, error) {
	// 邮箱
	if s.userRDB.ExistsEmail(ctx, req.Email) {
		return "", errors.New("user Email " + req.Email + " exists")
	}

	// 用户 ID
	uid := s.generateUserID(ctx)

	// 用户 password
	password, err := utils.EncryptPassword(req.Password)
	if err != nil {
		return "", err
	}

	// 生成身份验证密钥
	secret, err := utils.CreateSecret()
	if err != nil {
		return "", errors.New("failed to create secret")
	}

	// 交易地址
	address, key, _ := s.GenerateAddress(ctx, uid)

	user := &info.User{
		ID:          uid,
		Name:        req.Username,
		Password:    password,
		Secret:      secret,
		Address:     address,
		Email:       req.Email,
		Role:        role.User,
		LastBalance: 0,
		Status:      status.UserOn,
		InviteCode:  uid,
		Key:         key,
	}

	// 如果有邀请码，处理邀请关系
	if req.InviteCode != "" {
		user.InviteBy = req.InviteCode
		// 给邀请人增加积分
		if err = s.addInvitePoints(ctx, uid, req.InviteCode); err != nil {
			return "", err
		}
	}

	// 存储用户
	if err = s.userRDB.Set(ctx, user); err != nil {
		return "", errors.New("user cached failed")
	}

	return secret, nil
}

// Login 用户登录
func (s *UserService) Login(ctx *gin.Context, req *dto.LoginReq) ([]string, string, *info.User, error) {
	user, err := s.userRDB.GetByEmail(ctx, req.Email)
	if err != nil {
		return []string{""}, "", nil, errors.New("user not found")
	}

	//验证 Google 验证码
	//if ret, err := utils.VerifyCodeMoment(user.Secret, req.GoogleCode); ret || err != nil {
	//	return "", nil, errors.New("invalid GoogleCode")
	//}

	// 验证密码
	if !s.validPassword(user, req.Password) {
		return []string{""}, "", nil, errors.New("wrong password")
	}

	// 验证 Captcha
	// if !utils.VerifyCaptcha(ctx, req.CaptchaID, req.CaptchaValue) {
	// 	return "", nil, errors.New("wrong captcha")
	// }

	// 检查用户状态
	if user.Status != status.UserOn {
		return []string{""}, "", nil, errors.New("account is disabled")
	}

	// 检查积分是否欠费
	if user.RechargePoints+user.InvitePoints < 0 {
		return []string{""}, "", nil, errors.New("insufficient points")
	}

	// 检查IP是否变化
	//if user.LastLoginIP != "" && user.LastLoginIP != ctx.ClientIP() {
	//	return "", nil, errors.New("new ip detected")
	//}

	// 生成 JWT token
	token, err := utils.GenerateToken(user.ID, user.Name, 24)
	if err != nil {
		return []string{""}, "", nil, err
	}

	// 更新登录 IP 信息
	user.LastLoginIP = ctx.ClientIP()
	user.LastLoginAt = time.Now()
	var permissions []string
	if user.ID == "0" {
		permissions = []string{"*:*:*"}
	}

	ctx.Set("user_id", user.ID)

	if err := s.userRDB.Set(ctx, user); err != nil {
		return []string{""}, "", nil, errors.New("failed to update")
	}

	return permissions, token, user, nil
}

// Logout 用户注销
func (s *UserService) Logout(ctx *gin.Context) error {
	// 从请求头中获取 authorization 信息
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		return errors.New("authorization header is missing")
	}

	// 验证 token 格式是否正确
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return errors.New("invalid authorization format")
	}
	token := parts[1]

	// 添加 token 到 ban token
	if err := s.userRDB.AddBanToken(ctx, token); err != nil {
		return fmt.Errorf("failed to logout: %w", err)
	}

	return nil
}

// UpdatePasswd 修改用户密码
func (s *UserService) UpdatePasswd(ctx *gin.Context, req *dto.UpdatePasswdReq) error {
	userID, exists := ctx.Value("user_id").(string)
	if !exists {
		return errors.New("invalid user_id in context")
	}
	user, err := s.userRDB.GetByID(ctx, userID)
	if err != nil {
		return errors.New("user not found")
	}
	if !s.validPassword(user, req.OldPasswd) {
		return errors.New("wrong old password")
	}
	hashPassword, err := utils.EncryptPassword(req.NewPasswd)
	if err != nil {
		return errors.New("failed to encrypt new password")
	}
	user.Password = hashPassword
	if err := s.userRDB.Set(ctx, user); err != nil {
		return errors.New("failed to update user password")
	}
	return nil
}

// GetPointsBalance 获取用户积分余额
func (s *UserService) GetPointsBalance(ctx *gin.Context) (float32, error) {
	userID, exists := ctx.Value("user_id").(string)
	if !exists {
		return -1, errors.New("invalid user_id in context")
	}
	// 查找用户
	user, err := s.userRDB.GetByID(ctx, userID)
	if err != nil {
		return -1, errors.New("user not found")
	}

	// 取出用户的上次代币余额
	lastBalanceF := user.LastBalance

	// 获取代币转换率
	ratio, err := s.adminRDB.GetRechargeRatio(ctx)
	if err != nil {
		return 0, err
	}

	// 获取当前用户代币余额
	curBalance, err := s.getUserCurBalance(ctx, user.Address)
	if err != nil {
		return 0, err
	}

	curBalanceF64, err := strconv.ParseFloat(curBalance, 32)
	if err != nil {
		return 0, errors.New("str conversion failed")
	}

	recharge := float32(curBalanceF64) - lastBalanceF

	// 兑换比率
	addPoints := recharge * float32(ratio)

	// 用户增加积分
	user.RechargePoints += addPoints
	// 更新余额
	user.LastBalance += recharge

	// 更新用户
	s.userRDB.Set(ctx, user)

	// 邀请所获得的积分 + 充值所获的积分
	points := user.InvitePoints + user.RechargePoints

	return points, nil
}

func (s *UserService) getUserCurBalance(ctx context.Context, address string) (string, error) {
	apikey, err := s.bscApiKeyRDB.ZRangeWithScore(ctx)
	if err != nil {
		return "", err
	}
	// 增加 apikey 使用次数
	if err := s.bscApiKeyRDB.ZIncrBy(ctx, apikey, 1); err != nil {
		return "", err
	}
	result, err := s.requestBscApi(address, apikey)
	defer func() {
		// 减少 apikey 使用次数
		if err := s.bscApiKeyRDB.ZIncrBy(ctx, apikey, -1); err != nil {
			return
		}
	}()
	if err != nil {
		return "", err
	}

	return result, nil
}

func (s *UserService) GetCoins(ctx context.Context) (*[]string, error) {
	infos, err := s.coinRDB.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	var coins []string
	for _, info := range *infos {
		coins = append(coins, info.Name)
	}
	return &coins, nil
}

func (s *UserService) GetPools(ctx context.Context, coinName string) (*[]info.Pool, error) {
	return s.poolRDB.GetAll(ctx, coinName)
}

//// AddSoft 应用 custom miner soft 信息
//func (s *UserService) AddSoft(ctx context.Context, name string, soft *info.Soft) error {
//	return s.softRDB.Set(ctx, name, soft)
//}
//
//// DelSoft 删除 custom miner soft 信息
//func (s *UserService) DelSoft(ctx context.Context, name string) error {
//	return s.softRDB.Del(ctx, name)
//}
//
//// Update 修改 custom miner soft 信息
//func (s *UserService) UpdateSoft(ctx context.Context, name string, soft *info.Soft) error {
//	return s.softRDB.Set(ctx, name, soft)
//}
//
//// GetSoft 获取 custom miner soft 信息
//func (s *UserService) GetSoft(ctx context.Context, fsID string) (*info.Soft, error) {
//	return s.softRDB.Get(ctx, fsID)
//}

// requestBscApi 调用 bsc api
func (s *UserService) requestBscApi(address string, apikey string) (string, error) {
	url := fmt.Sprintf(utils.Config.Bsc.Api, address, apikey)
	go func() {
		utils.Logger.Info("generate " + url + "failed")
	}()

	// client := http.Client{
	// 	Timeout: 2 * time.Second,
	// }
	rsp, err := http.Get(url)

	if err != nil {
		return "", fmt.Errorf("HTTP request failed: %v", err)
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("BSC API returned non-200 status code: %d", rsp.StatusCode)
	}

	var body dto.BscApiRspBody
	if err := json.NewDecoder(rsp.Body).Decode(&body); err != nil {
		return "", fmt.Errorf("failed to parse response: %v", err)
	}

	return body.Result, nil
}

// 获取用户信息
func (s *UserService) GetUserInfo(ctx *gin.Context, userID string) (*info.User, error) {
	return s.userRDB.GetByID(ctx, userID)
}

// 获取用户充值地址
func (s *UserService) GetUserAddress(ctx *gin.Context, userID string) (string, error) {
	user, err := s.GetUserInfo(ctx, userID)
	if err != nil {
		return "", err
	}
	return user.Address, nil
}

// 验证密码
func (s *UserService) validPassword(user *info.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

// 给邀请者增加积分
func (s *UserService) addInvitePoints(ctx *gin.Context, uid string, inviterID string) error {
	// 获取邀请积分
	invitePoints, err := s.adminRDB.GetInviteReward(ctx)
	if err != nil {
		return errors.New("invite rewards are not set")
	}
	// 查找邀请人
	user, err := s.userRDB.GetByID(ctx, inviterID)
	if err != nil {
		return errors.New("inviter not long exists")
	}
	// 增加邀请积分
	user.InvitePoints += invitePoints
	// 积分记录
	detail := fmt.Sprintf("%s invite %s", inviterID, uid)
	go func() {
		record := &model.Pointslog{
			UserID:  inviterID,
			Type:    points.PointInvite,
			Amount:  invitePoints,
			Balance: user.InvitePoints,
			Time:    time.Now(),
			Detail:  detail,
		}
		s.pointsRecordDAO.CreatePointsRecord(record)
	}()
	return s.userRDB.Set(ctx, user)
}

// 生成用户 ID
// 原先用户 ID 无法用于生成钱包的账户
// [0, 4294967295]
func (s *UserService) generateUserID(ctx *gin.Context) string {
	for {
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)

		// 生成 [0, 4294967295] 范围内的随机数
		userIDInt := r.Uint32()
		userID := strconv.FormatUint(uint64(userIDInt), 10)

		if s.userRDB.ExistsSameID(ctx, userID) {
			continue
		}

		return userID
	}
}

// 通过助记词为每个用户生成地址
/*
m：主路径，表示从助记词派生出根密钥。
44'：BIP-44 标准，允许使用多种加密货币。
60'：指定以太坊（Ethereum）作为目标货币。
0'：账户索引，指定第一个账户。
0：变化类型，指定生成外部地址。
%d：用户索引，用于区分不同用户生成不同的地址
*/
func (s *UserService) GenerateAddress(ctx *gin.Context, userID string) (string, string, error) {
	if utils.TxWallet == nil {
		mnemonoic, err := s.adminRDB.GetMnemonic(ctx)
		if err != nil {
			return "", "", fmt.Errorf("mnemonoic not set")
		}
		if err = utils.UpdateTxWallet(mnemonoic); err != nil {
			fmt.Println(mnemonoic)
			return "", "", fmt.Errorf("failed to update TxWallet")
		}
	}
	account, err := utils.TxWallet.Derive(utils.DerivationPath(userID), false)
	if err != nil {
		return "", "", err
	}
	// TODO 是否需要存储私钥
	// 测试传出私钥
	privateKey, err := utils.TxWallet.PrivateKeyHex(account)
	if err != nil {
		return "", "", err
	}
	return account.Address.Hex(), privateKey, nil
}

func (s *UserService) GetSoftAll(ctx context.Context, coinName string) (*[]info.Soft, error) {
	return s.softRDB.GetAll(ctx, coinName)
}
