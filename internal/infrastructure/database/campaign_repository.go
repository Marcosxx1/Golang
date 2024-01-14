package database

import (
	"Golang/internal/domain/campaign"
	"fmt"
)

type CampaignRepository struct {
	campaigns []*campaign.Campaign
}
 
func (cr *CampaignRepository) Save(campaign *campaign.Campaign) error {
	cr.campaigns = append(cr.campaigns, campaign)
	return nil
}

func (cr *CampaignRepository) List() ([]*campaign.Campaign, error){
	 allCampaigns := cr.campaigns
	 return allCampaigns, nil
}

func (cr *CampaignRepository) Update(id string, updatedCampaign *campaign.Campaign) (*campaign.Campaign, error) {
	var index int = -1

	for campaignIdFound, campaigns := range cr.campaigns {
			if campaigns.ID == id {
					index = campaignIdFound
					break
			}
	}

	if index == -1 {
			return nil, fmt.Errorf("Campaign with ID %s not found", id)
	}

	cr.campaigns[index] = updatedCampaign
	return updatedCampaign, nil
}


func (cr *CampaignRepository) FindById(id string) (*campaign.Campaign, error)  {
	
		for _, campaigns := range cr.campaigns {
			if id == campaigns.ID{
				return campaigns, nil
			}
		}
		return nil, fmt.Errorf("Campaign with ID %s not found", id)

}