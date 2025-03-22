package services

import (
	"context"
	"errors"
	"fmt"
	"miner/common/dto"
	"miner/common/points"
	"miner/common/role"
	"miner/common/status"
	"miner/dao/mysql"
	"miner/dao/redis"
	"miner/model"
	"miner/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	userDAO      *mysql.UserDAO
	userRDB      *redis.UserRDB
	adminDAO     *mysql.AdminDAO
	adminRDB     *redis.AdminRDB
	pointslogDAO *mysql.PointslogDAO
}

func NewUserService() *UserService {
	return &UserService{
		userDAO:  mysql.NewUserDAO(),
		userRDB:  redis.NewUserRDB(),
		adminDAO: mysql.NewAdminDAO(),
		adminRDB: redis.NewAdminRDB(),
	}
}

// 登录
func (m *UserService) Login(ctx context.Context, clientIP string, req *dto.LoginReq) ([]string, string, *model.User, error) {
	user, err := m.userDAO.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, "", nil, err
	}

	if !utils.ValidPassword(user.Password, req.Password) {
		return nil, "", nil, errors.New("wrong password")
	}

	if user.Status != status.UserOn {
		return nil, "", nil, errors.New("account is disabled")
	}

	if user.RechargePoints+user.InvitePoints < 0 {
		return nil, "", nil, errors.New("insufficient points")
	}

	token, err := utils.GenerateToken(user.ID, user.Name, 24)
	if err != nil {
		return nil, "", nil, errors.New("failed to generate token")
	}

	user.LastLoginIP = clientIP
	user.LastLoginAt = time.Now()
	var permissions []string
	if user.Role == role.Admin {
		permissions = []string{"*:*:*"}
	}

	return permissions, token, user, nil
}

func (s *UserService) Logout(ctx *gin.Context) error {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		return errors.New("authorization header is missing")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return errors.New("invalid authorization format")
	}
	token := parts[1]

	if err := s.userRDB.AddBanToken(ctx, token); err != nil {
		return fmt.Errorf("failed to logout: %w", err)
	}

	return nil
}

// 注册
func (m *UserService) Register(ctx *gin.Context, req *dto.RegisterReq) (string, error) {
	_, err := m.userDAO.GetUserByEmail(ctx, req.Email)
	if err == nil {
		return "", errors.New("user email exists")
	}

	password, err := utils.EncryptPassword(req.Password)
	if err != nil {
		return "", errors.New("encrypt pasword error " + err.Error())
	}

	secret, err := utils.CreateSecret()
	if err != nil {
		return "", errors.New("failed to create secret " + err.Error())
	}

	uid, err := utils.GenerateUID()
	if err != nil {
		return "", err
	}

	// 取出助记词
	mn, err := m.adminRDB.GetMnemonic(ctx)
	if err != nil {
		return "", err
	}

	address, key, err := utils.GenerateAddress(mn, uid)
	if err != nil {
		return "", err
	}

	user := &model.User{
		Name:        req.Username,
		Password:    password,
		Secret:      secret,
		Address:     address,
		Email:       req.Email,
		Role:        role.User,
		LastBalance: 0,
		Status:      status.UserOn,
		UID:         uid,
		InviteCode:  uid,
		Key:         key,
	}

	if err := m.userDAO.CreateUser(ctx, user); err != nil {
		return "", err
	}

	// 处理 invite
	if req.InviteCode != "" {
		user.InvitedBy = req.InviteCode
		if err := m.addInvitePoints(ctx, user.ID, req.InviteCode); err != nil {
			return secret, errors.New("failed to add invite point")
		}
		if err := m.userDAO.UpdateUser(ctx, user); err != nil {
			return secret, errors.New("failed to update user")
		}
	}

	return secret, nil
}

func (s *UserService) addInvitePoints(ctx *gin.Context, userID int, inviteCode string) error {
	invitePoints, err := s.adminDAO.GetInviteReward(ctx)
	if err != nil {
		return err
	}
	user, err := s.userDAO.GetUserByInviteCode(ctx, inviteCode)
	if err != nil {
		return err
	}
	user.InvitePoints += invitePoints
	detail := fmt.Sprintf("%d invite %d", user.ID, userID)
	go func() {
		record := &model.Pointslog{
			UserID:  userID,
			Type:    points.PointInvite,
			Amount:  invitePoints,
			Balance: user.InvitePoints,
			Time:    time.Now(),
			Detail:  detail,
		}
		s.pointslogDAO.CreatePointslog(ctx, record)
	}()
	return s.userDAO.UpdateUser(ctx, user)
}

func (m *UserService) GetUserAddress(ctx context.Context, userID int) (string, error) {
	return m.userDAO.GetUserAddress(ctx, userID)
}

func (m *UserService) GetUserPointsBalance(ctx context.Context, userID int) (float32, error) {
	return m.userDAO.GetUserPointsBalance(ctx, userID)
}

func (m *UserService) GetUserOperlogs(ctx context.Context, userID int, query map[string]any) (*[]model.Operlog, error) {
	return m.userDAO.GetUserOperlogs(ctx, userID, query)
}

func (m *UserService) GetUserPointslog(ctx context.Context, userID int, query map[string]any) (*[]model.Pointslog, error) {
	return m.userDAO.GetUserPointslogs(ctx, userID, query)
}

func (m *UserService) QuitFarm(ctx context.Context, userID int, farmID int) error {

	return nil
}
