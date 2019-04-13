package serializers

type UserResponse struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
}
