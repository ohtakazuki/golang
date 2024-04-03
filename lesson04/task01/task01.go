package main

import (
	"fmt"
)

func main() {
	var scores []int
	var score int

	for {
		fmt.Print("点数を入力してください：")
		fmt.Scan(&score)

		if score == -1 {
			break
		}

		scores = append(scores, score)
	}

	if len(scores) > 0 {
		var sum int
		for _, s := range scores {
			sum += s
		}
		average := float64(sum) / float64(len(scores))
		fmt.Printf("%d人のテストの平均点は%.1f点です\n", len(scores), average)
	} else {
		fmt.Println("点数が入力されませんでした")
	}
}
