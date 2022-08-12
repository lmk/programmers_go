package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func solution(n int64) []int {
	s := fmt.Sprintf("%v", n)
	as := strings.Split(s, "")

	result := []int{}
	last := len(as) - 1
	for i, _ := range as {
		m, _ := strconv.Atoi(as[last-i])
		result = append(result, m)
	}

	return result
}

func report(n int64, desireResult []int) {

	input := fmt.Sprintf("Input:[%v]", n)

	start := time.Now()

	result := solution(n)

	duration := time.Since(start)

	isSame := true
	if len(result) != len(desireResult) {
		isSame = false
	} else {
		for i, v := range result {
			if desireResult[i] != v {
				isSame = false
				break
			}
		}
	}

	ox := ""
	if isSame {
		ox = fmt.Sprintf("[O]")
	} else {
		ox = fmt.Sprintf("[X]")
	}

	fmt.Printf("%s %s result:[%v] duration:%v\n", ox, input, result, duration)
}

func main() {
	report(12345, []int{5, 4, 3, 2, 1})
}
