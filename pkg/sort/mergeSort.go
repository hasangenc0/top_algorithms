package sort

import (
	utils "github.com/hasangenc0/top_algorithms/pkg/utility"

	"fmt"
	"testing"
)

/* Merge two halves */
func merge(first []int, second []int) {

	temp := append(first, second...)

	for i := 0; i < len(temp); i++ {
		for j := i; j < len(temp); j++ {
			if temp[i] > temp[j] {
				temp[i], temp[j] = temp[j], temp[i]
			}
		}
	}

}

/* Sort arrays */
func MergeSort(arr []int) {
	length := len(arr)

	if length <= 1 {
		return
	}

	middle := length / 2
	first := arr[0:middle]
	second := arr[middle:length]

	MergeSort(first)
	MergeSort(second)

	merge(first, second)
}

func mergeSortBenchmark(b *testing.B) {
	for n := 0; n < b.N; n++ {
		arr := utils.GenerateIntArray()
		MergeSort(arr)
	}

}

func Benchmark() {
	// Merge Sort Bench
	mbr := testing.Benchmark(mergeSortBenchmark)

	fmt.Print("Merge Sort Bench Test: ")
	fmt.Println(mbr)
}
