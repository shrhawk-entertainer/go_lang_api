package forms

type UserSignupValidation struct {
	Name     string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Gender   string `json:"gender"`
}
