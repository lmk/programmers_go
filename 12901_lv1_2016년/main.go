package main

import (
	"fmt"
	"time"
)

func solution(a int, b int) string {
	weeks := []string{"SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT"}
	days := []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	result := 0
	for i := 1; i < a; i++ {
		result += days[i-1]
	}

	return weeks[(result+b+4)%len(weeks)]
}

func report(a int, b int, desireResult string) {

	input := fmt.Sprintf("Input:[%v, %v]", a, b)

	start := time.Now()

	result := solution(a, b)

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
	report(5, 24, "TUE")
}
