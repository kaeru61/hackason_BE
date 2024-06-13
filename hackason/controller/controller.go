package controller

import (
	"fmt"
	"net/http"
)

func Handler() {

	http.HandleFunc(fmt.Sprintf("/post"), postController)
	http.HandleFunc(fmt.Sprintf("/user"), userController)
}
