package service

import (
	"errors"
	"miner/common/dto"
	"miner/common/role"
	"miner/common/status"
	"miner/dao/mysql"
	"miner/dao/redis"
	"miner/model/info"
	"miner/utils"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userDAO         *mysql.UserDAO
	userRDB         *redis.UserRDB
	operLog         *mysql.OperLogDAO
	pointsRecordDAO *mysql.PointsRecordDAO
}

func NewUserSerivce() *UserService {
	return &UserService{
		userDAO:         mysql.NewUserDAO(),
		userRDB:         redis.NewUserRDB(),
		operLog:         mysql.NewOperLogDAO(),
		pointsRecordDAO: mysql.NewPointRecordDAO(),
	}
}

// Register 用户注册
func (s *UserService) Register(ctx *gin.Context, req *dto.RegisterReq) error {
	// 用户名
	if s.userRDB.ExistsName(ctx, req.Username) {
		return errors.New("user Name " + req.Username + " exists")
	}

	// 邮箱
	if s.userRDB.ExistsEmail(ctx, req.Email) {
		return errors.New("user Email " + req.Email + " exists")
	}

	// 生成邀请码
	// inviteCode := utils.GenerateInviteCode() // 使用 ID 作为邀请码，方便查找用户

	// 生成身份验证密钥
	secret, err := utils.CreateSecret()
	if err != nil {
		return errors.New("failed to create secret")
	}

	uid, err := utils.GenerateUID()
	if err != nil {
		return errors.New("uid create failed")
	}

	// password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("uid create failed")
	}
	password := string(hashPassword)

	user := &info.User{
		ID:         uid,
		Name:       req.Username,
		Password:   password,
		Email:      req.Email,
		Role:       role.User,
		Status:     status.UserOn,
		InviteCode: uid,
		Secret:     secret,
	}

	// 如果有邀请码，处理邀请关系
	if req.InviteCode != "" {
		user.InviteBy = uid

		// 给邀请人增加积分
		// TODO 记录日志
		// err = s.addInvitePoints(ctx, inviter.ID)
		// if err != nil {
		// 	return errors.New("add invite points failed")
		// }
	}

	// 缓存
	if err = s.userRDB.Set(ctx, user); err != nil {
		return errors.New("user cached failed")
	}

	return nil
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
		return "", nil, errors.New("RDB failed")
	}

	// // 缓存 info token
	// if err := s.userCache.SetUser(ctx, user); err != nil {
	// 	return "", nil, err
	// }
	// if err := s.userCache.SetUserTokenByID(ctx, user.ID, token); err != nil {
	// 	return "", nil, err
	// }

	return token, user, nil
}

// Logout 用户注销
func (s *UserService) Logout(ctx *gin.Context) error {
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

// 更新密码
// func (s *UserService) UpdatePassword(ctx *gin.Context, userID int, oldPassword, newPassword string) error {
// 	user, err := s.userDAO.GetUserByID(userID)
// 	if err != nil {
// 		return err
// 	}

// 	// 验证旧密码
// 	if !s.validatePassword(user, oldPassword) {
// 		return errors.New("invalid old password")
// 	}

// 	// 更新密码
// 	if err := s.userDAO.UpdatePassword(userID, newPassword); err != nil {
// 		return err
// 	}

// 	// 清除缓存
// 	return s.userCache.DeleteUserCache(ctx)
// }

// GetPointsBalance 获取用户积分余额
func (s *UserService) GetPointsBalance(ctx *gin.Context) (int, error) {
	userID, exists := ctx.Value("user_id").(int)
	if !exists {
		return -1, errors.New("invalid user_id in context")
	}
	user, err := s.userDAO.GetUserByID(userID)
	if err != nil {
		return -1, errors.New("user not found")
	}
	return user.Points, err
}

// 添加积分
// func (s *UserService) UpdatePoints(ctx *gin.Context, req *dto.AddPointsReq) error {
// 	user, err := s.userRDB.GetByID(ctx, req.UserID)
// 	if err != nil {
// 		return err
// 	}
// 	if err := s.userRDB.UpdatePoints(ctx, req.UserID, req.Point); err != nil {
// 		return err
// 	}

// 	pointsRecord := &model.PointsRecord{
// 		UserID:  req.UserID,
// 		Type:    req.Type,
// 		Amount:  req.Point,
// 		Balance: newPoint,
// 		Time:    req.Time,
// 	}
// 	return s.pointsRecordDAO.CreatePointsRecord(pointsRecord)
// }

// 添加邀请积分
// func (s *UserService) addInvitePoints(ctx *gin.Context, inviterID string) error {
// 	// 邀请奖励积分
// 	// to
// 	const invitePoints int = 100
// 	req := &dto.AddPointsReq{
// 		UserID: inviterID,
// 		Type:   "invite",
// 		Point:  invitePoints,
// 		Time:   time.Now(),
// 	}
// 	return s.UpdatePoints(ctx, req)
// }

// 获取用户信息
// func (s *UserService) GetUserInfo(ctx *gin.Context) (*model.User, error) {
// 	if user, err := s.userCache.GetUserInfoByID(ctx, req.UserID); err == nil {
// 		return user, nil
// 	}
// 	return s.userDAO.GetUserByID(req.UserID)
// }

// 验证密码
func (s *UserService) validPassword(user *info.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
