/**
 * @Author: lj
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/03/18 10:46
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	pool := New(4)

	for i := 0; i < 8; i++ {
		pool.NewTask(func() {
			time.Sleep(1 * time.Millisecond)
			fmt.Println(time.Now())
		})
	}

	// 保证所有的协程都执行完毕
	time.Sleep(5 * time.Second)
}

type Pool struct {
	work chan func()
	sem  chan struct{}
}

func New(size int) *Pool {
	return &Pool{
		work: make(chan func()),
		sem:  make(chan struct{}, size),
	}
}

func (p *Pool) NewTask(task func()) {
	select {
	case p.work <- task:
	case p.sem <- struct{}{}:
		go p.worker(task)
	}
}

func (p *Pool) worker(task func()) {
	defer func() {
		<-p.sem
	}()
	for {
		task()
		<-p.work
	}
}

/*func main() {
	wp := New(3, 50).Start()
	lenth := 100
	for i := 0; i < lenth; i++ {
		wp.PushTaskFunc(func(args ...interface{}) {
			fmt.Print(args[0].(int), " ")
		}, i)
	}
	wp.Stop()
}

type TaskFunc func(args ...interface{})

type Task struct {
	f    TaskFunc
	args []interface{}
}

type WorkPool struct {
	pool        chan *Task
	workerCount int

	stopCtx        context.Context
	stopCancelFunc context.CancelFunc
	wg             sync.WaitGroup
}

func (t *Task) Execute() {
	t.f(t.args...)
}

func New(workerCount, poolLen int) *WorkPool {
	return &WorkPool{
		pool:        make(chan *Task, poolLen),
		workerCount: workerCount,
	}
}

func (w *WorkPool) PushTask(t *Task) {
	w.pool <- t
}

func (w *WorkPool) PushTaskFunc(f TaskFunc, args ...interface{}) {
	w.pool <- &Task{f: f, args: args}
}

func (w *WorkPool) work() {
	for {
		select {
		case <-w.stopCtx.Done():
			w.wg.Done()
		case t := <-w.pool:
			t.Execute()
		}
	}
}

func (w *WorkPool) Start() *WorkPool {
	w.wg.Add(w.workerCount)
	w.stopCtx, w.stopCancelFunc = context.WithCancel(context.Background())
	for i := 0; i < w.workerCount; i++ {
		go w.work()
	}
	return w
}

func (w *WorkPool) Stop() {
	w.stopCancelFunc()
	w.wg.Wait()
}
*/
