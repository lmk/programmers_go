package main

import (
	"fmt"
	"time"
)

/*
type Navigation struct {
	Direction []int // 0: minus, 1: plus

}

type Node struct {
	value int
	plus  *Node
	minus *Node
}

type BTree struct {
	node *Node
}

func (n *Node) Insert(v int) {

	if n.plus == nil {
		n.plus = &Node{value: v}
	} else {
		n.plus.Insert(v)
	}

	if n.minus == nil {
		n.minus = &Node{value: v}
	} else {
		n.minus.Insert(v)
	}
}

func (n *Node) Print(depth int, dir int) {
	if n == nil {
		return
	}

	sign := "+"
	if dir == 0 {
		sign = "-"
	}

	fmt.Printf("%s %s%d\n", strings.Repeat("\t", depth), sign, n.value)

	n.plus.Print(depth+1, 1)
	n.minus.Print(depth+1, 0)
}

func (t *BTree) Insert(v int) {
	if t.node == nil {
		t.node = &Node{value: 0}
		t.node.Insert(v)
	} else {
		t.node.Insert(v)
	}
}

func (t *BTree) Print() {
	if t.node == nil {
		fmt.Println("nil")
	}

	t.node.Print(0, 1)
}

func solution(numbers []int, target int) int {

	// insert tree
	t := &BTree{}
	for _, v := range numbers {
		t.Insert(v)
	}

	t.Print()

	// navigate
	//dir := []int{}

	return 0
}
*/
var gSame int
var gNumbers []int
var gTarget int

func calc(i int, sign int, sum int) {
	if len(gNumbers) <= i {
		if sum == gTarget && sign > 0 {
			gSame++
		}

		return
	}

	sum += gNumbers[i] * sign

	calc(i+1, 1, sum)
	calc(i+1, -1, sum)
}

func solution(numbers []int, target int) int {

	gSame = 0
	gNumbers = numbers
	gTarget = target

	calc(0, 1, 0)
	calc(0, -1, 0)

	return gSame
}

func report(numbers []int, target int, desireResult int) {

	input := fmt.Sprintf("Input:[%v %v]", numbers, target)

	start := time.Now()

	result := solution(numbers, target)

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

	//report([]int{1, 2, 3, 4, 5, 6, 7}, 6, 2)
	report([]int{1, 1, 1, 1, 1}, 3, 5)
	report([]int{4, 1, 2, 1}, 4, 2)
}
