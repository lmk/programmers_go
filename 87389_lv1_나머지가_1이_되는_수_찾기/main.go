package main

import (
	"fmt"
	"time"
)

func solution(a int, b int) int64 {
	var sum int64

	if a > b {
		a, b = b, a
	}

	for i := a; i <= b; i++ {
		sum += int64(i)
	}
	return sum
}

func report(a int, b int, desireResult int64) {

	start := time.Now()

	result := solution(a, b)

	duration := time.Since(start)

	if result == desireResult {
		fmt.Printf("[O] ")
	} else {
		fmt.Printf("[X] ")
	}

	fmt.Printf("Input:[%v, %v], result:[%v], duration:%v\n", a, b, result, duration)
}

func main() {

	report(3, 5, 12)
	report(3, 3, 3)
	report(5, 3, 12)
}
