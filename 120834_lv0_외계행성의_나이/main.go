package main

import (
	"fmt"
	"time"
)

func solution(age int) string {

	s := ""
	for age/10 > 0 {
		s = string(age-(age/10*10)+'a') + s
		age /= 10
	}

	return string(age+'a') + s
}

func report(age int, desireResult string) {

	input := fmt.Sprintf("Input:[%v]", age)

	start := time.Now()

	result := solution(age)

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
	report(23, "cd")
	report(51, "fb")
	report(100, "baa")
}
