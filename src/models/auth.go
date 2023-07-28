package models

type AuthRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthReponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
