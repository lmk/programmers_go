package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func solution(numbers string) int64 {

	strs := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for i, s := range strs {
		numbers = strings.ReplaceAll(numbers, s, string('0'+i))
	}

	i, _ := strconv.ParseInt(numbers, 10, 64)

	return i
}

func report(numbers string, desireResult int64) {

	input := fmt.Sprintf("Input:[%v]", numbers)

	start := time.Now()

	result := solution(numbers)

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
	report("onetwothreefourfivesixseveneightnine", 123456789)
	report("onefourzerosixseven", 14067)
}
