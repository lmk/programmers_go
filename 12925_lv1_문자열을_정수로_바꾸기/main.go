package main

import (
	"fmt"
	"time"
)

func solution(s string) int {

	dig := []int{10000, 1000, 100, 10, 1}
	i := 0
	offset := len(dig) - len(s)

	if s[i] == '+' || s[i] == '-' {
		i++
	}

	result := 0
	for i < len(s) {
		result += int(s[i]-'0') * dig[offset+i]
		i++
	}

	if s[0] == '-' {
		result *= -1
	}

	return result
}

func report(s string, desireResult int) {

	start := time.Now()

	result := solution(s)

	duration := time.Since(start)

	if result == desireResult {
		fmt.Printf("[O] ")
	} else {
		fmt.Printf("[X] ")
	}

	fmt.Printf("Input:[%v], result:[%v], duration:%v\n", s, result, duration)
}

func main() {

	report("-123", -123)
	report("1234", 1234)
	report("12345", 12345)
	report("-1234", -1234)
	report("+1234", 1234)
}
