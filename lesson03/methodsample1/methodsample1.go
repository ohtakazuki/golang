package main

import "fmt"

func main() {
	number1 := 100
	number2 := 10
	sumMethod1(number1, number2)

	result2 := sumMethod2(number1, number2)
	fmt.Println("sumMethod2の結果は", result2)

	// ----- 追記（ここから）-----
	a := 1
	b := 2
	fmt.Println("a=", a, ", b=", b)
	a, b = swap(a, b)
	fmt.Println("a=", a, ", b=", b)
	// ----- 追記（ここまで）-----
}

func sumMethod1(num1 int, num2 int) {
	result := num1 + num2
	fmt.Println("sumMethod1の結果は", result)
}

func sumMethod2(num1 int, num2 int) int {
	result := num1 + num2
	return result
}

// ----- 追記（以下の関数）-----
func swap(x, y int) (int, int) {
	return y, x
}
