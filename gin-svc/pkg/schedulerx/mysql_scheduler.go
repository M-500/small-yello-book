package schedulerx

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// MySQL实现的调度器

type SchedulerDao interface {
	Insert(ctx context.Context, task TaskInterface) error
	Delete(ctx context.Context, task TaskInterface) error
	SafeGetOneRandom(ctx context.Context) (TaskInterface, error) // 随机从DB中捞取一个任务
}

type MySQLScheduler struct {
	dao SchedulerDao
	//tasks        []TaskInterface                // 任务队列
	mu           sync.Mutex                     // 互斥锁 用于控制并发冲突
	executors    map[TaskType]ExecutorInterface // 存放任务类型和执行器关系的映射
	stopCh       chan struct{}
	maxWorkerCnt int // 最大并发的worker数量
}

func NewMySQLScheduler(maxCnt int, dao SchedulerDao) SchedulerInterface {
	return &MySQLScheduler{
		executors: make(map[TaskType]ExecutorInterface),
		stopCh:    make(chan struct{}),
		//tasks:        make([]TaskInterface, 0),
		dao:          dao,
		maxWorkerCnt: maxCnt,
	}
}

func (m *MySQLScheduler) Run() {
	// 死循环，每隔2s中去数据库捞一批任务，来执行
	for i := 0; i < m.maxWorkerCnt; i++ {
		go func() {
			defer recover()
			m.worker()
		}()
	}
}

func (m *MySQLScheduler) Stop() {
	//TODO implement me
	panic("implement me")
}

func (m *MySQLScheduler) RegisterExecutor(taskType TaskType, executor ExecutorInterface) {
	m.executors[taskType] = executor
}

func (m *MySQLScheduler) SubmitTask(task TaskInterface) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	//m.tasks = append(m.tasks, task)
	// 需要持久化到DB
	return m.dao.Insert(context.Background(), task)
}

func (m *MySQLScheduler) UnRegisterExecutor(task TaskInterface) error {
	//TODO  去DB中删除这个任务
	return m.dao.Delete(context.Background(), task)
}

func (m *MySQLScheduler) worker() {
	for {
		select {
		case <-m.stopCh:
			return
		default:
			m.mu.Lock()
			task, err := m.dao.SafeGetOneRandom(context.Background()) // 随机从DB中捞一个健康的task任务
			if err != nil {
				m.mu.Unlock()
				time.Sleep(100 * time.Millisecond) // 执行休眠
			}
			m.mu.Unlock()
			m.executeTask(task)
		}
	}
}
func (m *MySQLScheduler) executeTask(task TaskInterface) {
	executor, ok := m.executors[task.GetTaskType()]
	if !ok {
		fmt.Printf("没有执行器可以运行这个task，task的类型为： %v\n", task.GetTaskType())
		return
	}
	err := executor.Execute(task)
	if err != nil {
		fmt.Printf("Error executing task %v\n", task)
	} else {
		fmt.Printf("Task %v executed successfully\n", task)
	}
}
