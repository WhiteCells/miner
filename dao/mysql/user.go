package mysql

import (
	"miner/model"
	"miner/model/info"
	"miner/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserDAO struct{}

func NewUserDAO() *UserDAO {
	return &UserDAO{}
}

func (dao *UserDAO) CreateUser(user *model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return utils.DB.Create(user).Error
}

func (dao *UserDAO) GetUserByID(id int) (*model.User, error) {
	var user model.User
	err := utils.DB.First(&user, id).Error
	return &user, err
}

func (dao *UserDAO) GetUserByName(name string) (*model.User, error) {
	var user model.User
	err := utils.DB.Where("name = ?", name).First(&user).Error
	return &user, err
}

func (dao *UserDAO) GetUserByInviteCode(inviteCode string) (*info.User, error) {
	var user info.User
	err := utils.DB.Where("invite_code = ?", inviteCode).First(&user).Error
	return &user, err
}

func (dao *UserDAO) UpdateUser(user *model.User) error {
	return utils.DB.Save(user).Error
}

func (dao *UserDAO) UpdatePassword(userID int, newPassword string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return utils.DB.Model(&model.User{}).
		Where("id = ?", userID).
		Update("password", string(hashedPassword)).Error
}

func (dao *UserDAO) UpdatePoints(userID int, points int) error {
	return utils.DB.Model(&model.User{}).
		Where("id = ?", userID).
		Update("points", points).Error
}

func (dat *UserDAO) TrasnferPoints(fromUserID int, toUserID int, points uint) error {
	return utils.DB.Transaction(func(tx *gorm.DB) error {
		// 减少源用户积分
		if err := tx.Model(&model.User{}).
			Where("id = ? AND points >= ?", fromUserID, points).
			Update("points", gorm.Expr("points - ?", points)).Error; err != nil {
			return err // 如果扣分失败，则回滚
		}

		// 增加目标用户积分
		if err := tx.Model(&model.User{}).
			Where("id = ?", toUserID).
			Update("points", gorm.Expr("points + ?", points)).Error; err != nil {
			return err // 如果加分失败，则回滚
		}

		return nil
	})
}
