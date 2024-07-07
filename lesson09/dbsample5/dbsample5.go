package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Book struct {
	ID        int
	Title     string
	Price     int
	CreatedAt time.Time
}

func main() {
	// データベースに接続
	db, err := sql.Open("mysql", "root:_MySQLPassword1@tcp(localhost:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("データベースへの接続に失敗しました:", err)
		return
	}
	defer db.Close()

	// トランザクションを開始
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("トランザクションの開始に失敗しました:", err)
		return
	}

	// プリペアドステートメントを使用してデータを更新
	updateStmt, err := tx.Prepare("UPDATE books SET title = ?, price = ?, created_at = ? WHERE id = ?")
	if err != nil {
		fmt.Println("プリペアドステートメントの作成に失敗しました:", err)
		tx.Rollback()
		return
	}
	defer updateStmt.Close()

	// 更新するデータを準備
	book := Book{
		ID:        16,
		Title:     "Updated Book",
		Price:     1500,
		CreatedAt: time.Now(),
	}

	// データを更新
	result, err := updateStmt.Exec(book.Title, book.Price, book.CreatedAt, book.ID)
	if err != nil {
		fmt.Println("データの更新に失敗しました:", err)
		tx.Rollback()
		return
	}

	// 変更された行数を取得
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("変更された行数の取得に失敗しました:", err)
		tx.Rollback()
		return
	}

	if rowsAffected == 0 {
		fmt.Println("該当するレコードが存在しません")
		tx.Rollback()
		return
	}

	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		fmt.Println("トランザクションのコミットに失敗しました:", err)
		return
	}

	fmt.Println("データの更新が成功しました")
}
