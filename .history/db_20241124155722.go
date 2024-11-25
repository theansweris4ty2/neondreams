package main

import (
	"database/sql"
	"fmt"
)

func dg() {
	db, err := sql.Open("mysql", "database=test1")
}
