package presenter


type UserResponse struct {
	Email string `json:"email"`
	Username string `json:"username"`
	Name string `json:"name"`
}
type UserSuccessResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
	User UserResponse `json:"user"`
}