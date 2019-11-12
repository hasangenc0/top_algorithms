package binary_search_tree

import (
	"fmt"
	"math/rand"
	"testing"
)

var tree = Node{}

func InsertBenchmark(b *testing.B) {
	min := 1
	max := 20

	for n := 0; n < b.N; n++ {
		tree.Insert(rand.Intn(max - min) + min)
	}
}

func SearchBenchmark(b *testing.B) {
	min := 1
	max := 20

	for n := 0; n < b.N; n++ {
		tree.Search(rand.Intn(max - min) + min)
	}
}

func DeleteBenchmark(b *testing.B) {
	min := 1
	max := 20

	for n := 0; n < b.N; n++ {
		tree.Delete(rand.Intn(max - min) + min)
	}
}

func Benchmark() {
	// Insert Bench
	insertB := testing.Benchmark(InsertBenchmark)
	searchB := testing.Benchmark(SearchBenchmark)
	deleteB := testing.Benchmark(DeleteBenchmark)

	fmt.Print("Bst Insert Test: ")
	fmt.Println(insertB)

	fmt.Print("Bst Search Test: ")
	fmt.Println(searchB)

	fmt.Print("Bst Delete Test: ")
	fmt.Println(deleteB)
}