package main

import (
	"github.com/corroleaus/task-mapper-go/internal/concurrent"
)

func main() {
	pool := concurrent.NewPool(10)
	concurrent.Execute(pool)
	for i := 0; i < 100; i++ {
		pool.Submit(concurrent.TimeConsuming)
	}

	// time.Sleep(12 * time.Second)
}
