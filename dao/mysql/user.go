package mysql

import (
	"context"
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

func (UserDAO) CreateUser(ctx context.Context, user *model.User) error {
	return utils.DB.WithContext(ctx).
		Create(user).Error
}

func (UserDAO) DelUser(ctx context.Context, userID int) error {
	return utils.DB.WithContext(ctx).
		Where("id=?", userID).
		Delete(model.User{}).Error
}

func (UserDAO) UpdateUser(ctx context.Context, user *model.User) error {
	return utils.DB.WithContext(ctx).
		Save(user).Error
}

func (UserDAO) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	var user model.User
	err := utils.DB.WithContext(ctx).
		First(&user, id).Error
	return &user, err
}

func (UserDAO) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := utils.DB.WithContext(ctx).
		Where("email=?", email).
		First(&user).Error
	return &user, err
}

func (UserDAO) GetUserByName(ctx context.Context, name string) (*model.User, error) {
	var user model.User
	err := utils.DB.WithContext(ctx).
		Where("name = ?", name).
		First(&user).Error
	return &user, err
}

func (UserDAO) GetUserByInviteCode(ctx context.Context, inviteCode string) (*model.User, error) {
	var user model.User
	err := utils.DB.WithContext(ctx).
		Where("invite_code = ?", inviteCode).
		First(&user).Error
	return &user, err
}

func (UserDAO) GetAllUsers(ctx context.Context) (*[]model.User, error) {
	var user []model.User
	err := utils.DB.WithContext(ctx).
		Find(&user).Error
	return &user, err
}

func (UserDAO) GetUserAddress(ctx context.Context, userID int) (string, error) {
	var user model.User
	err := utils.DB.WithContext(ctx).First(&user, userID).Error
	return user.Address, err
}

func (UserDAO) GetUserPointsBalance(ctx context.Context, userID int) (float32, error) {
	var user model.User
	err := utils.DB.WithContext(ctx).
		First(&user, userID).Error
	return user.InvitePoints + user.RechargePoints, err
}

func (UserDAO) GetUserOperlogs(ctx context.Context, userID int, query map[string]any) (*[]model.Operlog, error) {
	var operlogs []model.Operlog
	err := utils.DB.WithContext(ctx).
		Where("user_id = ?", userID).
		Find(&operlogs).Error
	return &operlogs, err
}

func (UserDAO) GetUserPointslogs(ctx context.Context, userID int, query map[string]any) (*[]model.Pointslog, error) {
	var pointslogs []model.Pointslog
	err := utils.DB.WithContext(ctx).
		Where("user_id = ?", userID).
		Find(&pointslogs).Error
	return &pointslogs, err
}

func (UserDAO) UpdatePassword(ctx context.Context, userID int, newPassword string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return utils.DB.WithContext(ctx).Model(&model.User{}).
		Where("id = ?", userID).
		Update("password", string(hashedPassword)).Error
}

func (UserDAO) UpdatePoints(ctx context.Context, userID int, num float32, t points.PointsType) error {
	var err error
	switch t {
	case points.PointInvite:
		err = utils.DB.WithContext(ctx).Model(&model.User{}).
			Where("id=?", userID).
			Update("invite_points", num).Error
	case points.PointRecharge:
		err = utils.DB.WithContext(ctx).Model(&model.User{}).
			Where("id=?", userID).
			Update("recharge_points", num).Error
	}
	return err
}

func (UserDAO) UpdateLastCheckAt(ctx context.Context, userID int, t time.Time) error {
	var user model.User
	if err := utils.DB.WithContext(ctx).Where("id=?", userID).
		First(&user).Error; err != nil {
		return err
	}
	return utils.DB.WithContext(ctx).Model(&model.User{}).
		Update("check_at", t).Error
}

func (UserDAO) TransferPoints(ctx context.Context, fromUserID int, toUserID int, points float32) error {
	return utils.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
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
