package main

import (
	"fmt"
	"time"
)

func solution(n int) int {
	sum := n
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			sum += i
		}
	}

	return sum
}

func report(n int, desireResult int) {

	input := fmt.Sprintf("Input:[%v]", n)

	start := time.Now()

	result := solution(n)

	duration := time.Since(start)

	ox := ""
	if result == desireResult {
		ox = fmt.Sprintf("[O]")
	} else {
		ox = fmt.Sprintf("[X]")
	}

	fmt.Printf("%s %s result:[%v] duration:%v\n", ox, input, result, duration)
}

func main() {
	report(12, 28)
	report(5, 6)
}
