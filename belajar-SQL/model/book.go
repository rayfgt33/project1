package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Book struct {
	ID        int       `json:"id" gorm:"primaryKey;type:serial"`
	NameBook  string    `json:"name_book" gorm:"not null;type:varchar(50)"`
	Author    string    `json:"Author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (e Book) Validation() error { // custom validation
	return validation.ValidateStruct(&e,
		validation.Field(&e.NameBook, validation.Required),
		validation.Field(&e.Author, validation.Required))
}
