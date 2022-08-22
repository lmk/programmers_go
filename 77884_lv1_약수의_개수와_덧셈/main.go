package main

import (
	"fmt"
	"time"
)

func solution(left int, right int) int {

	result := 0
	for n := left; n <= right; n++ {
		count := 1
		for i := n / 2; i > 0; i-- {
			if n%i == 0 {
				count++
			}
		}

		if count%2 == 0 {
			result += n
		} else {
			result -= n
		}
	}

	return result
}

func report(left int, right int, desireResult int) {

	start := time.Now()

	result := solution(left, right)

	duration := time.Since(start)

	if result == desireResult {
		fmt.Printf("[O] ")
	} else {
		fmt.Printf("[X] ")
	}

	fmt.Printf("Input:[%v, %v], result:[%v], duration:%v\n", left, right, result, duration)
}

func main() {

	report(13, 17, 43)
	report(24, 27, 52)
	report(1, 1000, 52)
}
