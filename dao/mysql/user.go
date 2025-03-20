package mysql

import (
	"miner/common/points"
	"miner/model"
	"miner/utils"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserDAO struct{}

func NewUserDAO() *UserDAO {
	return &UserDAO{}
}

func (UserDAO) CreateUser(user *model.User) error {
	return utils.DB.
		Create(user).Error
}

func (UserDAO) DelUser(userID int) error {
	return utils.DB.
		Where("id=?", userID).
		Delete(model.User{}).Error
}

func (UserDAO) UpdateUser(user *model.User) error {
	return utils.DB.
		Save(user).Error
}

func (UserDAO) GetUserByID(id int) (*model.User, error) {
	var user model.User
	err := utils.DB.
		First(&user, id).Error
	return &user, err
}

func (UserDAO) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := utils.DB.
		Where("email=?", email).
		First(&user).Error
	return &user, err
}

func (UserDAO) GetUserByName(name string) (*model.User, error) {
	var user model.User
	err := utils.DB.
		Where("name = ?", name).
		First(&user).Error
	return &user, err
}

func (UserDAO) GetUserByInviteCode(inviteCode string) (*model.User, error) {
	var user model.User
	err := utils.DB.
		Where("invite_code = ?", inviteCode).
		First(&user).Error
	return &user, err
}

func (UserDAO) GetAllUsers() (*[]model.User, error) {
	var user []model.User
	err := utils.DB.
		Find(&user).Error
	return &user, err
}

func (UserDAO) GetUserAddress(userID int) (string, error) {
	var user model.User
	err := utils.DB.First(&user, userID).Error
	return user.Address, err
}

func (UserDAO) GetUserPointsBalance(userID int) (float32, error) {
	var user model.User
	err := utils.DB.First(&user, userID).Error
	return user.InvitePoints + user.RechargePoints, err
}

func (UserDAO) UpdatePassword(userID int, newPassword string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return utils.DB.Model(&model.User{}).
		Where("id = ?", userID).
		Update("password", string(hashedPassword)).Error
}

func (UserDAO) UpdatePoints(userID int, num float32, t points.PointsType) error {
	// return utils.DB.Model(&model.User{}).
	// 	Where("id = ?", userID).
	// 	Update("points", points).Error
	// var err error
	// switch t {
	// case points.PointInvite:
	// 	err = utils.DB.
	// case points.PointRecharge:

	// }

	return nil
}

func (UserDAO) UpdateLastCheckAt(userID int, t time.Time) error {
	var user model.User
	if err := utils.DB.Where("id=?", userID).First(&user).Error; err != nil {
		return err
	}
	return utils.DB.Model(&model.User{}).Update("check_at", t).Error
}

func (UserDAO) TransferPoints(fromUserID int, toUserID int, points float32) error {
	return utils.DB.Transaction(func(tx *gorm.DB) error {
		// 减少源用户积分
		if err := tx.Model(&model.User{}).
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("id = ?", fromUserID).
			Update("points", gorm.Expr("points - ?", points)).Error; err != nil {
			return err
		}
		// 增加目标用户积分
		if err := tx.Model(&model.User{}).
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("id = ?", toUserID).
			Update("points", gorm.Expr("points + ?", points)).Error; err != nil {
			return err
		}
		return nil
	})
}
