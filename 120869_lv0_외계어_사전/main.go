package main

import (
	"fmt"
	"strings"
	"time"
)

func solution(spell []string, dic []string) int {

	for _, d := range dic {
		ok := false

		for _, s := range spell {
			ok = true
			if !strings.Contains(d, s) {
				ok = false
				break
			}
		}

		if ok {
			return 1
		}
	}

	return 2
}

func report(spell []string, dic []string, desireResult int) {

	input := fmt.Sprintf("Input:[%v, %v]", spell, dic)

	start := time.Now()

	result := solution(spell, dic)

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
	report([]string{"p", "o", "s"}, []string{"sod", "eocd", "qixm", "adio", "soo"}, 2)
	report([]string{"z", "d", "x"}, []string{"def", "dww", "dzx", "loveaw"}, 1)
	report([]string{"s", "o", "m", "d"}, []string{"moos", "dzx", "smm", "sunmmo", "som"}, 2)
}
