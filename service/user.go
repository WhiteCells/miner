package service

import (
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
	operLogDAO      *mysql.OperLogDAO
	pointsRecordDAO *mysql.PointsRecordDAO
}

func NewUserSerivce() *UserService {
	return &UserService{
		userDAO:         mysql.NewUserDAO(),
		userRDB:         redis.NewUserRDB(),
		adminRDB:        redis.NewAdminRDB(),
		operLogDAO:      mysql.NewOperLogDAO(),
		pointsRecordDAO: mysql.NewPointRecordDAO(),
	}
}

// Register 用户注册
func (s *UserService) Register(ctx *gin.Context, req *dto.RegisterReq) (string, error) {
	// 用户名
	if s.userRDB.ExistsName(ctx, req.Username) {
		return "", errors.New("user Name " + req.Username + " exists")
	}

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
		LastBalance: 0.0,
		Status:      status.UserOn,
		InviteCode:  uid,
		Key:         key,
	}

	// 如果有邀请码，处理邀请关系
	if req.InviteCode != "" {
		user.InviteBy = uid
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
func (s *UserService) Login(ctx *gin.Context, req *dto.LoginReq) (string, *info.User, error) {
	user, err := s.userRDB.GetByName(ctx, req.Username)
	if err != nil {
		return "", nil, errors.New("user not found")
	}

	// 验证 Google 验证码
	// if ret, err := utils.VerifyCodeMoment(user.Secret, req.GoogleCode); ret || err != nil {
	// 	return "", nil, errors.New("invalid GoogleCode")
	// }

	// 验证密码
	if !s.validPassword(user, req.Password) {
		return "", nil, errors.New("invalid password")
	}

	// 检查用户状态
	if user.Status != status.UserOn {
		return "", nil, errors.New("account is disabled")
	}

	// 检查积分是否欠费
	if user.RechargePoints+user.InvitePoints < 0 {
		return "", nil, errors.New("insufficient points")
	}

	// 检查IP是否变化
	if user.LastLoginIP != "" && user.LastLoginIP != ctx.ClientIP() {
		return "", nil, errors.New("new ip detected")
	}

	// 生成 JWT token
	token, err := utils.GenerateToken(user.ID, user.Name, 24)
	if err != nil {
		return "", nil, err
	}

	// 更新登录 IP 信息
	user.LastLoginIP = ctx.ClientIP()
	user.LastLoginAt = time.Now()

	ctx.Set("user_id", user.ID)

	if err := s.userRDB.Set(ctx, user); err != nil {
		return "", nil, errors.New("failed to update")
	}

	return token, user, nil
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

// 更新用户信息
func (s *UserService) UpdateUserInfo(ctx *gin.Context, req *dto.UpdateInfoReq) error {
	userID, exists := ctx.Value("user_id").(string)
	if !exists {
		return errors.New("invalid user_id in context")
	}
	user, err := s.userRDB.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	return s.userRDB.Set(ctx, user)
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
	// 邀请所获得的积分 + 充值所获的积分
	points := user.InvitePoints + user.RechargePoints
	return points, nil
}

// 获取用户信息
func (s *UserService) GetUserInfo(ctx *gin.Context, userID string) (*info.User, error) {
	return s.userRDB.GetByID(ctx, userID)
}

// 获取用户充值地址
// func (s *UserService) GetUserAddress(ctx *gin.Context, userID string) (string, error) {
// 	return s.userRDB.
// }

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
		record := &model.PointsRecord{
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
