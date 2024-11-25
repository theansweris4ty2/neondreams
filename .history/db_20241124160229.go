package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

func dg() {
	db, err := sql.Open("mysql", "database=test1")
}
