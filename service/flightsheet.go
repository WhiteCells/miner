package service

import "miner/dao/mysql"

type FlightsheetService struct {
	flightsheetDAO *mysql.FlightsheetDAO
}

func NewFlightsheetService() *FlightsheetService {
	return &FlightsheetService{
		flightsheetDAO: mysql.NewFlightsheetDAO(),
	}
}
