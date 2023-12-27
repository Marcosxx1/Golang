package main

import (
	"Golang/internal/domain/campaign"
	"fmt"
)

func main() {

	expectedName := "Nome Campanha"
	expectedContent := "Body"
	expectedContacts := []string{"test1@example.com", "test2@example.com"}

	campaign := campaign.NewCampaign(expectedName, expectedContent, expectedContacts)

	for _, actual := range campaign.Contacts {
		for _, expected := range expectedContacts {
			if actual.Email == campaign {
				fmt.Println(actual.Email)
				fmt.Println(expected)
			}
		}
	}
}
