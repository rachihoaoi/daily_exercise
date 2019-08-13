package algorithm

import (
	"fmt"
	"sort"
	"testing"
)

func TestTwoSum(t *testing.T) {
	//fmt.Println(twoSum([]int{2, 7, 11, 15}, 17))
	// 1 4 4 4 6 7 8 9
	//fmt.Println(arrayPairSum([]int{1,4,7,4,6,8,4,9}))
	fmt.Println(smallestRangeII([]int{1,4,6,8,4,7,3,1,7,9}, 1))
}

func twoSum(nums []int, target int) []int {
	lenth := len(nums)
	hsMap := make(map[int]int)
	for i := 0; i < lenth; i ++ {
		if _,ok := hsMap[target - nums[i]];ok {
			return []int{hsMap[target - nums[i]], i}
		}
		hsMap[nums[i]] = i
	}
	return []int{99,99}
}


func arrayPairSum(nums []int) int {
	sum := 0
	sort.Ints(nums)
	for i:=0; i< len(nums); i+=2 {
		sum += nums[i]
	}
	return sum

}
