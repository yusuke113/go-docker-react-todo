package main

import (
	"app/db"
	"app/model"
	"fmt"
)

func main() {
	// データベースに接続
	dbConn := db.NewDB()

	// main 関数終了時にデータベースの切断を保証するため、defer ステートメントを使用
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)

	// モデルのマイグレーションを実行
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
