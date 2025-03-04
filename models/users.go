package models

type (
	Users struct {
		Gorm
		Username string `gorm:"not null;uniqueIndex" json:"username"`
		Email    string `gorm:"not null;uniqueIndex" json:"email"`
		Password string `gorm:"not null" json:"password,omitempty"`
	}

	RegisterRequest struct {
		Username string `json:"username" form:"username" valid:"required~Your username is required"`
		Email    string `json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
		Password string `json:"password" form:"password" valid:"required~Your password is required,minstringlength(8)~Password has to have a minimum length of 8 characters"`
	}

	LoginRequest struct {
		Email    string `json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
		Password string `json:"password" form:"password" valid:"required~Your password is required,minstringlength(8)~Password has to have a minimum length of 8 characters"`
	}
)
