package algorithm

import (
	"fmt"
	"math"
	"testing"
)

func TestIntReverse(t *testing.T) {
	fmt.Println(reverse(-123))
}

func reverse(x int) int {
	ret := 0
	for x != 0 {
		end :=  x % 10
		fmt.Println(end)
		x = x / 10
		ret = ret * 10 + end
		if ret < math.MinInt32 || ret > math.MaxInt32 {
			return 0
		}
	}
	return ret
}
