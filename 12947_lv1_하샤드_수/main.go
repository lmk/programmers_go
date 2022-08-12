package main

import (
	"fmt"
	"time"
)

func solution(x int) bool {
	sum := 0
	m := x
	for m > 0 {
		sum += m % 10
		m /= 10
	}

	return x%sum == 0
}

func report(x int, desireResult bool) {

	input := fmt.Sprintf("Input:[%v]", x)

	start := time.Now()

	result := solution(x)

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

	report(1, true)
	report(5, true)
	report(10, true)
	report(12, true)
	report(11, false)
	report(13, false)
	report(5500, true)
	report(1234, false)
	report(10000, true)
}
