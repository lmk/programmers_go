package main

import (
	"fmt"
	"strings"
	"time"
)

var BABY = []string{"aya", "ye", "woo", "ma"}

// 조카가 발음할 수 있는지 체크
func isBabbling(w string) bool {

	for 0 < len(w) {
		isBreak := true
		for _, b := range BABY {
			if strings.HasPrefix(w, b) {
				// 연속 체크
				if strings.HasPrefix(w[len(b):], b) {
					return false
				}
				w = w[len(b):]
				isBreak = false
			}
		}

		if isBreak {
			break
		}
	}

	return 0 == len(w)
}

func solution(babbling []string) int {

	count := 0
	for _, b := range babbling {
		if isBabbling(b) {
			count++
		}
	}

	return count
}

func report(babbling []string, desireResult int) {

	input := fmt.Sprintf("Input:[%v]", babbling)

	start := time.Now()

	result := solution(babbling)

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
	report([]string{"aya", "yee", "u", "maa"}, 1)
	report([]string{"ayaye", "uuu", "yeye", "yemawoo", "ayaayaa"}, 2)
}
