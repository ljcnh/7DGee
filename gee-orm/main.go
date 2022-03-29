/**
 * @Author: lj
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/03/24 22:29
 */

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

var (
	user1 = &User{"Tom", 18}
	user2 = &User{"Sam", 25}
	user3 = &User{"Jack", 25}
)

func main() {
	db, err := sql.Open("sqlite3", "gee.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() { _ = db.Close() }()
	_, _ = db.Exec("CREATE TABLE IF NOT EXISTS User(`Name` text);")

	tx, _ := db.Begin()
	_, err1 := tx.Exec("INSERT INTO User(`Name`) VALUES (?)", "Tom")
	_, err2 := tx.Exec("INSERT INTO User(`Name`) VALUES (?)", "Jack")
	if err1 != nil || err2 != nil {
		_ = tx.Rollback()
		log.Println("Rollback", err1, err2)
	} else {
		_ = tx.Commit()
		log.Println("Commit")
	}
}
