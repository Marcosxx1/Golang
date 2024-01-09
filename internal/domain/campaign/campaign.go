package campaign

import (
	"Golang/internal/domain/error_handling"
	"fmt"
	"time"

	"github.com/rs/xid"
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

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
	contacts := make([]Contact, len(emails))

	for indice, valor := range emails {
		contacts[indice] = Contact{Email: valor}
	}

	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedAt: time.Now(),
		Contacts:  contacts,
	}

	err := error_handling.ValidateStruct(campaign)
	if err == nil {
		return campaign, nil
	}
	fmt.Println("Validation Error:", err) // Add this line for debugging

	return nil, err
}
