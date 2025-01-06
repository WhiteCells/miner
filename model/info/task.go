package info

type TaskType string

var Cmd TaskType = "cmd"
var Config TaskType = "config"

type TaskStatus string

var Pending TaskStatus = "pending"
var Done TaskStatus = "done"

type Task struct {
	ID      string     `json:"id"`
	Type    TaskType   `json:"type"`
	Status  TaskStatus `json:"status"`
	Content string     `json:"content"`
	Result  string     `json:"result"`
}
