package redis

import (
	"context"
	"encoding/json"
	"errors"
	"miner/dao/mysql"
	"miner/model"
	"miner/model/info"
	"miner/utils"
	"strconv"
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

// 获取矿机的队头任务
func (c *TaskRDB) GetTask(ctx context.Context, rigID int) (*model.Task, error) {
	rigIDStr := strconv.Itoa(rigID)
	// redis list 中弹出队头
	taskID, err := c.LPop(ctx, rigIDStr)
	if err != nil {
		return nil, errors.New("no task in redis list")
	}
	// 数据库中找到对应 id 的任务信息
	var task model.Task
	if err := utils.DB.Where("id = ?", taskID).First(&task).Error; err != nil {
		return nil, errors.New("no task in db")
	}
	return &task, nil
}
