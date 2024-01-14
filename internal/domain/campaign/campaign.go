package campaign

import (
	"time"
)

type Contact struct {
	Email string `validate:"required,email"`
}

type Campaign struct {
	ID        string    `validate:"required"`
	Name      string    `validate:"min=5,max=25"`
	CreatedAt time.Time `validate:"required"`
	Content   string    `validate:"min=5,max=1024"`
	Contacts  []Contact `validate:"min=1,max=25,dive"`
}

