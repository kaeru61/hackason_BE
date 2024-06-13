package controller

import (
	"fmt"
	"net/http"
	"os"
)

func Handler() {
	version := os.Getenv("VERSION")

	http.HandleFunc(fmt.Sprintf("/v%s/post", version), postController)
}
