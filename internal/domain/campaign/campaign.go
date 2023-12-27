package campaign

import "time"

type Contact struct{
	Email string
}

type Campaign struct {
	ID string
	Name string
	CreatedAt time.Time
	Content string
	Contacts []Contact
}

func NewCampaign(name string, content string, emails []string) *Campaign {
 
	contacts := make([]Contact, len(emails))

	for indice, valor:= range emails{
		contacts[indice].Email = valor
	}

	return &Campaign{
		ID: "1",
		Name: name,
		Content: content,
		CreatedAt: time.Now(),
		Contacts: contacts ,
	}
}