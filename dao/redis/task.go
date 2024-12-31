package redis

import (
	"context"
	"encoding/json"
	"miner/model/info"
	"miner/utils"
)

type TaskRDB struct {
}

func NewTaskRDB() *TaskRDB {
	return &TaskRDB{}
}

// 添加任务
// list
// +-------+-------------+
// | key   | val         |
// +-------+-------------+
// | task  | <task_info> |
// +-------+-------------+
//
// +----------------+------------+
// | key            | val        |
// +----------------+------------+
// | task:<task_id> | <task_res> |
// +----------------+------------+
func (c *TaskRDB) RPush(ctx context.Context, task *info.Task) error {
	key := MakeKey(TaskField)
	taskJSON, err := json.Marshal(task)
	if err != nil {
		return err
	}
	return utils.RDB.RPush(ctx, key, string(taskJSON))
}

// 取出任务
func (c *TaskRDB) LPop(ctx context.Context) (*info.Task, error) {
	key := MakeKey(TaskField)
	taskStr, err := utils.RDB.LPop(ctx, key)
	if err != nil {
		return nil, err
	}
	var task info.Task
	err = json.Unmarshal([]byte(taskStr), &task)
	return &task, err
}

// 数量
func (c *TaskRDB) LLen(ctx context.Context) (int64, error) {
	key := MakeKey(TaskField)
	return utils.RDB.LLen(ctx, key)
}
