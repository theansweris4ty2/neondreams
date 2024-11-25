package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func dbOpen() {
	db, err := sql.Open("mysql", "database=test1")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
}
