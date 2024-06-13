package maindao

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var Db *sql.DB

func OpenSQL() {
	// ①-1
	// DB接続のための準備
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPwd := os.Getenv("MYSQL_PWD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	connStr := fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)
	Db_, err := sql.Open("mysql", connStr)

	// ①-2
	if err != nil {
		log.Fatalf("fail: sql.Open, %v\n", err)
	}
	// ①-3
	if err := Db_.Ping(); err != nil {
		log.Fatalf("fail: db.Ping, %v\n", err)
	}
	Db = Db_
}

func CloseDBWithSysCall() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		s := <-sig
		log.Printf("received syscall, %v\n", s)

		if err := Db.Close(); err != nil {
			log.Fatal(err)
		}

		log.Printf("success: Db.Close")
		os.Exit(0)
	}()
}
