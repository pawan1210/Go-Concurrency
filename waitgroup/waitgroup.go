package waitgroup

import (
	"fmt"
	"sync"
	"time"
)

func SkipGoCoroutines() {
	go func() {
		fmt.Print("Go coroutine 1")
	}()

	go func() {
		fmt.Print("Go coroutine 2")
	}()
}

func WaitForGoCoroutines() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Go coroutine 1")
		defer wg.Done()
	}()

	go func() {
		fmt.Println("Go coroutine 2")
		defer wg.Done()
	}()

	time.Sleep(1000000)
}