package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Book struct {
	ID        int
	Title     string
	Price     int
	CreatedAt time.Time
}

func main() {
	// データベースに接続
	dsn := "root:_MySQLPassword1@tcp(localhost:3306)/bookstore?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("データベースへの接続に失敗しました:", err)
		return
	}

	// ユーザーからタイトルを入力してもらう
	var title string
	fmt.Print("検索するタイトルを入力してください: ")
	fmt.Scan(&title)

	// クエリを実行
	var book Book
	result := db.Where("title = ?", title).First(&book)

	// 結果を処理
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			fmt.Println("指定されたタイトルのレコードが見つかりませんでした")
			return
		} else {
			fmt.Println("クエリの実行に失敗しました:", result.Error)
			return
		}
	}

	// 結果を出力
	fmt.Println(book.ID, book.Title, book.Price, book.CreatedAt)
}
