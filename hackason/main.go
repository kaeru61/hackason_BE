package main

import (
	"db/controller"
	"db/dao/maindao"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	controller.Handler()

	maindao.CloseDBWithSysCall()

	log.Println("Listening...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func init() {
	maindao.OpenSQL()
}
