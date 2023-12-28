package campaign

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	expectedName     = "Nome Campanha"
	expectedContent  = "Body"
	expectedContacts = []string{"test1@example.com", "test2@example.com"}
)

func Test_New_Campaign_Creation(t *testing.T) {

	// arrange
	assert := assert.New(t)
	campaign, _ := NewCampaign(expectedName, expectedContent, expectedContacts)

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
	campaign, _ := NewCampaign(expectedName, expectedContent, expectedContacts)

	// assert
	assert.NotNil(campaign.ID)
}

func Test_New_Campaign_CreatedAt_Must_Be_Now(t *testing.T) {

	// arrange
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	// act
	campaign, _ := NewCampaign(expectedName, expectedContent, expectedContacts)

	// assert
	assert.Greater(campaign.CreatedAt, now)
}

func Test_NewCampaign_MustValidateName(t *testing.T){
	assert := assert.New(t)

	_, err := NewCampaign("", expectedContent, expectedContacts)

	assert.Equal("Name must be filled", err.Error())
}

func Test_NewCampaign_MustValidateContent(t *testing.T){
	assert := assert.New(t)

	_, err := NewCampaign(expectedName, "", expectedContacts)

	assert.Equal("Content must be filled", err.Error())
}

 

func TestNewCampaignMustValidateContacts(t *testing.T){
	assert := assert.New(t)

	expectedName := "Test Campaign"
	expectedContent := "Test Content"

	emptyEmail := []string{"", "test2@example.com"}
	_, errEmpty := NewCampaign(expectedName, expectedContent, emptyEmail)
	assert.Error(errEmpty)
	assert.Contains(errEmpty.Error(), "Invalid email: ")

	onlyNumbersEmail := []string{"23d23d", "test2@example.com"}
	_, errNumbersEmail := NewCampaign(expectedName, expectedContent, onlyNumbersEmail)
	assert.Error(errNumbersEmail)
	assert.Contains(errNumbersEmail.Error(), "Invalid email: ")

		
	validEmail := []string{"valid@email.com", "test2@example.com"}
	_, validEmailCreated := NewCampaign(expectedName, expectedContent, validEmail)
	assert.NoError(validEmailCreated)
}