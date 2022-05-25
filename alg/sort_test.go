package alg

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	arr := []int{1,2,2,8,2,4,5,3,7,9,8}
	HeapSort(arr)
	fmt.Println(arr)
}

func TestMergeSort(t *testing.T) {
	arr := []int{1,2,2,8,2,4,5,3,7,9,8}
	arr = MergeSort(arr)
	fmt.Println(arr)
}

func TestQuickSort(t *testing.T) {
	arr := []int{1,2,2,8,2,4,5,3,7,9,8}
	QuickSort(arr)
	fmt.Println(arr)
}

func TestFindMaxn(t *testing.T) {
	arr := []int{1,2,2,8,2,4,5,3,7,9,8}
	fmt.Println(FindMaxn(arr, 5))
}