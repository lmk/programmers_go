package main

import (
	"fmt"
	"time"
)

func solution(s string) bool {
	if len(s) != 4 && len(s) != 6 {
		return false
	}

	i := 0
	for i < len(s) {
		if "0"[0] > s[i] || s[i] > "9"[0] {

			return false
		}
		i++
	}

	return true
}

func report(s string, desireResult bool) {

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

	report("a234", false)
	report("1234", true)
}
