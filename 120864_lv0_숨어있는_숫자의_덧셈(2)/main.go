package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func solution(my_string string) int {

	rep_string := ""
	for _, r := range my_string {
		if ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') {
			rep_string += " "
		} else {
			rep_string += string(r)
		}
	}

	arr := strings.Split(rep_string, " ")
	sum := 0
	for _, a := range arr {
		n, _ := strconv.Atoi(a)
		sum += n
	}

	return sum
}

func report(my_string string, desireResult int) {

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
	report("aAb1B2cC34oOp", 37)
	report("1a2b3c4d123Z", 133)
}
