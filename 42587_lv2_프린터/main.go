package main

import (
	"fmt"
	"time"
)

type printer struct {
	index    int
	priority int
}

type Queue []printer

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Push(data printer) {
	*q = append(*q, data)
}

func (q *Queue) Pop() printer {
	if q.IsEmpty() {
		return printer{-1, 0}
	}
	data := (*q)[0]
	*q = (*q)[1:]
	return data
}

func (q *Queue) MaxPriority() int {
	max := 0
	for _, p := range *q {
		if max < p.priority {
			max = p.priority
		}
	}

	return max
}

func solution(priorities []int, location int) int {

	result := []printer{} //  index: priorities

	q := Queue{}
	for i, p := range priorities {
		item := printer{index: i, priority: p}
		q.Push(item)
	}

	for q.IsEmpty() == false {
		hi := q.MaxPriority()
		p := q.Pop()
		if hi == p.priority {
			result = append(result, p)
		} else {
			q.Push(p)
		}
	}

	for i, p := range result {
		if p.index == location {
			return i + 1
		}
	}

	return 1
}

func report(priorities []int, location int, desireResult int) {

	start := time.Now()

	result := solution(priorities, location)

	duration := time.Since(start)

	if result == desireResult {
		fmt.Printf("[O] ")
	} else {
		fmt.Printf("[X] ")
	}

	fmt.Printf("Input:[%v, %v], result:[%v], duration:%v\n", priorities, location, result, duration)
}

func main() {

	report([]int{2, 1, 3, 2}, 2, 1)
	report([]int{1, 1, 9, 1, 1, 1}, 0, 5)
}
