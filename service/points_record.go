package service

// import (
// 	"miner/dao/mysql"
// )

// type PointsRecordService struct {
// 	pointslogDAO *mysql.PointslogDAO
// }

// func NewPointRecordService() *PointsRecordService {
// 	return &PointsRecordService{
// 		pointslogDAO: mysql.NewPointRecordDAO(),
// 	}
// }

// // 获取用户积分记录
// // func (s *PointsRecordService) GetPointsRecords(ctx *gin.Context, query map[string]any) ([]model.Pointslog, int64, error) {
// // 	records, total, err := s.pointslogDAO.GetUserPointslog(query)
// // 	if err != nil {
// // 		return nil, -1, errors.New("get user points records failed")
// // 	}
// // 	return records, total, err
// // }
