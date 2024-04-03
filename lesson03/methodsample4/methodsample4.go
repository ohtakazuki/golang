package main

import "fmt"

func main() {
	number1 := 100
	printNum(number1)
	fmt.Println("number1の値は", number1)
}

func printNum(num1 int) {
	num1 *= 10
	fmt.Println("num1の10倍は：", num1)
}
