package main

import (
	"fmt"
	"time"
)

func solution(slice int, n int) int {
	if n%slice == 0 {
		return n / slice
	}
	return n/slice + 1
}

func report(slice int, num int, desireResult int) {

	input := fmt.Sprintf("Input:[%v, %v]", slice, num)

	start := time.Now()

	result := solution(slice, num)

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
	report(7, 10, 2)
	report(4, 12, 3)
}
