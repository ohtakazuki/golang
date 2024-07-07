package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Book struct {
	ID int
}

func main() {
	// データベースに接続
	dsn := "root:_MySQLPassword1@tcp(localhost:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("データベースへの接続に失敗しました:", err)
		return
	}

	// トランザクションを開始
	tx := db.Begin()

	// 削除するデータのIDを指定
	bookID := 4

	// データを削除
	result := tx.Delete(&Book{}, bookID)
	if result.Error != nil {
		fmt.Println("データの削除に失敗しました:", result.Error)
		tx.Rollback()
		return
	}

	// 変更された行数を確認
	if result.RowsAffected == 0 {
		fmt.Println("該当するレコードが存在しません")
		tx.Rollback()
		return
	}

	// トランザクションをコミット
	if err := tx.Commit().Error; err != nil {
		fmt.Println("トランザクションのコミットに失敗しました:", err)
		return
	}

	fmt.Println("データの削除が成功しました")
}
