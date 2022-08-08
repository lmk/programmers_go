package main

import (
	"fmt"
	"time"
)

func solution(s string) string {
	result := ""
	if len(s)%2 == 0 {
		i := len(s) / 2
		result = s[i-1 : i+1]
	} else {
		i := int(len(s) / 2)
		result = s[i : i+1]
	}
	return result
}

func report(s string, desireResult string) {

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

	report("abcde", "c")
	report("qwer", "we")
}
