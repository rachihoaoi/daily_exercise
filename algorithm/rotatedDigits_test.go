package algorithm

import (
	"fmt"
	"testing"
)

func TestRotateDigits (t *testing.T) {
	fmt.Println(rotatedDigits(100))
	fmt.Println(8000 / 1111)
}


func rotatedDigits(N int) int {
	return N / 1111 + N / 111 + N / 11 + N / 1
}
