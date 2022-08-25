package main

import (
	"fmt"
	"time"
)

func solution(survey []string, choices []int) string {
	indicator := []string{"R", "T", "C", "F", "J", "M", "A", "N"}

	score := map[string]int{}

	for i, v := range survey {
		if choices[i] < 4 {
			score[string(v[0])] += 4 - choices[i]
		} else if choices[i] > 4 {
			score[string(v[1])] += choices[i] - 4
		}
	}

	result := ""
	for i := 0; i < len(indicator); i += 2 {
		if score[indicator[i]] < score[indicator[i+1]] {
			result += indicator[i+1]
		} else {
			result += indicator[i]
		}
	}

	return result
}

func report(survey []string, choices []int, desireResult string) {

	input := fmt.Sprintf("Input:[%v]", survey, choices)

	start := time.Now()

	result := solution(survey, choices)

	duration := time.Since(start)

	isSame := true
	if result != desireResult {
		isSame = false
	}

	if isSame == true {
		fmt.Printf("[O] %s result:[%v] duration:%v\n", input, result, duration)
	} else {
		fmt.Printf("[X] %s result:[%v]->[%v] duration:%v\n", input, result, desireResult, duration)
	}
}

func main() {

	report([]string{"AN", "CF", "MJ", "RT", "NA"}, []int{5, 3, 2, 7, 5}, "TCMA")
	report([]string{"TR", "RT", "TR"}, []int{7, 1, 3}, "RCJA")
}
