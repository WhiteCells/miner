package mysql

import (
	"context"
	"fmt"
	"miner/model"
	"miner/utils"
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
func (dao *TaskDAO) AddTask(ctx context.Context, rigID string, task *model.Task) error {
	key := fmt.Sprintf("%s:%s", "task_id", rigID)
	if err := utils.DB.Create(task).Error; err != nil {
		return err
	}
	return utils.RDB.RPush(ctx, key, task.ID)
}

// 获取任务
func (dao *TaskDAO) GetTask(taskID string) (*model.Task, error) {
	var task model.Task
	if err := utils.DB.Where("id = ?", taskID).First(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

// 更新任务
func (dao *TaskDAO) UpdateTask(task *model.Task) error {
	return utils.DB.Save(task).Error
}
