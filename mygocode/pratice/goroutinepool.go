package main

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 定义任务
type Task struct {
	Handler func(v ...interface{}) //可变参数 ，处理可变参数
	Params  []interface{}          // 定义空接口
}

// 任务池的定义
type Pool struct {
	capacity       uint64
	runningWorkers uint64
	status         int64
	chTask         chan *Task
	sync.Mutex
	PanicHandler func(interface{})
}

// 任务池的构造函数
var ErrInvalidPoolCap = errors.New("invalid pool cap")

const (
	RUNNING = 1
	STOPED  = 0
)

func NewPool(capacity uint64) (*Pool, error) {
	if capacity <= 0 {
		return nil, ErrInvalidPoolCap
	}

	return &Pool{
		capacity: capacity,
		status:   RUNNING,
		// 初始化任务队列
		chTask: make(chan *Task, capacity),
	}, nil

}

func (p *Pool) run() {
	p.runningWorkers++

	go func() {
		defer func() {
			p.runningWorkers--
		}()

		for {
			select { //阻塞等待任务，结束信号到来
			case task, ok := <-p.chTask:
				if !ok { //如果 channel 被关闭， 结束worker 运行
					return
				}
				// 执行任务
				task.Handler(task.Params...)
			}
		}
	}()

}

func (p *Pool) incRunning() {
	atomic.AddUint64(&p.runningWorkers, 1)
}

func (p *Pool) decRunning() {
	atomic.AddUint64(&p.runningWorkers, ^uint64(0))
}

func (p *Pool) GetRunningWorkers() uint64 {
	return atomic.LoadUint64(&p.runningWorkers)
}

func (p *Pool) GetCap() uint64 {
	return p.capacity
}

func (p *Pool) setStatus(status int64) bool {
	p.Lock()
	defer p.Unlock()

	if p.status == status {
		return false
	}

	p.status = status
	return true
}

// 任务池关闭
var ErrPoolAlreadyClosed = errors.New("pool already closed")

func (p *Pool) Close() {
	p.setStatus(STOPED)

	for len(p.chTask) > 0 {
		time.Sleep(1e6)
	}

	close(p.chTask)
}

// 生产任务
func (p *Pool) Put(task *Task) error {
	// 枷锁防止启动多个worker
	p.Lock()
	defer p.Unlock()

	// 执行任务
	if p.GetRunningWorkers() < p.GetCap() { //如果任务池满，则不创建worker
		p.run()
	}

	//将任务推入队列，等待消费
	if p.status == RUNNING {
		p.chTask <- task
	}

	return nil
}

func main() {
	//创建任务池
	pool, err := NewPool(4)

	if err != nil {
		panic(err)
	}

	for i := 0; i < 1024; i++ {
		//任务放入池子
		pool.Put(&Task{
			Handler: func(v ...interface{}) {
				fmt.Println(v)
			},
			Params: []interface{}{i},
		})
	}
	time.Sleep(1e9)
}
