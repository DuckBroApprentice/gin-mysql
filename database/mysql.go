package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var MySq *sql.DB

const (
	UserName     string = "root"
	Password     string = "123123"
	Addr         string = "127.0.0.1"
	Port         int    = 3306
	Database     string = "demo"
	MaxLifetime  int    = 10
	MaxOpenConns int    = 10
	MaxIdleConns int    = 10
)

var B int = 1

func DBConn() {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", UserName, Password, Addr, Port, Database)
	MySq, err := sql.Open("mysql", conn)
	log.Println("start")
	if err != nil {
		log.Fatal(err)
		log.Println("MySQL connect fail!")
	}
	B = 10
	err = MySq.Ping()
	if err != nil {
		log.Fatal(err)
		log.Println("MySQL connect fail!")
	}
	log.Println("success")

	MySq.SetConnMaxLifetime(time.Duration(MaxLifetime) * time.Second)
	MySq.SetMaxOpenConns(MaxOpenConns)
	MySq.SetMaxIdleConns(MaxIdleConns)
}
