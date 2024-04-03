package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// ポインタを使ってミュータブルな値を変更する
	p := &Person{Name: "John", Age: 30}
	fmt.Println(p) // Output: &{John 30}
	p.Age = 31
	fmt.Println(p) // Output: &{John 31}

	// nilポインタへのアクセス
	var p2 *Person
	fmt.Println(p2) // Output: <nil>
	// p2.Age = 32    // ランタイムパニック：nilポインタへの間接参照

	// nilチェック
	p2 = &Person{}
	if p2 != nil {
		p2.Age = 32
	}
}
