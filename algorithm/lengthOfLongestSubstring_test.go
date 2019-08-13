package algorithm

import (
	"fmt"
	"testing"
)

func TestLengthOfSubString(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("asdasdz"))
	}

func lengthOfLongestSubstring(s string) int {
	byteArray := []byte(s)
	indexArray := make([]int,128)
	var Max,Num int
	for i,j := 0, 0;j < len(s) && i<len(s); j++ {
		if indexArray[byteArray[j]] > i {
			i = indexArray[byteArray[j]]
		}
		Num = j - i + 1
		if Num > Max {
			Max = Num
		}
		indexArray[byteArray[j]] = j + 1
	}

	return Max
}
//
//func TestLengthOfSubString(t *testing.T) {
//	fmt.Println(ifMult([]byte("c")))
//}