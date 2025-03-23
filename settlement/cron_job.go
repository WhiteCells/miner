package settlement

import (
	"context"
	"fmt"
	"math"
	"miner/common/perm"
	"miner/common/points"
	"miner/common/status"
	"miner/dao/mysql"
	"miner/dao/mysql/relationdao"
	"miner/dao/redis"
	"miner/model"
	"miner/utils"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

func InitCronJob() {
	c := cron.New()
	c.AddFunc("45 9 * * *", func() {
		ctx := context.Background()
		processPointsDeduct(ctx)
	})
	c.Start()
}

var (
	adminDAO     *mysql.AdminDAO          = mysql.NewAdminDAO()
	adminRDB     *redis.AdminRDB          = redis.NewAdminRDB()
	farmDAO      *mysql.FarmDAO           = mysql.NewFarmDAO()
	userFarmDAO  *relationdao.UserFarmDAO = relationdao.NewUserFarmDAO()
	userDAO      *mysql.UserDAO           = mysql.NewUserDAO()
	pointslogDAO *mysql.PointslogDAO      = mysql.NewPointRecordDAO()
)

func processPointsDeduct(ctx context.Context) {
	users, err := adminDAO.GetAllUsers(ctx)
	if err != nil {
		return
	}

	workerCnt := 10
	userChan := make(chan model.User, len(*users))
	var wg sync.WaitGroup

	for range workerCnt {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for user := range userChan {
				if user.Status == status.UserOn {
					processUserPoints(ctx, &user)
				}
			}
		}()
	}

	for _, user := range *users {
		userChan <- user
	}
	close(userChan)

	wg.Wait()
}

func processUserPoints(ctx context.Context, user *model.User) {
	userLock := getUserLock(user.ID)
	userLock.Lock()
	defer userLock.Unlock()

	if user.LastCheckAt.IsZero() {
		user.LastCheckAt = time.Now().Add(-24 * time.Hour)
	}

	now := time.Now()

	hours := now.Sub(user.LastCheckAt).Hours()
	days := int(math.Ceil(hours / 24))
	if hours < 24 {
		days = 1
	}

	if user.LastCheckAt.Before(now.Add(-24*time.Hour)) ||
		user.LastCheckAt.Before(time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, now.Location())) {

		farms, err := farmDAO.GetAllFarmsByUserID(ctx, user.ID)
		if err != nil {
			log := fmt.Sprintf("%d failed to get all farm", user.ID)
			utils.Logger.Error(log)
			return
		}

		gpuNum := 0
		for _, farm := range *farms {
			p, err := userFarmDAO.GetPerm(ctx, user.ID, farm.ID)
			if err != nil {
				continue
			}
			if p == perm.FarmOwner {
				gpuNum += farm.GpuNum
			}
		}

		freeGpuNum, err := adminRDB.GetFreeGpuNum(ctx)
		if err != nil {
			utils.Logger.Error("admin rdb get free gpu num error")
			return
		}
		if gpuNum <= freeGpuNum {
			utils.Logger.Info("gpu num is less than free gpu num")
			return
		}

		// 计算消耗积分
		num := min(gpuNum, 6)
		// cost := float32(gpuNum-freeGpuNum) * price[num-1] / 30 * float32(days) * discount[gpuNum]
		cost := float32(gpuNum-freeGpuNum) * price[num-1] / 30 * float32(days)

		balance := user.InvitePoints + user.RechargePoints

		// 更新积分
		if err = userDAO.UpdatePoints(ctx, user.ID, -cost, points.PointSettlement); err != nil {
			utils.Logger.Error("update points error")
			return
		}

		// 更新扣除积分时间
		if err = userDAO.UpdateLastCheckAt(ctx, user.ID, now); err != nil {
			utils.Logger.Error("set last check at error")
			return
		}

		user, err = userDAO.GetUserByID(ctx, user.ID)
		if err != nil {
			return
		}

		detail := fmt.Sprintf("farm num:%d gpu num:%d", len(*farms), gpuNum)
		pointslog := &model.Pointslog{
			UserID:  user.ID,
			Type:    points.PointSettlement,
			Amount:  -cost,
			Balance: balance - cost,
			Time:    now,
			Detail:  detail,
		}
		if err := pointslogDAO.CreatePointslog(ctx, pointslog); err != nil {
			log := fmt.Sprintf("%d failed to create points logs", user.ID)
			utils.Logger.Error(log)
		}
	}
}
