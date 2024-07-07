package main

import "fmt"

// Addableインターフェースは、Add メソッドを持つ型を表します
type Addable interface {
	Add(other Addable) Addable
}

// Sum関数は、Addableインターフェースを満たす型の配列の要素を合計します
func Sum[T Addable](values []T) T {
	var sum T
	for _, value := range values {
		sum = sum.Add(value).(T)
	}
	return sum
}

// MyInt型は、int型をラップし、Addableインターフェースを実装します
type MyInt int

// Add メソッドは、MyInt型の値を加算します
func (a MyInt) Add(other Addable) Addable {
	return a + other.(MyInt)
}

func main() {
	values := []MyInt{1, 2, 3, 4, 5}
	sum := Sum(values)
	fmt.Println(sum) // 出力: 15
}
