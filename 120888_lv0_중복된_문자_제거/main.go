package main

import (
	"fmt"
	"time"
)

func solution(my_string string) string {

	m := map[rune]int{}
	answer := ""

	for _, r := range my_string {
		_, ok := m[r]
		if !ok {
			m[r] = 1
			answer += string(r)
		}
	}

	return answer
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
	report("people", "peol")
	report("We are the world", "We arthwold")
}
