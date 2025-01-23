package settlement

import (
	"context"
	"math"
	"miner/common/perm"
	"miner/common/points"
	"miner/dao/mysql"
	"miner/dao/redis"
	"miner/model"
	"miner/model/info"
	"miner/utils"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

func InitCronJob() {
	c := cron.New()
	c.AddFunc("0 12 * * *", func() {
		ctx := context.Background()
		processPointsDeduct(ctx)
	})
	c.Start()
}

var (
	adminRDB       *redis.AdminRDB        = redis.NewAdminRDB()
	farmRDB        *redis.FarmRDB         = redis.NewFarmRDB()
	userRDB        *redis.UserRDB         = redis.NewUserRDB()
	pointsCordsDAO *mysql.PointsRecordDAO = mysql.NewPointRecordDAO()
)

func processPointsDeduct(ctx context.Context) {
	users, err := adminRDB.GetAllUsers(ctx)
	if err != nil {
		return
	}

	workerCnt := 10
	userChan := make(chan info.User, len(*users))
	var wg sync.WaitGroup

	for i := 0; i < workerCnt; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for user := range userChan {
				processUserPoints(ctx, &user)
			}
		}()
	}

	for _, user := range *users {
		userChan <- user
	}
	close(userChan)

	wg.Wait()
}

func processUserPoints(ctx context.Context, user *info.User) {
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

		// user 拥有的所有 farm
		farms, err := farmRDB.GetAll(ctx, user.ID)
		if err != nil {
			utils.Logger.Error(user.ID + " farm rdb get all error")
			return
		}

		// 计算所有 farm 的 GPU
		gpuNum := 0
		for _, farm := range *farms {
			// 只计算所有者
			if farm.Perm == perm.FarmOwner {
				gpuNum += farm.GpuNum
			}
		}

		freeGpuNum, err := adminRDB.GetFreeGpuNum(ctx)
		if err != nil {
			utils.Logger.Error("admin rdb get free gpu num error")
			return
		}
		if gpuNum <= freeGpuNum {
			utils.Logger.Info(user.ID + " gpu num is less than free gpu num")
			return
		}

		// 计算消耗积分
		num := min(gpuNum, 6)
		cost := float32(gpuNum-freeGpuNum) * price[num-1] / 30 * float32(days)

		balance := user.InvitePoints + user.RechargePoints

		// 更新积分
		if err = userRDB.UpdatePoints(ctx, user.ID, -cost, points.PointSettlement); err != nil {
			utils.Logger.Error(user.ID + " update points error")
			return
		}

		// 更新扣除积分时间
		if err = userRDB.SetLastCheckAt(ctx, user.ID, now); err != nil {
			utils.Logger.Error(user.ID + " set last check at error")
			return
		}

		user, err = userRDB.GetByID(ctx, user.ID)
		if err != nil {
			utils.Logger.Error(user.ID + " get user error")
			return
		}

		pointsRecord := &model.PointsRecord{
			UserID:  user.ID,
			Type:    points.PointSettlement,
			Amount:  -cost,
			Balance: balance - cost,
			Time:    time.Now(),
			Detail:  "",
		}
		if err := pointsCordsDAO.CreatePointsRecord(pointsRecord); err != nil {
			utils.Logger.Error(user.ID + " create points record error")
		}
	}
}
