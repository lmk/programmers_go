package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func solution(s string) int {

	nums := []string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	for i, num := range nums {
		s = strings.Replace(s, num, string(i+'0'), -1)
	}

	result, _ := strconv.Atoi(s)

	return result
}

func report(s string, desireResult int) {

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

	report("one4seveneight", 1478)
	report("23four5six7", 234567)
	report("2three45sixseven", 234567)
	report("123", 123)
}
