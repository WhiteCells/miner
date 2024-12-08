package service

import "miner/dao/mysql"

type FlightsheetService struct {
	flightsheetDAO      *mysql.FlightsheetDAO
	minerFlightsheetDAO *mysql.MinerFlightsheetDAO
}

func NewFlightsheetService() *FlightsheetService {
	return &FlightsheetService{
		flightsheetDAO:      mysql.NewFlightsheetDAO(),
		minerFlightsheetDAO: mysql.NewMinerFlightsheetDAO(),
	}
}
