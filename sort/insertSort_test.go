package sort

import (
	"fmt"
	"testing"
	"time"
)

func TestInsertSort (t *testing.T) {
	t1 := time.Now()
	array := []int{49, 38, 65, 97, 23, 22, 76, 1, 5, 8, 2, 0, -1, 22}
	InsertSort(array)
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
	fmt.Println(array)
}

func InsertSort(array []int) {
	for j := 1; j < len(array); j++ {
		// 判断需要插入的值
		if array[j] < array[j-1] {
			// 暂存需插入的值
			key := array[j]
			// 当前值的位置
			i := j - 1
			// 遍历当前值之前的位置，若大于这个值则说明被插入位置还在前面，目前元素向后一动一格
			for ; i >= 0 && array[i] > key; i-- {
				// 一动元素
				array[i+1] = array[i]
			}
			// 得到插入位置i，插入元素
			array[i+1] = key
		}
	}
}