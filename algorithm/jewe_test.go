package algorithm

import (
	"fmt"
	"testing"
)

func TestJewe(t *testing.T) {
	fmt.Println(numJewelsInStones("zs", "ZZzaaAAAzAAcsafas"))
}

func numJewelsInStones(J string, S string) int {
	hMap := make(map[int32]int)
	sum := 0
	for _,v := range J {
		hMap[v] = 1
	}
	for _, v := range S {
		if hMap[v] != 0 {
			sum ++
		}
	}

	//fmt.Println()
	return sum
}