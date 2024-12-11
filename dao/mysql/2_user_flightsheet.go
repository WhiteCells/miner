package mysql

import (
	"miner/model"
	"miner/utils"
)

type UserFlightsheetDAO struct{}

func NewUserFlightsheetDAO() *UserFlightsheetDAO {
	return &UserFlightsheetDAO{}
}

func (dao *UserFlightsheetDAO) CreateUserFlightsheet(userFlightsheet *model.UserFlightsheet) error {
	return utils.DB.Create(userFlightsheet).Error
}

func (dao *UserFlightsheetDAO) DeleteUserFlightsheet(userID int, userFlightsheetID int) error {
	return utils.DB.
		Where("user_id = ? AND flightsheet_id = ?", userID, userFlightsheetID).
		Delete(model.UserFlightsheet{}).
		Error
}

func (dao *UserFlightsheetDAO) UpdateUserFlightsheet(userFlightsheet *model.UserFlightsheet) error {
	return utils.DB.Save(userFlightsheet).Error
}

func (dao *UserFlightsheetDAO) GetUserAllFlightsheet(userID int) (*[]model.Flightsheet, error) {
	var flightsheets *[]model.Flightsheet
	err := utils.DB.Where("user_id = ?", userID).First(&flightsheets).Error
	return flightsheets, err
}
