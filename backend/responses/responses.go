package responses

type UserResponse struct {
	Message string `json:"message"`
	UserID  string `json:"userID"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
