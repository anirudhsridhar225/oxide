package handler

import (
	"fmt"
	"net/http"
)

func CloneHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CloneHandler called")
}
