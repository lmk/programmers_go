package main

import (
	"fmt"
	"math"
	"time"
)

func solution(n int64) int64 {
	q := math.Sqrt(float64(n))
	if float64(int64(q)) == q {
		return int64(math.Pow(q+1, 2))
	}

	return -1
}

func report(n int64, desireResult int64) {

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

	report(121, 144)
	report(3, -1)
}
