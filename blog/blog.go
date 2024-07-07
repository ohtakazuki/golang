package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/rs/cors"
)

// Post 構造体を定義
type Post struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Likes     int       `json:"likes"`
	CreatedAt time.Time `json:"created_at"`
}

var db *gorm.DB

func main() {
	// データベースに接続
	dsn := "root:_MySQLPassword1@tcp(localhost:3306)/blogsite?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("データベースへの接続に失敗しました:", err)
	}
	// マイグレーションを実行
	db.AutoMigrate(&Post{})

	// ルートハンドラの登録
	http.HandleFunc("/posts", handlePosts)
	http.HandleFunc("/posts/", handlePost)

	// CORSの設定
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://127.0.0.1:5500"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})

	// ハンドラにCORSミドルウェアを適用
	handler := c.Handler(http.DefaultServeMux)

	// サーバーの起動
	log.Fatal(http.ListenAndServe(":8080", handler))
}

// handlePosts は /posts エンドポイントのリクエストを処理
func handlePosts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// 全ての記事を取得
		getPosts(w, r)
	case http.MethodPost:
		// 新しい記事を作成
		createPost(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handlePost は /posts/:id エンドポイントのリクエストを処理
func handlePost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/posts/"):])
	if err != nil {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case http.MethodGet:
		// 特定の記事を取得
		getPost(w, r, id)
	case http.MethodPut:
		// 特定の記事を更新
		updatePost(w, r, id)
	case http.MethodDelete:
		// 特定の記事を削除
		deletePost(w, r, id)
	case http.MethodPost:
		// 特定の記事にいいねを追加
		likePost(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

// getPosts は全ての記事を取得
func getPosts(w http.ResponseWriter, r *http.Request) {
	var posts []Post
	result := db.Find(&posts)
	if result.Error != nil {
		http.Error(w, "クエリの実行に失敗しました", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(posts)
}

// createPost は新しい記事を作成
func createPost(w http.ResponseWriter, r *http.Request) {
	var post Post
	json.NewDecoder(r.Body).Decode(&post)
	post.CreatedAt = time.Now()
	result := db.Create(&post)
	if result.Error != nil {
		http.Error(w, "記事の作成に失敗しました", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(post)
}

// getPost は特定の記事を取得
func getPost(w http.ResponseWriter, r *http.Request, id int) {
	var post Post
	result := db.First(&post, id)
	if result.Error != nil {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(post)
}

// updatePost は特定の記事を更新
func updatePost(w http.ResponseWriter, r *http.Request, id int) {
	var post Post
	result := db.First(&post, id)
	if result.Error != nil {
		http.NotFound(w, r)
		return
	}
	var updatedPost Post
	json.NewDecoder(r.Body).Decode(&updatedPost)
	post.Title = updatedPost.Title
	post.Content = updatedPost.Content

	result = db.Save(&post)
	if result.Error != nil {
		http.Error(w, "記事の更新に失敗しました", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(post)
}

// deletePost は特定の記事を削除
func deletePost(w http.ResponseWriter, r *http.Request, id int) {
	var post Post
	result := db.First(&post, id)
	if result.Error != nil {
		http.NotFound(w, r)
		return
	}
	result = db.Delete(&post)
	if result.Error != nil {
		http.Error(w, "記事の削除に失敗しました", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// likePost は特定の記事にいいねを追加
func likePost(w http.ResponseWriter, r *http.Request, id int) {
	var post Post
	result := db.First(&post, id)
	if result.Error != nil {
		http.NotFound(w, r)
		return
	}
	post.Likes++

	result = db.Save(&post)
	if result.Error != nil {
		http.Error(w, "いいねの追加に失敗しました", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(post)
}
