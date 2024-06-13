package usercontroller

import (
	"db/application/userapplication"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"unicode/utf8"

	"github.com/oklog/ulid/v2"
)

func UserRegisterController(w http.ResponseWriter, r *http.Request) {
	type info struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var Info info
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if err := json.NewDecoder(r.Body).Decode(&Info); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("decodeerror")
		return
	}
	name := Info.Name
	nameC := utf8.RuneCountInString(name)
	if nameC > 50 || nameC == 0 {
		w.WriteHeader(http.StatusBadRequest)
	}
	age := Info.Age
	if age > 80 || age < 20 {
		w.WriteHeader(http.StatusBadRequest)
	}
	idA := ulid.Make()
	id := idA.String()
	err := userapplication.UserRegisterApplication(id, name, age)
	if err != nil {
		log.Printf("fail: db.Exec, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Printf("{'id': '%s¥n'", id)
	w.WriteHeader(http.StatusOK)
}
