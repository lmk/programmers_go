package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func countPrime(n int) int {
	count := 1 // 1

	if n <= 1 {
		return count
	}

	for i := 2; i*2 <= n; i++ {
		if n%i == 0 {
			count++
		}
	}

	return count + 1 // self
}

func solution(number int, limit int, power int) int {

	var sum int32 = 0

	var wg sync.WaitGroup

	for i := 1; i <= number; i++ {
		wg.Add(1)
		go func(i int) {
			p := countPrime(i)
			if p > limit {
				p = power
			}
			atomic.AddInt32(&sum, int32(p))
			defer wg.Done()
		}(i)
	}

	wg.Wait()
	return int(sum)
}

func report(number int, limit int, power int, desireResult int) {

	input := fmt.Sprintf("Input:[%v, %v, %v]", number, limit, power)

	start := time.Now()

	result := solution(number, limit, power)

	duration := time.Since(start)

	isSame := true
	if result != desireResult {
		isSame = false
	}

	if isSame {
		fmt.Printf("[O] %s result:[%v] duration:%v\n", input, result, duration)
	} else {
		fmt.Printf("[X] %s result:[%v]->[%v] duration:%v\n", input, result, desireResult, duration)
	}
}

func main() {
	report(5, 3, 2, 10)
	report(10, 3, 2, 21)
}
