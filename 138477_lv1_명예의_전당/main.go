package main

import (
	"fmt"
	"sort"
	"time"
)

// 명예의 전당
type Fame struct {
	score []int
}

func (t *Fame) Push(score int, size int) {

	t.score = append(t.score, score)

	sort.Slice(t.score, func(i int, j int) bool {
		return t.score[i] > t.score[j]
	})

	if size < len(t.score) {
		t.score = t.score[:size]
	}
}

func (t *Fame) Low() int {
	return t.score[len(t.score)-1]
}

func solution(k int, score []int) []int {

	answer := make([]int, len(score))

	var fame Fame
	for i, s := range score {
		fame.Push(s, k)
		answer[i] = fame.Low()
	}

	return answer

}

func report(k int, score []int, desireResult []int) {

	input := fmt.Sprintf("Input:[%v, %v]", k, score)

	start := time.Now()

	result := solution(k, score)

	duration := time.Since(start)

	isSame := true
	for i := range result {
		if result[i] != desireResult[i] {
			isSame = false
			break
		}
	}

	if isSame {
		fmt.Printf("[O] %s result:[%v] duration:%v\n", input, result, duration)
	} else {
		fmt.Printf("[X] %s result:[%v]->[%v] duration:%v\n", input, result, desireResult, duration)
	}
}

func main() {
	report(3, []int{10, 100, 20, 150, 1, 100, 200}, []int{10, 10, 10, 20, 20, 100, 100})
	report(4, []int{0, 300, 40, 300, 20, 70, 150, 50, 500, 1000}, []int{0, 0, 0, 0, 20, 40, 70, 70, 150, 300})
}
