package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func NormalChannelOperation() {

	channel := make(chan string)
	var wg sync.WaitGroup

	go func (ch chan string) {
		ch <- "sample"

		defer func() {
			wg.Done()
		}()
	}(channel)
	
	wg.Add(1)
	go func(ch chan string) {
		var value = <-ch
		fmt.Print(value)

		defer func() {
			wg.Done()
		}()
	}(channel)

	wg.Wait()
}

func DeadLockExample() {
	var wg sync.WaitGroup
	channel := make(chan string)

	wg.Add(1)
	
	go func(ch chan string) {
		ch <- "sample"
	}(channel)
	
	wg.Wait()
}

func worker(tasks chan string) {
	for task := range tasks {
		time.Sleep(5000)
		fmt.Println(task)	
	}
}

func orderedWorker(tasks chan string, mu *sync.Mutex) {
	/*
		why do we need to use mutex ?
	*/
	mu.Lock()
	fmt.Println(<-tasks)
	mu.Unlock()
}

func MockQueue() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var orderedTasksChannel = make(chan string, 1)
	var unorderedTasksChannel = make(chan string, 5)

	orderedTasksList := []int{1, 2, 3, 4, 5}
	unorderedTasksList := []int{6, 7, 8, 9, 10}

	for i:=0; i<len(orderedTasksList); i++ {	
		orderedTasksChannel <- fmt.Sprintf("ordered task %v", orderedTasksList[i])
		
		wg.Add(1)

		go func() {
			orderedWorker(orderedTasksChannel, &mu)
			wg.Done()
		}()
	}

	for _, unorderedTask := range unorderedTasksList {
		unorderedTasksChannel <- fmt.Sprintf("unordered task %v", unorderedTask)

		wg.Add(1)

		go func() {
			worker(unorderedTasksChannel)
			wg.Done()
		}()
	}

	/*
		closure of the channel means that no
		more values will be sent to the channel
		and channel can process the values inside it.
	*/
	close(unorderedTasksChannel)
	wg.Wait()
}