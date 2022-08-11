package main

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

type Bigger []string

func (b Bigger) Len() int {
	return len(b)
}

func (b Bigger) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b Bigger) Less(i, j int) bool {
	ij, _ := strconv.Atoi(b[i] + b[j])
	ji, _ := strconv.Atoi(b[j] + b[i])
	return ij < ji
}

func solution(numbers []int) string {

	allZero := true
	for _, n := range numbers {
		if n != 0 {
			allZero = false
			break
		}
	}

	if allZero {
		return "0"
	}

	s := []string{}
	for _, v := range numbers {
		s = append(s, fmt.Sprintf("%d", v))
	}

	sum := ""
	sort.Sort(Bigger(s))

	for _, v := range s {
		sum = v + sum
	}

	return sum
}

func report(numbers []int, desireResult string) {

	input := fmt.Sprintf("Input:[%v]", numbers)

	start := time.Now()

	result := solution(numbers)

	duration := time.Since(start)

	ox := ""
	if result == desireResult {
		ox = fmt.Sprintf("[O]")
	} else {
		ox = fmt.Sprintf("[X]")
	}

	fmt.Printf("%s %s result:[%v] duration:%v\n", ox, input, result, duration)
}

func main() {

	report([]int{232, 23}, "23232")
	report([]int{212, 21}, "21221")
	report([]int{70, 0, 0, 0, 0}, "700000")
	report([]int{6, 10, 2}, "6210")
	report([]int{3, 30, 34, 5, 9}, "9534330")
	report([]int{0, 0, 0, 0}, "0")
}
