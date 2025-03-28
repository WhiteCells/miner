package mysql

import (
	"context"
	"fmt"
	"miner/model"
	"miner/utils"
	"strconv"
)

type TaskDAO struct {
	// taskRDB *redis.TaskRDB
}

func NewTaskDAO() *TaskDAO {
	return &TaskDAO{
		// taskRDB: redis.NewTaskRDB(),
	}
}

// 添加任务
func (TaskDAO) AddTask(ctx context.Context, userID, farmID, minerID int, task *model.Task) error {
	key := fmt.Sprintf("%s:%d", "task_id", minerID)
	if err := utils.DB.WithContext(ctx).Create(task).Error; err != nil {
		return err
	}
	taskIDStr := strconv.Itoa(task.ID)
	return utils.RDB.RPush(ctx, key, taskIDStr)
}

// 获取任务
func (dao *TaskDAO) GetTask(ctx context.Context, taskID string) (*model.Task, error) {
	var task model.Task
	if err := utils.DB.WithContext(ctx).Where("id = ?", taskID).First(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

// 更新任务
func (dao *TaskDAO) UpdateTask(ctx context.Context, taskID int, updateInfo map[string]any) error {
	return utils.DB.WithContext(ctx).
		Model(&model.Task{}).
		Where("id=?", taskID).
		Updates(updateInfo).Error
}
