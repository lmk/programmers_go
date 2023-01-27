package main

import (
	"fmt"
	"time"
	"unicode"
)

func solution(my_string string) string {

	answer := []rune(my_string)

	for i := range answer {
		if 'a' <= answer[i] && 'z' >= answer[i] {
			answer[i] = unicode.ToUpper(answer[i])
		} else {
			answer[i] = unicode.ToLower(answer[i])
		}
	}

	return string(answer)
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
	report("cccCCC", "CCCccc")
	report("abCdEfghIJ", "ABcDeFGHij")
}
