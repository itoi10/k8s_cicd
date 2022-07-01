/*
データベースの初期化処理を定義
*/
package model

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql"
)

var db *gorm.DB

func init() {
	// 接続情報
	DBMS := "mysql"
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	HOST := os.Getenv("MYSQL_HOST")
	DBNAME := os.Getenv("MYSQL_DATABASE")

	CONNECT := USER + ":" + PASS + "@tcp(" + HOST + ":3306)/" + DBNAME

	// DB接続
	var err error
	db, err = gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(fmt.Sprintf("failed to connect database %v", err))
	}
	// マイグレート
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Todo{})
}
