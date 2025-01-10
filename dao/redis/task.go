package redis

import (
	"context"
	"encoding/json"
	"miner/dao/mysql"
	"miner/model"
	"miner/model/info"
	"miner/utils"
)

type TaskRDB struct {
	taskDAO *mysql.TaskDAO
}

func NewTaskRDB() *TaskRDB {
	return &TaskRDB{
		taskDAO: mysql.NewTaskDAO(),
	}
}

// 添加任务
// hash
// +---------------------+--------------+
// | key                 | val          |
// +---------------------+--------------+
// | task_info:<task_id> | <task_info>  |
// +---------------------+--------------+
func (c *TaskRDB) Set(ctx context.Context, taskID string, task *info.Task) error {
	key := MakeKey(TaskInfoField, taskID)
	taskBytes, err := json.Marshal(task)
	if err != nil {
		return err
	}
	return utils.RDB.Set(ctx, key, string(taskBytes))
}

// 获取任务信息，其中就包括任务的结果
func (c *TaskRDB) Get(ctx context.Context, taskID string) (*info.Task, error) {
	key := MakeKey(TaskInfoField, taskID)
	taskJSON, err := utils.RDB.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	var task info.Task
	if err := json.Unmarshal([]byte(taskJSON), &task); err != nil {
		return nil, err
	}
	return &task, nil
}

// list
// +------------------+------------+
// | key              | val        |
// +------------------+------------+
// | task_id:<rig_id> | <task_id>  |
// +------------------+------------+
func (c *TaskRDB) RPush(ctx context.Context, rigID string, taskID string) error {
	key := MakeKey(TaskIDField, rigID)
	return utils.RDB.RPush(ctx, key, taskID)
}

// 取出队头任务 ID
func (c *TaskRDB) LPop(ctx context.Context, rigID string) (string, error) {
	key := MakeKey(TaskIDField, rigID)
	return utils.RDB.LPop(ctx, key)
}

// 数量
func (c *TaskRDB) LLen(ctx context.Context, rigID string) (int64, error) {
	key := MakeKey(TaskIDField)
	return utils.RDB.LLen(ctx, key)
}

// 添加任务
// list
// +------------------+------------+
// | key              | val        |
// +------------------+------------+
// | task_id:<rig_id> | <task_id>  |
// +------------------+------------+
// hash
// +---------------------+--------------+
// | key                 | val          |
// +---------------------+--------------+
// | task_info:<task_id> | <task_info>  |
// +---------------------+--------------+
// func (c *TaskRDB) AddTask(ctx context.Context, rigID string, taskID string, task *info.Task) error {
// 	list_key := MakeKey(TaskIDField, rigID)
// 	hash_key := MakeKey(TaskInfoField, taskID)

// 	pipe := utils.RDB.Client.Pipeline()
// 	// list
// 	pipe.RPush(ctx, list_key, taskID)
// 	taskJSON, err := json.Marshal(task)
// 	if err != nil {
// 		return err
// 	}
// 	// hash
// 	pipe.Set(ctx, hash_key, string(taskJSON), 0)
// 	_, err = pipe.Exec(ctx)

// 	return err
// }

// 获取矿机的队头任务
func (c *TaskRDB) GetTask(ctx context.Context, rigID string) (*model.Task, error) {
	// redis list 中弹出队头
	taskID, err := c.LPop(ctx, rigID)
	if err != nil {
		return nil, err
	}
	// 数据库中找到对应 id 的任务信息
	var task model.Task
	if err := utils.DB.Where("id = ?", taskID).First(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}
