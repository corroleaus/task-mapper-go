package main

import (
	"fmt"
	"github.com/corroleaus/task-mapper-go/internal"
)


func main() {
	x := Concurrent.ThreadPool{Threadsmax:9}
	fmt.Printf("%d\n", x.Threadsmax)
}