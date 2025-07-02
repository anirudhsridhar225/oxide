package handler

import (
	"encoding/json"
	"net/http"

	"oxide/db"
	"oxide/models"
)

//function to add a user
// @Summary add a new user to the db
// @Description adds a new user to the database
// @Tags user
// @Param user body models.User true "User object"
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string "Invalid request payload"
// @Failure 500 {object} map[string]string "Failed to create user"
// @Router /api/user/add [post]
func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := db.DB.Create(&user).Error; err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := map[string]interface{}{
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
// @Failure 500 {object} map[string]string "Failed to fetch users"
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
// @Description Deletes a user from the database (not implemented yet)
// @Tags user
// @Produce json
// @Success 200 {object} map[string]string "Delete user functionality not implemented yet"
// @Failure 400 {object} map[string]string "Failed to delete user"
// @Router /api/user/delete [delete]
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// This function is a placeholder for future implementation
	// Currently, it does not perform any action
	w.WriteHeader(http.StatusNotImplemented)
	response := map[string]string{
		"message": "Delete user functionality not implemented yet",
	}
	json.NewEncoder(w).Encode(response)
}