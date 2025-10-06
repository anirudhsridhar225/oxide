package handler

import (
	"fmt"
	"net/http"

	"oxide/models"
	_ "oxide/responses"
)

// not implemented yet
func GolangHandler(w http.ResponseWriter, r *http.Response) {
	user := models.User{}

	fmt.Println(user.ID)
}
