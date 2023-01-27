package main

import (
	"fmt"
	"strings"
	"time"
)

func solution(str1 string, str2 string) int {

	if strings.Contains(str1, str2) {
		return 1
	}

	return 2
}

func report(str1 string, str2 string, desireResult int) {

	input := fmt.Sprintf("Input:[%v, %v]", str1, str2)

	start := time.Now()

	result := solution(str1, str2)

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
	report("ab6CDE443fgh22iJKlmn1o", "6CD", 1)
	report("ppprrrogrammers", "pppp", 2)
	report("AbcAbcA", "AAA", 2)
}
