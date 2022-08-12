package main

import (
	"fmt"
	"time"
)

func solution(num int) int {

	if num == 1 {
		return 0
	}

	count := 0
	for num > 1 {
		if num%2 == 0 {
			num /= 2
		} else {
			num = num*3 + 1
		}

		if count > 500 {
			return -1
		}

		count++
	}

	return count
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

	report(6, 8)
	report(16, 4)
	report(626331, -1)
}
