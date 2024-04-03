package main

import (
	"fmt"
)

func main() {
	var word string
	words := make([]string, 0)
	counts := make(map[rune]int)

	fmt.Println("英単語を入力してください（endと入力するとループを終了します）：")

	for {
		fmt.Scan(&word)
		if word == "end" {
			break
		}
		words = append(words, word)
		countLetters(word, counts)
	}

	fmt.Println("\n入力した英単語：")
	for _, word := range words {
		fmt.Println(word)
	}

	fmt.Println("\nアルファベットごとの文字数：")
	for letter, count := range counts {
		fmt.Printf("%c: %d\n", letter, count)
	}
}

func countLetters(word string, counts map[rune]int) {
	for _, letter := range word {
		counts[letter]++
	}
}
