package main

import (
	"fmt"
	"time"
)

type Bridge struct {
	on          []int
	totalWeight int
	totalSec    int
}

func makeBridge(l int) *Bridge {
	bridge := Bridge{}
	bridge.on = make([]int, l)
	bridge.totalWeight = 0
	bridge.totalSec = 0

	return &bridge
}

func (b *Bridge) Move(w int) int {
	out := b.on[0]

	b.on = b.on[1:]
	b.on = append(b.on, w)

	b.totalWeight += w - out

	if b.totalWeight > 0 {
		b.totalSec++
	}

	return out
}

// Weight Move시 무게 계산
func (b *Bridge) Weight(w int) int {
	return b.totalWeight - b.on[0] + w
}

// Left 나머지 시간 계산
func (b *Bridge) Left() {
	for i := len(b.on) - 1; i >= 0; i-- {
		if b.on[i] == 0 {
			continue
		}

		b.totalSec += i + 1
		break
	}
}

func solution(bridge_length int, weight int, truck_weights []int) int {

	onBridge := makeBridge(bridge_length)

	for _, w := range truck_weights {

		for onBridge.Weight(w) > weight {
			onBridge.Move(0)
		}

		onBridge.Move(w)
	}

	onBridge.Left()

	return onBridge.totalSec
}

func report(bridge_length int, weight int, truck_weights []int, desireResult int) {

	input := fmt.Sprintf("Input:[%v, %v, %v]", bridge_length, weight, truck_weights)

	start := time.Now()

	result := solution(bridge_length, weight, truck_weights)

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
	report(2, 10, []int{7, 4, 5, 6}, 8)
	report(100, 100, []int{10}, 101)
	report(100, 100, []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, 110)
	report(5, 5, []int{2, 2, 2, 2, 1, 1, 1, 1, 1}, 19)
}
