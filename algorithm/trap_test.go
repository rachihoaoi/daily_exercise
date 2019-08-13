package algorithm

import (
	"fmt"
	"testing"
)

func TestTrap (t *testing.T) {
	fmt.Println(trap([]int{9,8,10, 0, 2}))
}


func trap(height []int) int {
	sum := 0
	topIndex, _ := getTop(height)
	leftMax := 0
	for i,j  := 0, len(height) - 1; i < topIndex && j > topIndex; {
		if height[i] > leftMax {
			leftMax = height[i]
		}
		sum += leftMax - height[i]
	}

	return sum
}

func getTop(height []int) (int, int) {
	index, value := 0,0
	for k, v := range height {
		if v > value {
			value = v
			index = k
		}
	}
	return index, value
}


func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}