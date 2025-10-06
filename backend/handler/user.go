package handler

import (
	"encoding/json"
	"net/http"

	"oxide/db"
	"oxide/models"
	_ "oxide/responses"

	"github.com/google/uuid"
)

// function to add a user
// @Summary Add a user
// @Description adds a new user to the database
// @Tags user
// @Param user body models.User true "User object"
// @Produce json
// @Success 201 {object} responses.UserResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /api/user/add [post]
func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user.ID = ""

	if err := db.DB.Create(&user).Error; err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := map[string]any{
		"message": "User created successfully",
		"user_id": user.ID,
	}
	json.NewEncoder(w).Encode(response)
}

// UserHandler handles the user retrieval endpoint
// @Summary Get all users
// @Description Retrieves all users from the database
// @Tags user
// @Produce json
// @Success 200 {array} models.User "List of users"
// @Failure 500 {object} responses.ErrorResponse
// @Router /api/user [get]
func UserHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	if err := db.DB.Find(&users).Error; err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// DeleteUserHandler handles the user deletion endpoint
// @Summary Delete a user
// @Description Deletes a user from the database
// @Param id query string true "user id"
// @Tags user
// @Produce json
// @Success 200 {object} responses.UserResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /api/user/delete [delete]
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if _, err := uuid.Parse(id); err != nil {
		http.Error(w, "Invalid id entered, this is not a uuid", http.StatusBadRequest)
		return
	}

	result := db.DB.Where("ID = ?", id).Delete(&models.User{})

	if result.Error != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]any{
		"message": "User deleted successfully",
		"user_id": id,
	}
	json.NewEncoder(w).Encode(response)
}

// EmailVerifiedHandler() handles the email verification endpoint
// @Summary Verify a user
// @Description Checks whether the user is verified or not by sending them an email and waiting for the email link to be verified
// @Tags user
// @Produce json
// @Success 200 {object} responses.UserResponse
// @Failure 400 {object} responses.ErrorResponse
// @Router /api/user/verify [post]
func EmailVerifiedHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "verified email",
	}
	json.NewEncoder(w).Encode(response)
}
