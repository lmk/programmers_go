package main

import (
	"fmt"
	"time"
)

func solution(num int, k int) int {
	snum := fmt.Sprintf("%d", num)
	sk := fmt.Sprintf("%d", k)

	for i, r := range snum {
		if string(r) == sk {
			return i + 1
		}
	}
	return -1
}

func report(num int, k int, desireResult int) {

	input := fmt.Sprintf("Input:[%v, %v]", num, k)

	start := time.Now()

	result := solution(num, k)

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
	report(29183, 1, 3)
	report(232443, 4, 4)
	report(123456, 7, -1)
}
