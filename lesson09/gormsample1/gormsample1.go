package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Book struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	Price     int
	CreatedAt time.Time
}

func main() {
	// データベースに接続
	dsn := "root:_MySQLPassword1@tcp(localhost:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// テーブルを自動的に作成
	err = db.AutoMigrate(&Book{})
	if err != nil {
		panic("failed to migrate database")
	}

	// レコードを挿入
	books := []Book{
		{Title: "はじめてのMySQL", Price: 2980},
		{Title: "はじめてのGo", Price: 1980},
		{Title: "はじめてのHTML", Price: 1000},
		{Title: "はじめてのCSS", Price: 1000},
	}
	result := db.Create(&books)
	if result.Error != nil {
		panic("failed to create books")
	}
	fmt.Printf("Inserted %d records.\n", result.RowsAffected)
}
