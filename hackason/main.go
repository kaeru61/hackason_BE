package main

import (
	"db/controller"
	"db/controller/usercontroller"
	"db/dao/maindao"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		usercontroller.UserSearchController(w, r)
		return
	case http.MethodPost:
		usercontroller.UserRegisterController(w, r)
		return
	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func main() {
	controller.Handler()
	http.HandleFunc("/user", handler)

	maindao.CloseDBWithSysCall()

	log.Println("Listening...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func init() {
	maindao.OpenSQL()
}
