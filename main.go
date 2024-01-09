package main

import (
	"Golang/internal/domain/campaign"
	"Golang/internal/domain/error_handling"

	"fmt"
)

func main() {

	campaignInstance, err := campaign.NewCampaign(
		"asd32s",
		"asdasdasd",
		[]string{"email@example.com"})

	if err != nil {
		error_handling.ValidateStruct(err)
		return 
	}

	fmt.Println("ID:", campaignInstance.ID)
	fmt.Println("Name:", campaignInstance.Name)
	fmt.Println("CreatedAt:", campaignInstance.CreatedAt)
	fmt.Println("Content:", campaignInstance.Content)

	for i, contact := range campaignInstance.Contacts {
		fmt.Printf("Contact %d:\n", i)
		fmt.Println("  Email:", contact.Email)
	}

}
