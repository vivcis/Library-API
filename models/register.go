package models

type Register struct {
	Model
	UserID       string `json:"user_id" gorm:"not null"`
	User         User   `json:"user" gorm:"foreignkey:UserID"`
	BookID       string `json:"book_id" gorm:"not null"`
	Book         Book   `json:"book" gorm:"foreignkey:BookID"`
	BorrowedDate int64  `json:"borrowed_date" gorm:"not null"`
	ReturnDate   int64  `json:"return_date"`
}
