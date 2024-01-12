package endpoints

import "Golang/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}