package initdb

import (
	"database/sql"
	"log"
)

// Db 数据库连接
var Db *sql.DB

// 初始化
func init() {
	var err error
	Db, err = sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)
}
