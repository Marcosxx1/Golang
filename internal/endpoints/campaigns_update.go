package endpoints

import (
	"Golang/internal/domain/campaign"
	internalerrors "Golang/internal/internal-errors"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (h *Handler) CampaignUpdate(res http.ResponseWriter, req *http.Request) {
	var campaignRequest *campaign.Campaign
	render.DecodeJSON(req.Body, &campaignRequest)

	campaignId := chi.URLParam(req, "campaignId")

	campaigns, err := h.CampaignService.Repository.Update(campaignId, campaignRequest)

	if err != nil {
		if errors.Is(err, internalerrors.ErrInternal) {
			render.Status(req, 500)
		} else {
			render.Status(req, 400)
		}
		render.JSON(res, req, map[string]string{"error: ": err.Error()})
		return
	}

	render.Status(req, 201)
	render.Status(req, 201)
	render.JSON(res, req, map[string]interface{}{
		"campaignId":       campaigns.ID,
		"campaignName":     campaigns.Name,
		"campaignContacts": campaigns.Contacts,
	})
}
