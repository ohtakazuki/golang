package main

import (
	"fmt"
	"os"
	"sync"
)

func readFile(filename string, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("ファイル '%s' の読み込みエラー: %v\n", filename, err)
		ch <- 0
		return
	}

	count := len(string(content))
	ch <- count
}

func main() {
	filenames := []string{"file1.txt", "file2.txt", "file3.txt"}
	ch := make(chan int)
	var wg sync.WaitGroup

	// 各ファイルを読み込むためのゴルーチンを作成
	for _, filename := range filenames {
		wg.Add(1)
		go readFile(filename, ch, &wg)
	}

	// ゴルーチンの完了を待つ
	go func() {
		wg.Wait()
		close(ch)
	}()

	// 各ゴルーチンから文字数を受け取り、合計を計算
	totalCount := 0
	for count := range ch {
		totalCount += count
	}

	// 最終的な文字数を表示
	fmt.Printf("合計文字数: %d\n", totalCount)
}
