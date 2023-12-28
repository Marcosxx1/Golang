package campaign

import (
	"errors"
	"regexp"
	"strconv"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string
}

type Campaign struct {
	ID        string
	Name      string
	CreatedAt time.Time
	Content   string
	Contacts  []Contact
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
	if name == "" {
		return nil, errors.New("Name must be filled")
	}
	if content == "" {
		return nil, errors.New("Content must be filled")
	}

	for _, email := range emails {
		if !isValidEmail(email) {
			return nil, errors.New("Invalid email: " + email)
		}
	}

	contacts := make([]Contact, len(emails))

	for indice, valor := range emails {
		contacts[indice].Email = valor
	}

	return &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedAt: time.Now(),
		Contacts:  contacts,
	}, nil
}

func isValidEmail(email string) bool {
	if email == "" {
		return false
	}

	_, err := strconv.Atoi(email)
	if err == nil {
		return false
	}

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, email)
	return match
}
