package main

import "fmt"

func main() {
	numbers := []int{100} // スライスを初期化
	printNum(numbers)     // スライスをそのまま渡す
	fmt.Println("最初の要素の値は", numbers[0])
}

func printNum(nums []int) {
	nums[0] *= 10 // スライスの最初の要素を10倍にする
	fmt.Println("最初の要素の10倍は", nums[0])
}
