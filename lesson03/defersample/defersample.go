package main

import (
	"bufio"
	"fmt"
	"os"
)

func processFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func main() {
	err := processFile("example.txt")
	if err != nil {
		fmt.Println("ファイル処理に失敗しました:", err)
		return
	}
	fmt.Println("ファイル処理が完了しました")
}
