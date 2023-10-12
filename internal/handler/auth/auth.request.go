package auth

type registerRequest struct {
	Username string `json:"username" form:"username" binding:"required,min=1,max=255"`
	Password string `json:"password" form:"password" binding:"required,min=1,max=255"`
	Email    string `json:"email" form:"email" binding:"required,email"`
}
