package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() {
	fmt.Println("Go MySQL Tutorial")

	db, err := sql.Open("mysql", "root:12345678@(127.0.0.1:3306)/test")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO test VALUES ( 100, 'Sơn' )")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

}
