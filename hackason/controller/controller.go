package controller

import (
	"fmt"
	"net/http"
)

func Handler() {
	version := 1

	http.HandleFunc(fmt.Sprintf("/v%s/post", version), postController)
	http.HandleFunc(fmt.Sprintf("/v%s/user", version), userController)
}
