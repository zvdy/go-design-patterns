package main

import (
	"database/sql"
	"fmt"
	"sync"
	//_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Connection *sql.DB
}

var instance *Database
var once sync.Once

func GetDatabaseInstance() *Database {
	once.Do(func() {
		db, err := sql.Open("mysql", "user:password@/dbname")
		if err != nil {
			fmt.Println("Failed to create the handle")
		}
		if err2 := db.Ping(); err2 != nil {
			fmt.Println("Failed to keep connection alive")
		}
		instance = &Database{Connection: db}
	})
	return instance
}

func main() {
	db1 := GetDatabaseInstance()
	fmt.Println(db1.Connection)

	db2 := GetDatabaseInstance()
	fmt.Println(db2.Connection)

	if db1 == db2 {
		fmt.Println("db1 and db2 point to the same instance.")
	}
}
