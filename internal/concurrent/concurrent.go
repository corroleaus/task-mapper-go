package concurrent

import (
	"fmt"
	"time"
)

// WorkerPool - stuff
type WorkerPool struct {
	workermax int
	workers   []Worker
	channel   chan func()
}

func (pool *WorkerPool) Submit(task func()) {
	pool.channel <- task
}

//Worker - stuff
type Worker struct {
	channel chan func()
}

// NewPool - stuff
func NewPool(workermax int) *WorkerPool {
	workers := make([]Worker, workermax)
	for i := range workers {
		workers[i].channel = make(chan func(), 1)
		go workers[i].WorkerExecute()
	}

	// keep for ref
	// for i := 0; i < workermax; i++ {
	// 	workers[i].channel = make(chan func())
	// }
	return &WorkerPool{
		workermax: workermax,
		workers:   workers,
		channel:   make(chan func(), 1),
	}
}

// Execute - stuff
func Execute(pool *WorkerPool) {
	fmt.Printf("Starting Pool Execute\n")
	go func() {
		workerIndex := 0
		for {
			f := <-pool.channel
			pool.workers[workerIndex%pool.workermax].channel <- f
			workerIndex++
		}
	}()
}

func (worker *Worker) WorkerExecute() {
	fmt.Printf("Starting Worker Execute\n")
	for {
		f := <-worker.channel
		fmt.Printf("executing %s\n", time.Now().Format(time.RFC3339))
		f()
	}
}

func TimeConsuming() {
	time.Sleep(1 * time.Second)
}
