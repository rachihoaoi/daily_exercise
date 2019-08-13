package sort

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	array := []int{ 2, 0, -1, 22,2,3,7,3,7}
	heapSort(array)
	fmt.Println(array)
}

func heapSort (array []int) {
	arrLen := len(array)
	GenerateMaxHeap(array, arrLen)
	for i := arrLen - 1; i >=0 ; i-- {
		swap(array, 0, i)
		arrLen -= 1
		heapify(array, 0, arrLen)
	}

}

func GenerateMaxHeap (array []int, length int) {
	for i := length / 2; i >=0; i-- {
		heapify(array, i, length)
	}
}

func heapify(array []int, notChildIndex int, arrLen int) {
	leftChild := 2 * notChildIndex + 1
	rightChild := 2 * notChildIndex + 2
	root := notChildIndex
	if leftChild < arrLen && array[leftChild] > array[root] {
		root = leftChild
	}
	if rightChild < arrLen && array[rightChild] > array[root] {
		root = rightChild
	}
	if root != notChildIndex {
		swap(array, notChildIndex, root)
		heapify(array, notChildIndex, arrLen)
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}