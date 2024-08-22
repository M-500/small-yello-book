package scheduler

import (
	"container/heap"
	"fmt"
	"sync"
	"time"
)

// Scheduler 表示调度器
type Scheduler struct {
	tasks       PriorityQueue
	mu          sync.Mutex
	executors   map[TaskType]Executor
	stopCh      chan struct{}
	workerCount int
}

// NewScheduler 创建一个新的调度器
func NewScheduler(workerCount int) *Scheduler {
	return &Scheduler{
		tasks:       make(PriorityQueue, 0),
		executors:   make(map[TaskType]Executor),
		stopCh:      make(chan struct{}),
		workerCount: workerCount,
	}
}

// RegisterExecutor 注册一个执行器
func (s *Scheduler) RegisterExecutor(taskType TaskType, executor Executor) {
	s.executors[taskType] = executor
}

// SubmitTask 提交一个任务
func (s *Scheduler) SubmitTask(task *Task) {
	s.mu.Lock()
	defer s.mu.Unlock()
	heap.Push(&s.tasks, task)
}

// Start 启动调度器
func (s *Scheduler) Start() {
	for i := 0; i < s.workerCount; i++ {
		go s.worker()
	}
}

// Stop 停止调度器
func (s *Scheduler) Stop() {
	close(s.stopCh)
}

// worker 是调度器的工作协程
func (s *Scheduler) worker() {
	for {
		select {
		case <-s.stopCh:
			return
		default:
			s.mu.Lock()
			if s.tasks.Len() > 0 {
				task := heap.Pop(&s.tasks).(*Task)
				s.mu.Unlock()
				s.executeTask(task)
			} else {
				s.mu.Unlock()
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
}

// executeTask 执行一个任务
func (s *Scheduler) executeTask(task *Task) {
	executor, ok := s.executors[task.Type]
	if !ok {
		fmt.Printf("No executor found for task type %v\n", task.Type)
		return
	}

	err := executor.Execute(task)
	if err != nil {
		fmt.Printf("Error executing task %s: %v\n", task.ID, err)
	} else {
		fmt.Printf("Task %s executed successfully\n", task.ID)
	}
}
