package service

import (
	"errors"
	"miner/common/dto"
	"miner/common/role"
	"miner/common/status"
	"miner/dao/mysql"
	"miner/dao/redis"
	"miner/model"
	"miner/utils"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userDAO         *mysql.UserDAO
	userCache       *redis.UserCache
	operLog         *mysql.OperLogDAO
	pointsRecordDAO *mysql.PointsRecordDAO
}

func NewUserSerivce() *UserService {
	return &UserService{
		userDAO:         mysql.NewUserDAO(),
		userCache:       redis.NewUserCache(),
		operLog:         mysql.NewOperLogDAO(),
		pointsRecordDAO: mysql.NewPointRecordDAO(),
	}
}

// Register 用户注册
func (s *UserService) Register(ctx *gin.Context, req *dto.RegisterReq) error {
	// 用户名
	_, err := s.userCache.GetUserInfoByName(ctx, req.Username)
	if err == nil {
		return errors.New("user " + req.Username + " exists")
	}
	// 邮箱
	_, err = s.userCache.GetUserInfoByEmail(ctx, req.Email)
	if err == nil {
		return errors.New("user Email " + req.Email + " exists")
	}

	// if _, _, err := s.userDAO.GetUserByName(req.Username); err == nil {
	// 	return errors.New("username " + req.Username + " exists")
	// }

	// 生成邀请码
	inviteCode := utils.GenerateInviteCode()

	// 生成身份验证密钥
	secret, err := utils.CreateSecret()
	if err != nil {
		return errors.New("failed to create secret")
	}

	user := &model.User{
		Name:       req.Username,
		Password:   req.Password,
		Email:      req.Email,
		Role:       role.User,
		Status:     status.UserOn,
		InviteCode: inviteCode,
		Secret:     secret,
	}

	// 如果有邀请码，处理邀请关系
	if req.InviteCode != "" {
		inviter, err := s.userDAO.GetUserByInviteCode(req.InviteCode)
		if err != nil {
			return errors.New("invalid invite code")
		}
		user.InvitedBy = inviter.ID

		// 给邀请人增加积分
		err = s.addInvitePoints(ctx, inviter.ID)
		if err != nil {
			return errors.New("add invite points failed")
		}
	}

	// 创建用户
	err = s.userDAO.CreateUser(user)
	if err != nil {
		return errors.New("user create failed")
	}

	// 缓存
	if err = s.userCache.SetUserInfo(ctx, user); err != nil {
		return errors.New("user cached failed")
	}

	return nil
}

// Login 用户登录
func (s *UserService) Login(ctx *gin.Context, req *dto.LoginReq) (string, *model.User, error) {
	// 先读缓存
	user, err := s.userCache.GetUserInfoByName(ctx, req.Username)
	if err != nil {
		user, err = s.userDAO.GetUserByName(req.Username)
		if err != nil {
			return "", nil, errors.New("user not found")
		}
		return "", nil, errors.New("user not found")
	}

	// 验证 Google 验证码
	// if ret, err := utils.VerifyCodeMoment(user.Secret, req.GoogleCode); ret || err != nil {
	// 	return -1, errors.New("invalid GoogleCode")
	// }

	// 验证密码
	if !s.validatePassword(user, req.Password) {
		return "", nil, errors.New("invalid password")
	}

	// 检查用户状态
	if user.Status != status.UserOn {
		return "", nil, errors.New("account is disabled")
	}

	// 检查积分是否欠费
	if user.Points < 0 {
		return "", nil, errors.New("insufficient points")
	}

	// 检查IP是否变化
	if user.LastLoginIP != "" && user.LastLoginIP != ctx.ClientIP() {
		return "", nil, errors.New("ip")
	}

	// 生成 JWT token
	token, err := utils.GenerateToken(user.ID, user.Name, user.Role, 24)
	if err != nil {
		return "", nil, err
	}

	// 更新登录 IP 信息
	user.LastLoginIP = ctx.ClientIP()
	user.LastLoginAt = time.Now()
	s.userDAO.UpdateUser(user)

	// 缓存 info token
	if err := s.userCache.SetUserInfo(ctx, user); err != nil {
		return "", nil, err
	}
	if err := s.userCache.SetUserTokenByID(ctx, user.ID, token); err != nil {
		return "", nil, err
	}

	return token, user, nil
}

// 更新用户信息
func (s *UserService) UpdateUserInfo(ctx *gin.Context, userID int, req *dto.UpdateInfoReq) error {
	user, err := s.userDAO.GetUserByID(userID)
	if err != nil {
		return err
	}

	// todo 更新用户信息

	// 保存更新
	if err := s.userDAO.UpdateUser(user); err != nil {
		return err
	}

	// 清除缓存
	return s.userCache.DeleteUserCache(ctx)
}

// 更新密码
func (s *UserService) UpdatePassword(ctx *gin.Context, userID int, oldPassword, newPassword string) error {
	user, err := s.userDAO.GetUserByID(userID)
	if err != nil {
		return err
	}

	// 验证旧密码
	if !s.validatePassword(user, oldPassword) {
		return errors.New("invalid old password")
	}

	// 更新密码
	if err := s.userDAO.UpdatePassword(userID, newPassword); err != nil {
		return err
	}

	// 清除缓存
	return s.userCache.DeleteUserCache(ctx)
}

// 添加积分
func (s *UserService) AddPoints(ctx *gin.Context, req *dto.AddPointsReq) error {
	user, err := s.userDAO.GetUserByID(req.UserID)
	if err != nil {
		return err
	}
	newPoint := user.Points + req.Point
	if err := s.userDAO.UpdatePoints(req.UserID, newPoint); err != nil {
		return err
	}

	pointsRecord := &model.PointsRecord{
		UserID:  req.UserID,
		Type:    req.Type,
		Amount:  req.Point,
		Balance: newPoint,
		Time:    req.Time,
	}
	return s.pointsRecordDAO.CreatePointsRecord(pointsRecord)
}

// 添加邀请积分
func (s *UserService) addInvitePoints(ctx *gin.Context, inviterID int) error {
	// 邀请奖励积分
	// to
	const invitePoints int = 100
	req := &dto.AddPointsReq{
		UserID: inviterID,
		Type:   "invite",
		Point:  invitePoints,
		Time:   time.Now(),
	}
	return s.AddPoints(ctx, req)
}

// 获取用户信息
func (s *UserService) GetUserInfo(ctx *gin.Context, req *dto.GetUserInfoReq) (*model.User, error) {
	if user, err := s.userCache.GetUserInfoByID(ctx, req.UserID); err == nil {
		return user, nil
	}
	return s.userDAO.GetUserByID(req.UserID)
}

// 验证密码
func (s *UserService) validatePassword(user *model.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
