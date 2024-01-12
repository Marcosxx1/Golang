package database

import "Golang/internal/domain/campaign"

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