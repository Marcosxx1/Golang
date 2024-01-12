package campaign

import (
	"Golang/internal/contract"
)

type Service struct {
	Repository Repository
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

func (service *Service) listAll() ([]*Campaign, error) {
	return service.Repository.List()
}
