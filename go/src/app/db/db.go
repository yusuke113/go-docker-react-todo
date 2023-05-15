package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDBはデータベース接続を確立するための関数です。
func NewDB() *gorm.DB {
	// 開発環境の場合は環境変数を読み込む
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		}
	}

	// データベースに接続するためのURLを作成
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PW"), os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))

	// DBに接続
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("DB接続成功")
	return db
}

// CloseDBはデータベース接続を切断するための関数です。
func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	err := sqlDB.Close()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("DB切断成功")
}
