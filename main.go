package main

import (
	// "fmt"
	"go-tut/waitgroup"
	"go-tut/concurrency"
	customScheduler "go-tut/scheduler"

	"go-tut/parallelism"
	// "go-tut/efficiency"
)

func main() {
	waitgroup.SkipGoCoroutines()
	// waitgroup.WaitForGoCoroutines()
	customScheduler.SchedulingExample()
	parallelism.MockQueue()
	// concurrency.SchedulingExample()
	// concurrency.NormalChannelOperation()
	// concurrency.DeadLockExample()
	concurrency.MockQueue()
	// efficiency.FindSum()
}