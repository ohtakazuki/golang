package main

import "fmt"

func main() {
  var arr1 [5]int  // 長さ5の整数型の配列を宣言
  fmt.Println("arr1:", arr1)

	arr1[0] = 1  // 配列の最初の要素に1を代入
	arr1[1] = 2  // 配列の2番目の要素に2を代入
	fmt.Println("arr1:", arr1)

	arr2 := [5]int{1, 2, 3, 4, 5}  // 長さ5の整数型の配列を宣言し、初期値を設定
  fmt.Println("arr2:", arr2)

	arr1Len := len(arr1) // 配列arr1の長さを取得
	fmt.Println("arr1の長さ:", arr1Len)  
}