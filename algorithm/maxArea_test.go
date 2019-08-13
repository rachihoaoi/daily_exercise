package algorithm

import (
	"fmt"
	"testing"
)

func TestMaxArea (t *testing.T){
	fmt.Println(maxArea([]int{2,3,4,5,18,17,6}))
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	var currentVolume, maxVolume int
	for i,j := 0,len(height) - 1;i < len(height) && j > 0; {
		if height[i] <= height[j] {
			currentVolume = (j - i ) * height[i]
			i ++
		} else {
			currentVolume = (j - i ) * height[j]
			j --
		}
		if currentVolume > maxVolume {
			maxVolume = currentVolume
		}
		if (i - j) == 0 {
			break
		}
	}
	return maxVolume
}