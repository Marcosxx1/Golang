package main

import (
	internalerrors "Golang/internal"
	"Golang/internal/contract"
	"Golang/internal/domain/campaign"
	"Golang/internal/infrastructure/database"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	service := campaign.Service{
		Repository: &database.CampaignRepository{},
	}

	router.Post("/campaigns", func(res http.ResponseWriter, req *http.Request) {
		var request contract.NewCampaign
		render.DecodeJSON(req.Body, &request)

		id, err := service.Create(request)

		if err != nil {

			if errors.Is(err, internalerrors.ErrInternal) {
				render.Status(req, 500)
			} else {
				render.Status(req, 400)
			}
			render.JSON(res, req, map[string]string{"error": err.Error()})
			return
		}
		render.Status(req, 202)
		render.JSON(res, req, map[string]string{"id": id})
	})

	http.ListenAndServe(":3000", router)
}
