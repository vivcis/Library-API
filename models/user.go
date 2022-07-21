package models

type User struct {
	Model
	UserName     string `json:"user_name" gorm:"not null"`
	FirstName    string `json:"first_name" gorm:"not null" binding:"required"`
	LastName     string `json:"last_name" gorm:"nut null"`
	Email        string `json:"email" gorm:"unique; not null" binding:"required,email"`
	Password     string `json:"password, omitempty" gorm:"-"`
	PasswordHash string `json:"password_hash" gorm:"password_hash"`
}
