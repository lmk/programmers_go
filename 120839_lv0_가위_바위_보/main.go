package main

import (
	"fmt"
	"time"
)

func solution(rsp string) string {

	answer := ""
	for _, r := range rsp {
		switch r {
		case '0':
			answer += "5"
		case '2':
			answer += "0"
		case '5':
			answer += "2"
		}
	}
	return answer
}

func report(rsp string, desireResult string) {

	input := fmt.Sprintf("Input:[%v]", rsp)

	start := time.Now()

	result := solution(rsp)

	duration := time.Since(start)

	isSame := true
	if result != desireResult {
		isSame = false
	}

	if isSame {
		fmt.Printf("[O] %s result:[%v] duration:%v\n", input, result, duration)
	} else {
		fmt.Printf("[X] %s result:[%v]->[%v] duration:%v\n", input, result, desireResult, duration)
	}
}

func main() {
	report("2", "0")
	report("205", "052")
}
