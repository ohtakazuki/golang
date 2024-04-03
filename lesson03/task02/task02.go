package main

import (
	"fmt"
	"strconv"
)

func main() {
	defer func() {
		fmt.Println("処理を終了します")
	}()

	var a, b int
	var err error

	fmt.Print("割られる数を入力してください：")
	if a, err = strconv.Atoi(readLine()); err != nil {
		fmt.Println("エラー：数値を入力してください")
		return
	}

	fmt.Print("割る数を入力してください：")
	if b, err = strconv.Atoi(readLine()); err != nil {
		fmt.Println("エラー：数値を入力してください")
		return
	}

	if b == 0 {
		fmt.Println("エラー：0で割り算しないでください")
		return
	}

	c := float64(a) / float64(b)
	fmt.Printf("%d ÷ %d = %.1f\n", a, b, c)
}

func readLine() string {
	var line string
	fmt.Scanln(&line)
	return line
}
