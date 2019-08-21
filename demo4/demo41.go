package demo4

/*
	_ "github.com/go-sql-driver/mysql" 说明:
	  1. 引入但不使用，目的在于调用包的 init() 初始化方法
*/

import (
	"database/sql"
	"fmt"

	// 调用 init
	_ "github.com/go-sql-driver/mysql"
)

// TestDemo41 数据库操作
func TestDemo41() {
	// 打开一个注册过的数据库驱动, go-sql-driver中注册了mysql这个数据库驱动
	//
	/*
		sql.Open(type, params)
			1. sql.Open()函数用来打开一个注册过的数据库驱动，go-sql-driver中注册了mysql这个数据库驱动
			2. 第二个参数是DSN(Data Source Name)，它是go-sql-driver定义的一些数据库链接和配置信息。
			  它支持如下格式：
				2.1 user@unix(/path/to/socket)/dbname?charset=utf8
				2.2 user:password@tcp(localhost:5555)/dbname?charset=utf8
				2.3 user:password@/dbname
				2.4 user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname
	*/
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/golang?charset=utf8")
	checkErr(err)

	// 插入数据
	add(db)
	addedID := add(db)

	// 更新数据
	update(db, addedID)

	// 查询
	query(db)

	// 删除
	// deleteByID(db, updateID)
}

// 插入
func add(db *sql.DB) int64 {
	// Stmt是一种准备好的状态，和Conn相关联，而且只能应用于一个goroutine中，不能应用于多个goroutine
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)

	// stmt.Exec()函数用来执行stmt准备好的SQL语句
	res, err := stmt.Exec("xiaot", "研发部门", "2019-12-12")
	checkErr(err)

	// 返回由数据库执行插入操作得到的自增ID号
	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	return id
}

// 更新数据
func update(db *sql.DB, id int64) int64 {
	stmt, err := db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err := stmt.Exec("sunxiaot", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	return affect
}

// 查询数据
func query(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string

		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)

		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}
}

// 删除
func deleteByID(db *sql.DB, id int64) int64 {
	stmt, err := db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err := stmt.Exec(id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	return affect
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
