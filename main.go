package main

import (
	"Golang/internal/domain/campaign"

	"github.com/go-playground/validator/v10"
)

func main() {

	contacts := []campaign.Contact{{Email: ""}}
	campaign := campaign.Campaign{
		Name:     "Campanha Teste",
		Contacts: contacts,
	}
	validate := validator.New()
	err := validate.Struct(campaign)
	if err == nil {
		println("nenhum erro")
	} else {
		validationsErrors := err.(validator.ValidationErrors)
		for _, v := range validationsErrors {
			println(v.StructField() + " is invalid:  " + v.Tag())
		}
	}
}
