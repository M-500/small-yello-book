package schedulerx

// TaskStatus 任务的状态
type TaskStatus int

const (
	TaskStatusPending TaskStatus = iota // 任务等待执行
	TaskStatusRunning                   // 任务正在执行
	TaskStatusFailed                    // 任务执行失败
	TaskStatusSuccess                   // 任务执行成功
	TaskStatusKilled                    // 任务被终止
	TaskStatusUnknown                   // 未知状态
)

// TaskType 表示任务类型
type TaskType int

const (
	Python   TaskType = iota // Python任务
	Nextflow                 // Nextflow任务
	K8sPod                   // K8sPod任务
)

// SchedulerInterface
// @Description: 调度器的抽象
type SchedulerInterface interface {
	Run()                                                           // 启动调度器 初始化一些东西
	Stop()                                                          // 停止调度器 清理一些调度器相关的资源
	RegisterExecutor(taskType TaskType, executor ExecutorInterface) // 注册执行器  一个调度器应该支持多种执行器，比如Python Nextflow K8S的Pod等
	SubmitTask(task TaskInterface) error                            // 提交任务 任务的类型应该和执行器的类型对应
	UnRegisterExecutor(executor TaskInterface) error                // 注销一个任务
}

// ExecutorInterface 执行器的抽象
type ExecutorInterface interface {
	Execute(task TaskInterface) error // 执行任务
	Kill(task TaskInterface) error    // 终止任务
}

// TaskInterface 任务的抽象
type TaskInterface interface {
	SaveTask(task TaskInterface) error                      // 持久化Task信息
	DeleteTask(task TaskInterface) error                    // 删除Task信息
	UpdateTask(task TaskInterface) error                    // 更新Taskz状态信息
	QueryTaskStatus(task TaskInterface) (TaskStatus, error) // 查询Task的状态
	GetTaskType() TaskType
}
