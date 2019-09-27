package sort

import (
	utils "github.com/hasangenc0/top_algorithms/pkg/utility"

	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{132,21,24,56,6,87,3,1}
	sortedArray := []int{1,3,6,21,24,56,87,132}
	MergeSort(arr)

	if !utils.AreIntArraysEqual(arr, sortedArray) {
		t.Errorf("Array = %d; wanted %d", arr, sortedArray)
	}
}
