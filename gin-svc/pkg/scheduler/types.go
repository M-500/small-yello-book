package scheduler

// 调度器的设计

// TaskType 表示任务类型
type TaskType int

const (
	Python   TaskType = iota // Python任务
	Nextflow                 // Nextflow任务
	K8sPod                   // K8sPod任务
)

// Task 表示一个任务
type Task struct {
	ID       string
	Type     TaskType
	Priority int
	Data     interface{}
}

// Executor 接口定义了执行器的行为
type Executor interface {
	Execute(*Task) error
	Kill(*Task) error
}
