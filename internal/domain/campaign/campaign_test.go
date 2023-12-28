package campaign

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func InitializeNewCampaign() (string, string, []string, *Campaign) {
	expectedName := "Nome Campanha"
	expectedContent := "Body"
	expectedContacts := []string{"test1@example.com", "test2@example.com"}

	// act
	campaign := NewCampaign(expectedName, expectedContent, expectedContacts)
	return expectedName, expectedContent, expectedContacts, campaign
}

func Test_New_Campaign_Creation(t *testing.T) {

	// arrange
	assert := assert.New(t)
	expectedName, expectedContent, expectedContacts, campaign := InitializeNewCampaign()

	// assert
	assert.Equal(expectedName, campaign.Name)
	assert.Equal(expectedContent, campaign.Content)
	assert.Equal(len(expectedContacts), len(campaign.Contacts))

	fmt.Println(len(expectedContacts), len(campaign.Contacts))
	for _, actual := range campaign.Contacts {
		for _, expected := range expectedContacts {
			if actual.Email == expected {
				fmt.Println(actual.Email)
				fmt.Println(expected)
			}
		}
	}

}

func Test_New_Campaign_Id_Not_Nil(t *testing.T) {

	// arrange
	assert := assert.New(t)
	_, _, _, campaign := InitializeNewCampaign()

	// assert
	assert.NotNil(campaign.ID)
}

func Test_New_Campaign_CreatedAt_Must_Be_Now(t *testing.T) {

	// arrange
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	// act
	_, _, _, campaign := InitializeNewCampaign()

	// assert
	assert.Greater(campaign.CreatedAt, now)
}
