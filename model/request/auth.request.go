package request

type Login struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

type ClaimsJWT struct {
	Expired   string `json:"expired"`
	Email     string `json:"email"`
	Token     string `json:"token,omitempty"`
	TokenType string `json:"token_type,omitempty"`
}
