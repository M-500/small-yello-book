package schedulerx

import "gorm.io/gorm"

type BaseTask struct {
	TaskUUID   string
	TaskName   string
	TaskType   TaskType
	TaskStatus TaskStatus
	TaskConfig string
	TaskResult string
}

type TaskShell struct {
	db *gorm.DB
	BaseTask
}

func NewTaskShell() TaskInterface {
	return &TaskShell{}
}

func (t *TaskShell) SaveTask(task TaskInterface) error {
	//TODO implement me
	panic("implement me")
}

func (t *TaskShell) DeleteTask(task TaskInterface) error {
	//TODO implement me
	panic("implement me")
}

func (t *TaskShell) UpdateTask(task TaskInterface) error {
	//TODO implement me
	panic("implement me")
}

func (t *TaskShell) QueryTaskStatus(task TaskInterface) (TaskStatus, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TaskShell) GetTaskType() TaskType {
	//TODO implement me
	panic("implement me")
}
