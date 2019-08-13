package sort

import (
	"fmt"
	"testing"
	"time"
)

func TestQuickSort (t *testing.T) {
	t1 := time.Now()
	array := []int{49, 38, 65, 97, 23, 22, 76, 1, 5, 8, 2, 0, -1, 22}
	quickSort(array, 0, len(array) - 1)
	fmt.Println(array)
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
}

func quickSort(array []int, begin int, end int) {
	if begin < end {
		index := getIndex(array, begin, end)
		// 递归排序
		quickSort(array, begin, index - 1)
		quickSort(array, index + 1, end)
	}

}

func getIndex(array []int, begin int, end int) int {
	tmp := array[begin]
	for begin < end {
		// 当队尾的元素大于等于基准数据时,向前挪动end指针
		for begin < end && array[end] >= tmp {
			end --
		}
		// 如果队尾元素小于tmp了,需要将其赋值给begin
		array[begin] = array[end]
		// 当队首元素小于等于tmp时,向前挪动begin指针
		for begin < end && array[begin] <= tmp {
			begin ++
		}
		// 当队首元素大于tmp时,需要将其赋值给end
		array[end] = array[begin]
	}
	array[begin] = tmp
	return begin
}