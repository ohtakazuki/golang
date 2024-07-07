package main

import "fmt"

// ジェネリクスを使った汎用的な構造体の例
type Array[T any] struct {
	data []T
}

// Array構造体に要素を追加するメソッド
func (a *Array[T]) Append(value T) {
	a.data = append(a.data, value)
}

// Array構造体の要素を表示するメソッド
func (a *Array[T]) Print() {
	fmt.Printf("Array: %v\n", a.data)
}

func main() {
	// int型の配列を作成
	intArray := &Array[int]{}
	intArray.Append(1)
	intArray.Append(2)
	intArray.Append(3)
	intArray.Print()

	// string型の配列を作成
	strArray := &Array[string]{}
	strArray.Append("Hello")
	strArray.Append("World")
	strArray.Print()
}
