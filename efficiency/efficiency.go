package efficiency

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func findSumInRange(nums []int, left int, right int, result chan int) {
	var sum int = 0
	for i:=left; i<=right; i++ {
		sum += nums[i]
	}

	result <- sum
}

func FindSum() {
	startTime := time.Now()
	var wg sync.WaitGroup
    nums := []int{}
	partition := 100
	left := 0
	result := make(chan int)


	for i:=0; i<10000; i++ {
		nums = append(nums, i+1)
	}

	sumChannel := make(chan int, len(nums)/partition)

    for {
		if left==len(nums){
			break
		}

		right := left + partition - 1
		right = int(math.Min(float64(right), float64(len(nums)-1)))
		wg.Add(1)

		go func(left int, right int) {
			fmt.Println("go coroutine spawned")
			findSumInRange(nums, left, right, sumChannel)
			wg.Done()
		}(left, right)
		
		left = right+1
	}


	wg.Wait()

	go func() {
		var total = 0
		for sum := range(sumChannel) {
			total += sum
		}

		result <- total
	}()

	close(sumChannel)
	
    fmt.Println("Total sum:", <-result)
	elapsedTime := time.Since(startTime)
	fmt.Println("Time taken:", elapsedTime)
}