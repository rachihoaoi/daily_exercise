package algorithm

import (
	"testing"
)

func TestRotate(t *testing.T) {
	rotate([][]int{
		{1 ,2 ,3 ,4 ,5},
		{6 ,7 ,8 ,9 ,10},
		{11,12,13,14,15},
		{16,17,18,19,20},
		{21,22,23,24,25},
		//{1,2,3},
		//{4,5,6},
		//{7,8,9},
	})
}
func rotate(matrix [][]int)  {
	length := len(matrix)
	for i := 0; i < length / 2; i++ {
		for j := i; j < length - 1 - i; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[length - 1 - j][i]
			matrix[length - 1 - j][i] = matrix[length - 1 - i][length - 1 -j]
			matrix[length - 1 - i][length - 1 -j] = matrix[j][length - 1 - i]
			matrix[j][length - 1 - i] = tmp

		}
	}
}