package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Book struct {
	ID         int
	Title      string
	Price      int
	Created_at time.Time
}

func main() {
	// データベースに接続
	db, err := sql.Open("mysql", "root:_MySQLPassword1@tcp(localhost:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("データベースへの接続に失敗しました:", err)
		return
	}
	defer db.Close()

	// ユーザーからタイトルを入力してもらう
	var title string
	fmt.Print("検索するタイトルを入力してください: ")
	fmt.Scan(&title)

	// クエリを実行
	rows, err := db.Query("SELECT * FROM books WHERE title = ?", title)
	if err != nil {
		fmt.Println("クエリの実行に失敗しました:", err)
		return
	}
	defer rows.Close()

	// 結果を処理
	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Price, &book.Created_at)
		if err != nil {
			fmt.Println("結果の取得に失敗しました:", err)
			return
		}
		books = append(books, book)
	}

	// エラーチェック
	err = rows.Err()
	if err != nil {
		fmt.Println("結果の処理中にエラーが発生しました:", err)
		return
	}

	// 結果を出力
	if len(books) == 0 {
		fmt.Println("指定されたタイトルのレコードが見つかりませんでした")
	} else {
		for _, book := range books {
			fmt.Println(book.ID, book.Title, book.Price, book.Created_at)
		}
	}
}
