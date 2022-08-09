package main

import (
	"fmt"
	"time"
)

func solution(progresses []int, speeds []int) []int {

	deploy := make([]int, len(progresses))

	for i, _ := range progresses {
		deploy[i] = (100 - progresses[i]) / speeds[i]
		if deploy[i]*speeds[i] != 100-progresses[i] {
			deploy[i] += 2
		}
	}

	if len(deploy) == 0 {
		return []int{-1}
	}

	result := []int{1}
	big := deploy[0]
	for i := 1; i < len(deploy); i++ {
		if deploy[i] <= big {
			result[len(result)-1]++
		} else {
			result = append(result, 1)
			big = deploy[i]
		}
	}

	return result
}

func report(progresses []int, speeds []int, desireResult []int) {

	input := fmt.Sprintf("Input:[%v %v]", progresses, speeds)

	start := time.Now()

	result := solution(progresses, speeds)

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
	report([]int{93, 30, 55}, []int{1, 30, 5}, []int{2, 1})
	report([]int{95, 90, 99, 99, 80, 99}, []int{1, 1, 1, 1, 1, 1}, []int{1, 3, 2})
}
