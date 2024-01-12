package endpoints

import (
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CampaingGet (res http.ResponseWriter, req *http.Request) {
	campaigns, err := h.CampaignService.Repository.List()

	if err != nil {
		render.Status(req, 500)
		render.JSON(res, req, map[string]string{"error": err.Error()})
		return
	}
	render.Status(req, 200)
	render.JSON(res, req, campaigns)
}