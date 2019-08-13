package algorithm

import (
	"fmt"
	"testing"
)

func TestRob (t *testing.T) {
	fmt.Println(rob([]int{7,8,3}))
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	sum1, sum2 := nums[0], nums[1]

	for lastIndex := 2; lastIndex < len(nums); lastIndex ++ {
		tmp := sum1

		if sum2 > sum1 {
			sum1 = sum2
		}

		sum2 = tmp + nums[lastIndex]
	}
	if sum1 > sum2 {
		return sum1
	}
	return sum2
}
