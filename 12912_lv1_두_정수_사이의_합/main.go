package main

import (
	"fmt"
	"time"
)

func solution(n int) int {

	for i := 1; i < n; i++ {
		if (n % i) == 1 {
			return i
		}
	}

	return 0
}

func report(n int, desireResult int) {

	start := time.Now()

	result := solution(n)

	duration := time.Since(start)

	if result == desireResult {
		fmt.Printf("[O] ")
	} else {
		fmt.Printf("[X] ")
	}

	fmt.Printf("Input:[%v], result:[%v], duration:%v\n", n, result, duration)
}

func main() {

	report(10, 3)
	report(12, 11)
}
