package campaign

import (
	"Golang/internal/contract"
	"Golang/internal/domain/error_handling"
	"fmt"
	"time"

	"github.com/rs/xid"
)

type Service struct {
	Repository Repository
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
	fmt.Println("Validation Error:", err)

	return nil, err
}

func (service *Service) Create(newCampaign contract.NewCampaign) (string, error) {

	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Contacts)

	if err != nil {
		return "", err
	}

	err = service.Repository.Save(campaign)
	if err != nil {
		return "", err
	}

	return campaign.ID, nil
}

func (service *Service) ListAll() ([]*Campaign, error) {
	return service.Repository.List()
}

func (service *Service) Update (id string, campaignToUpdate contract.UpdateCampaign)(*Campaign, error) {
	existingCampaign, err := service.Repository.FindById(id)

	if err != nil {
		return nil, err
	}

	size := len(campaignToUpdate.Contacts)
	emails := campaignToUpdate.Contacts
	contacts := make([]Contact, size)

	for indice, valor := range emails {
		contacts[indice] = Contact{Email: valor}
	}

	existingCampaign.Name = campaignToUpdate.Name
	existingCampaign.Content = campaignToUpdate.Content
	existingCampaign.Contacts = contacts

	err = error_handling.ValidateStruct(existingCampaign)
	if err != nil {
			fmt.Println("Validation Error:", err)
			return nil, err
	}

	updatedCampaign, err := service.Repository.Update(id, existingCampaign)
	if err != nil {
			return nil, err
	}

	return updatedCampaign, nil
}
