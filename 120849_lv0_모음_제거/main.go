package main

import (
	"fmt"
	"time"
)

func solution(my_string string) string {
	s := ""
	for _, r := range my_string {
		if r != 'a' && r != 'e' && r != 'i' && r != 'o' && r != 'u' {
			s += string(r)
		}
	}

	return s
}

func report(my_string string, desireResult string) {

	input := fmt.Sprintf("Input:[%v]", my_string)

	start := time.Now()

	result := solution(my_string)

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
	report("bus", "bs")
	report("nice to meet you", "nc t mt y")
}
