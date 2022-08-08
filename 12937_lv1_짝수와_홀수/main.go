package main

import (
	"fmt"
	"time"
)

func solution(num int) string {
	s := []string{"Even", "Odd"}

	if num == 0 {
		return s[0]
	}

	if 0 > num {
		num *= -1
	}

	return s[num%2]
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

	report(3, "Odd")
	report(4, "Even")
}
