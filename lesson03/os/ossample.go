package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "3.14159"
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("文字列を浮動小数点数に変換できませんでした:", err)
		return
	}
	fmt.Printf("%s -> %f\n", str, f)

}
