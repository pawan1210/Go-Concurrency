package parallelism

import (
	"fmt"
	"sync"
	"time"
)

func worker(tasks chan string) {
	for task := range tasks {
		// simulate long running task
		time.Sleep(5000)
		fmt.Println(task)
	}
}

func MockQueue() {
	var wg sync.WaitGroup
	tasks := make(chan string)

	for i:=0; i<10; i++ {
		tasks <- fmt.Sprintf("task %v", i)
	}

	for i:=0; i<5; i++ {
		wg.Add(1)
		go func() {
			worker(tasks)
			defer wg.Done()
		}()
	}

	wg.Wait()
}
