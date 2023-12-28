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
	fmt.Println(campaign.ID)

}
