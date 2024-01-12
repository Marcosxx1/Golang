package endpoints

import (
	"Golang/internal/contract"
	internalerrors "Golang/internal/internal-errors"
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignPost (res http.ResponseWriter, req *http.Request) {
	var request  contract.NewCampaign
	render.DecodeJSON(req.Body, &request)
	id, err := h.CampaignService.Create(request)

	if err != nil {
		if errors.Is(err , internalerrors.ErrInternal) {
			render.Status(req, 500)
		} else {
			render.Status(req, 400)
		}
		render.JSON(res, req, map[string]string{"error: ": err.Error()})
		return
	}

	render.Status(req, 201)
	render.JSON(res, req, map[string]string{"id": id})
}