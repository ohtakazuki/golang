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

	// トランザクションを開始
	tx := db.Begin()

	// 更新するデータを準備
	book := Book{
		ID:        6,
		Title:     "Updated Book",
		Price:     1500,
		CreatedAt: time.Now(),
	}

	// データを更新
	result := tx.Model(&Book{}).Where("id = ?", book.ID).Updates(book)
	if result.Error != nil {
		fmt.Println("データの更新に失敗しました:", result.Error)
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

	fmt.Println("データの更新が成功しました")
}
