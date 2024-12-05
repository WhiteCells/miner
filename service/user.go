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
		pointsRecordDAO: mysql.NewPointRecordDAO(),
	}
}

// Register 用户注册
func (s *UserService) Register(ctx *gin.Context, req *dto.RegisterReq) error {
	// 检查用户名是否已存在
	// 需要改成从缓存中查
	if _, _, err := s.userDAO.GetUserByUsername(req.Username); err == nil {
		return errors.New("username already exists")
	}

	// 生成邀请码
	inviteCode := utils.GenerateInviteCode()

	// 生成身份验证密钥
	secret := utils.CreateSecret()

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
			return err
		}
	}

	// 缓存
	err := s.userCache.SetUserInfo(ctx, user)
	if err != nil {
		return errors.New("user cached failed")
	}

	// 创建用户
	return s.userDAO.CreateUser(user)
}

// Login 用户登录
func (s *UserService) Login(ctx *gin.Context, req *dto.LoginReq) (string, int, error) {
	// 先读缓存

	// 获取用户信息
	user, id, err := s.userDAO.GetUserByUsername(req.Username)
	if err != nil {
		return "", -1, errors.New("user not found")
	}

	ctx.Set("user_id", id)
	ctx.Set("user_role", user.Role)

	if !utils.VerifyCodeMoment(user.Secret, req.GoogleCode) {
		return "", -1, errors.New("invalid GoogleCode")
	}

	// 验证密码
	if !s.validatePassword(user, req.Password) {
		return "", -1, errors.New("invalid password")
	}

	// 检查用户状态
	if user.Status != status.UserOn {
		return "", -1, errors.New("account is disabled")
	}

	// 检查积分是否欠费
	if user.Points < 0 {
		return "", -1, errors.New("insufficient points")
	}

	// 检查IP是否变化
	if user.LastLoginIP != "" && user.LastLoginIP != ctx.ClientIP() {
		return "", -1, errors.New("ip")
	}

	// 生成 JWT token
	token, err := utils.GenerateToken(user.ID, user.Name, user.Role, 24)
	if err != nil {
		return "", -1, err
	}

	// 更新登录 IP 信息
	user.LastLoginIP = ctx.ClientIP()
	s.userDAO.UpdateUser(user)

	// 缓存 info token ip
	if err := s.userCache.SetUserInfo(ctx, user); err != nil {
		return "", -1, err
	}
	if err := s.userCache.SetUserToken(ctx, user.ID, token); err != nil {
		return "", -1, err
	}
	if err := s.userCache.SetLoginIP(ctx, user.ID, ctx.ClientIP()); err != nil {
		return "", -1, err
	}

	// todo token 后期不返回
	return token, id, nil
}

// 更新用户信息
func (s *UserService) UpdateUserInfo(ctx *gin.Context, req *dto.UpdateInfoReq) error {
	user, err := s.userDAO.GetUserByID(req.UserID)
	if err != nil {
		return err
	}

	// todo 更新用户信息

	// 保存更新
	if err := s.userDAO.UpdateUser(user); err != nil {
		return err
	}

	// 清除缓存
	return s.userCache.DeleteUserCache(ctx, req.UserID)
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
	return s.userCache.DeleteUserCache(ctx, userID)
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
	}
	return s.AddPoints(ctx, req)
}

// 获取用户信息
func (s *UserService) GetUserInfo(ctx *gin.Context, req *dto.GetUserInfoReq) (*model.User, error) {
	// 先从缓存获取
	user, err := s.userCache.GetUserInfo(ctx, req.UserID)
	if err == nil {
		return user, nil
	}

	// 缓存未命中，从数据库获取
	user, err = s.userDAO.GetUserByID(req.UserID)
	if err != nil {
		return nil, err
	}

	// 更新缓存
	if err := s.userCache.SetUserInfo(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

// 验证密码
func (s *UserService) validatePassword(user *model.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
