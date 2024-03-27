package main

import "fmt"

func main() {
	// 変数の明示的な宣言
	var name string = "山田太郎"
	var age int = 25
	var birthplace string = "東京都"
	var hobby string = "読書"

	// 自己紹介の出力
	fmt.Printf("はじめまして、%vと申します。\n", name)
	fmt.Printf("年齢は%v歳で、%v出身です。\n", age, birthplace)
	fmt.Printf("趣味は%vです。\n", hobby)
	fmt.Println("よろしくお願いします。")
}
