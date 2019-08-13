package algorithm

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	fmt.Println(isPalindrome(1))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	array := make([]int,0)
	for x != 0 {
		array = append(array, x % 10)
		x /= 10
	}
	for i:=0;i<len(array)/2;i++{
		if array[i] != array[len(array)-i-1]{
			return false
		}
	}
	return true
}