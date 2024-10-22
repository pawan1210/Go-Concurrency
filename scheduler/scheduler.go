package concurrency

import (
	"fmt"
)


func SchedulingExample() {
    for i:=0; i<1000; i++ {
		go func() {
			fmt.Println("go coroutine")
		}()
		fmt.Println("main thread")
	}
}
// GOMAXPROCS=1 go run main.go