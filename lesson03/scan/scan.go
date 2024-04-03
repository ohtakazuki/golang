package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Print("名前を入力してください: ")
	fmt.Scan(&name)
	fmt.Print("年齢を入力してください: ")
	fmt.Scan(&age)
	fmt.Printf("%sは%d歳です。\n", name, age)
}
