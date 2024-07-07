package main

import "fmt"

// Remove関数は、スライスから指定された値を削除し、新しいスライスを返します
func Remove[T comparable](slice []T, value T) []T {
	for i, v := range slice {
		if v == value {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func main() {
	intSlice := []int{1, 2, 3, 4, 5}
	intSlice = Remove(intSlice, 3)
	fmt.Println(intSlice) // 出力: [1 2 4 5]

	stringSlice := []string{"apple", "banana", "cherry"}
	stringSlice = Remove(stringSlice, "banana")
	fmt.Println(stringSlice) // 出力: [apple cherry]
}
