package main

import (
	"fmt"
	"time"
)

func solution(n int) string {
	s := []string{"수", "박"}
	result := ""

	for i := 0; i < n; i++ {
		result += s[i%2]
	}

	return result
}

func report(num int, desireResult string) {

	start := time.Now()

	result := solution(num)

	duration := time.Since(start)

	if result == desireResult {
		fmt.Printf("[O] ")
	} else {
		fmt.Printf("[X] ")
	}

	fmt.Printf("Input:[%v], result:[%v], duration:%v\n", num, result, duration)
}

func main() {

	report(3, "수박수")
	report(4, "수박수박")
}
