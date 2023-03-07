package main

import (
	"fmt"

	"github.com/divilla/functional-programming/slices"
)

func main() {
	for i := 0; i < 10000; i++ {
		fmt.Println(
			slices.FromIntSlice([]int{1, 2, 3, 4, 5, 6}).
				Map(func(v, k int) int {
					return v * 3
				}).
				Filter(func(value, index int) bool {
					return value%2 == 0
				}).
				ToIntSlice(),
		)
	}
}
