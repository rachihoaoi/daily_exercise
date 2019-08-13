package sort

import (
	"fmt"
	"testing"
	"time"
)

func TestBubbleSort (t *testing.T) {
	array := []int{49, 38, 65, 97, 23, 22, 76, 1, 5, 8, 2, 0, -1, 22}
	array1 := []int{49, 38, 65, 97, 23, 22, 76, 1, 5, 8, 2, 0, -1, 22}
	BubbleSort(array)
	CookTailSort(array1)
	fmt.Println(array)
	fmt.Println(array1)
}

func BubbleSort (array []int) {
	t1 := time.Now()
	//fmt.Println(array)
	for i := 0; i < len(array) - 1; i++ {
		for j := 0; j < len(array) - i - 1; j++ {
			if array[j] > array[j + 1] {
				tmp := array[j]
				array[j] = array[j + 1]
				array[j + 1] = tmp
			}
		}
	}
	t2 := time.Now()
	fmt.Println("bubble:", t2.Sub(t1))
}

func CookTailSort (array []int) {
	t1 := time.Now()
	begin := 0
	end := len(array) - 1
	for begin < end {
		nbegin, nend := begin, end

		for i := begin; i < end; i++ {
			if array[i] > array[i + 1] {
				nend = i
				tmp := array[i]
				array[i] = array[i + 1]
				array[i + 1] = tmp
			}
		}
		if nbegin == nend {break}
		end = nend

		for j := end; j > begin; j-- {
			if array[j] < array[j - 1] {
				nbegin = j
				tmp := array[j]
				array[j] = array[j - 1]
				array[j - 1] = tmp
			}
		}
		if nbegin == begin {break}
		begin = nbegin
	}
	t2 := time.Now()
	fmt.Println("cookTail:", t2.Sub(t1))
}