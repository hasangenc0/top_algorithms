package utility

import "math/rand"

func GenerateInt(max int, min int) int {
	return rand.Intn(max-min)
}

func GenerateIntArray() (arr []int) {

	size := GenerateInt(100, 10)

	for i := 1; i <= size; i++ {
		arr = append(arr, GenerateInt(1000, 0))
	}
	return arr
}
