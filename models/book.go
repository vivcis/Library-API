package models

type Book struct {
	Model
	Title    string `json:"title" gorm:"not null" binding:"required"`
	Author   string `json:"author" gorm:"not null" binding:"required"`
	Copies   uint   `json:"copies" gorm:"not null" binding:"required"`
	ISBN     string `json:"ISBN" gorm:"unique;not null" binding:"required"`
	Returned bool   `json:"returned"`
}
