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
	if err != nil {
		fmt.Println(err)
		return
	}
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
	// Create a new conn
	fmt.Println("conn2 created time:", time.Now())
	conn2, err := DB.Conn(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = conn.Close()
	_ = conn2.Close()

	time.Sleep(30 * time.Second)

}
