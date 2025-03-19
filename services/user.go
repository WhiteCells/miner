package services

import (
	"errors"
	"miner/common/dto"
	"miner/common/role"
	"miner/common/status"
	"miner/dao/mysql"
	"miner/model"
	"miner/utils"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	userDAO *mysql.UserDAO
}

func NewUserService() *UserService {
	return &UserService{
		userDAO: mysql.NewUserDAO(),
	}
}

func (m *UserService) Register(ctx *gin.Context, req *dto.RegisterReq) (string, error) {
	// mysql 检查 email 唯一性
	_, err := m.userDAO.GetUserByEmail(req.Email)
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

	address, _ := m.GenerateAddress(ctx, "这里需要使用userid怎么处理")

	_ = &model.User{
		Name:        req.Username,
		Password:    password,
		Secret:      secret,
		Address:     address,
		Email:       req.Email,
		Role:        role.User,
		LastBalance: 0,
		Status:      status.UserOn,
		InviteCode:  "uid",
	}

	return "", nil
}

func (s *UserService) GenerateInviteCode(ctx *gin.Context) (string, error) {
	return "", nil
}

func (s *UserService) GenerateAddress(ctx *gin.Context, userID string) (string, error) {
	if utils.TxWallet == nil {
		// mnemonoic, err := s.adminRDB.GetMnemonic(ctx)
		// if err != nil {
		// 	return "", "", fmt.Errorf("mnemonoic not set")
		// }
		// if err = utils.UpdateTxWallet(mnemonoic); err != nil {
		// 	fmt.Println(mnemonoic)
		// 	return "", "", fmt.Errorf("failed to update TxWallet")
		// }
	}
	account, err := utils.TxWallet.Derive(utils.DerivationPath(userID), false)
	if err != nil {
		return "", err
	}
	// TODO 是否需要存储私钥
	// 测试传出私钥
	_, err = utils.TxWallet.PrivateKeyHex(account)
	if err != nil {
		return "", err
	}
	return account.Address.Hex(), nil
}
