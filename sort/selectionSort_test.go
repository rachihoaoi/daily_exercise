package sort

import (
	"fmt"
	"testing"
)

func TestSelectionSort (t *testing.T) {
	array := []int{49, 38, 65, 97, 23, 22, 76, 1, 5, 8, 2, 0, -1, 22}
	SelectionSort(array)
	fmt.Println(array)
}

func SelectionSort (array []int) {
	length := len(array)
	for i := 0; i < length; i++ {
		m := array[i]
		index := i
		for j := i + 1; j < length; j++ {
			if array[j] < m {
				m = array[j]
				index = j
			}
		}
		array[index], array[i] = array[i], array[index]
	}
}
