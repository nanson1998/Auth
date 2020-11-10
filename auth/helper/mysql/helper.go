package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", "root", "12345678", "127.0.0.1:3306", "test")
	dsn += "?collation=utf8mb4_bin&clientFoundRows=false&interpolateParams=true&maxAllowedPacket=0&multiStatements=false&parseTime=true&timeout=5000ms"
	dsn += "&wait_timeout=15"

	fmt.Println("start time:", time.Now())
	DB, err := sql.Open("mysql", dsn)
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(100)
	DB.SetConnMaxLifetime(time.Second * 3)
	ErrorCheck(err)
	defer DB.Close()

	var name string
	var value string
	if err := DB.QueryRow("show session variables where variable_name='wait_timeout'").Scan(&name, &value); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name, value)

	time.Sleep(4 * time.Second)
	// Get the existing conn
	conn, err := DB.Conn(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	insert, err := DB.Prepare("INSERT INTO test VALUES ( 2, 'Nam' )")
	ErrorCheck(err)
	defer insert.Close()

	_ = conn.Close()

	time.Sleep(3 * time.Second)

}
func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}
