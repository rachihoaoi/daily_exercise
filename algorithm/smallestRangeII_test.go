package algorithm

import (
	"fmt"
	"sort"
	"testing"
)

func TestSmallestRangeII (t *testing.T) {
	fmt.Println(smallestRangeII([]int{1,3,4,6,8,3,4,2,56,7}, 10))
}

func SolveMax(a int,b int ) int {
	if a > b {
		return a
	}
	return  b
}

func SolveMin(a int,b int ) int {
	if a > b {
		return b
	}
	return  a
}

func smallestRangeII(A []int, K int) int {
	sort.Ints(A)
	lenth := len(A)
	res := A[lenth - 1] - A[0]
	for i := 1;i<lenth;i++ {
		min := SolveMin(A[0] + K, A[i] - K)
		max := SolveMax(A[lenth - 1] - K, A[i -1] + K)
		res = SolveMin(max-min, res)
	}
	return res
}