package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := [5]int{3, 1, 4, 5, 2}

	// 昇順にソートする
	sort.Slice(numbers[:], func(i, j int) bool {
		return numbers[i] < numbers[j]
	})

	fmt.Println("Sorted numbers:", numbers)
}
