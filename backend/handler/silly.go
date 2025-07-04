package handler

import (
	"net/http"

	_ "oxide/responses"
)

// TeapotHandler handles your teapot
// @Summary Teapot
// @Description Determines if you are worthy of chai
// @Tags tea
// @Produce json
// @Failure 418 {object} responses.ErrorResponse
// @Router /teapot [get]
func TeapotHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "no chai for you fn", http.StatusTeapot)
	return
}
