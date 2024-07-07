package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

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

	// プリペアドステートメントを使用してデータを削除
	deleteStmt, err := tx.Prepare("DELETE FROM books WHERE id = ?")
	if err != nil {
		fmt.Println("プリペアドステートメントの作成に失敗しました:", err)
		tx.Rollback()
		return
	}
	defer deleteStmt.Close()

	// 削除するデータのIDを指定
	bookID := 81

	// データを削除
	result, err := deleteStmt.Exec(bookID)
	if err != nil {
		fmt.Println("データの削除に失敗しました:", err)
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

	fmt.Println("データの削除が成功しました")
}
