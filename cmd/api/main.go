package main

import (
	"Golang/internal/domain/campaign"
	"Golang/internal/endpoints"
	"Golang/internal/infrastructure/database"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	campaignService := campaign.Service{
		Repository: &database.CampaignRepository{},
	}
	handler := endpoints.Handler{
		CampaignService: campaignService,
	}

	router.Post("/campaigns", handler.CampaignPost)
	router.Get("/", handler.CampaingGet)
	router.Patch("/campaigns/{campaignId}", handler.CampaignUpdate)

	http.ListenAndServe(":3000", router)
}
