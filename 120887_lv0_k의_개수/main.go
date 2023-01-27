package main

import (
	"fmt"
	"time"
)

func solution(i int, j int, k int) int {

	c := 0
	for ; i <= j; i++ {
		for n := i; n > 0; n /= 10 {
			if n%10 == k {
				c++
			}
		}
	}

	return c
}

func report(i int, j int, k int, desireResult int) {

	input := fmt.Sprintf("Input:[%v, %v, %v]", i, j, k)

	start := time.Now()

	result := solution(i, j, k)

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
	report(1, 13, 1, 6)
	report(10, 50, 5, 5)
	report(3, 10, 2, 0)
}
