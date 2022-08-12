package main

import (
	"fmt"
	"time"
)

func solution(n int) int {
	sum := 0
	m := n
	for m > 0 {
		sum += m % 10
		m /= 10
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

	report(123, 6)
	report(987, 24)

}
