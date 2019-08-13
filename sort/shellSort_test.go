package sort

import (
	"fmt"
	"testing"
)

func TestShellSort (t *testing.T) {
	array := []int{49, 38, 65, 97, 23, 22, 76, 1, 5, 8, 2, 0, -1, 22}
	ShellSort(array, len(array) / 2)
	fmt.Println(array)
}

func ShellSort (arr []int, batchSize int) {
	if batchSize < 1 {
		return
	}
	// 遍历所有batch
	for i := 0; i < batchSize; i++ {
		// 遍历得到每个batch内部有多少个元素
		for j := 1; batchSize * j + i < len(arr); j++ {
			// 对于每一个新增的元素，与前一个排序比较大小
			for n := j; n > 0; n -- {
				pre := batchSize * (n - 1) + i
				next := batchSize * n + i
				if arr[next] < arr[pre] {
					arr[next], arr[pre] = arr[pre], arr[next]
				}
			}
		}
	}
	// 递归调用直至batchsize < 1
	ShellSort(arr, batchSize / 2)
}
