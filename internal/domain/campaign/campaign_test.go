package campaign_test

import (
	"testing"

	"Golang/internal/domain/campaign"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

func TestNewCampaignMinimumNameLength(t *testing.T) {
	name := "Min"
	content := "This is a valid campaign."
	emails := []string{"email@example.com"}

	newCampaign, err := campaign.NewCampaign(name, content, emails)

	assert.Error(t, err)
	assert.Nil(t, newCampaign)
}

func TestNewCampaignMaximumNameLength(t *testing.T) {
	name := "VeryLongName123456789000245"
	content := "This is a valid campaign."
	emails := []string{"email@example.com"}

	newCampaign, err := campaign.NewCampaign(name, content, emails)

	assert.Error(t, err)
	assert.Nil(t, newCampaign)
}

func TestNewCampaignMinimumMaximumContentLength(t *testing.T) {
	fake := faker.New()
	name := "Valid Campaign"
	minContent := "Min"
	maxContent := fake.Lorem().Text(1100)
	emails := []string{"email@example.com"}

	// Test for minimum content length
	newCampaign, err := campaign.NewCampaign(name, minContent, emails)
	assert.Error(t, err)
	assert.Nil(t, newCampaign)

	// Test for maximum content length
	newCampaign, err = campaign.NewCampaign(name, maxContent, emails)
	assert.Error(t, err)
	assert.Nil(t, newCampaign)
}

func TestNewCampaignMinimumMaximumEmails(t *testing.T) {
	name := "Valid Campaign"
	content := "This is a valid campaign."
	minEmails := []string{"email@example.com"}
	maxEmails := make([]string, 26) // 26 emails to exceed the maximum allowed

	// Test for minimum number of emails
	newCampaign, err := campaign.NewCampaign(name, content, minEmails)
	assert.NoError(t, err)
	assert.NotNil(t, newCampaign)

	// Test for maximum number of emails
	newCampaign, err = campaign.NewCampaign(name, content, maxEmails)
	assert.Error(t, err)
	assert.Nil(t, newCampaign)
}

func TestNewCampaignValid(t *testing.T) {
	name := "Valid Campaign"
	content := "This is a valid campaign."
	emails := []string{"email@example.com"}

	newCampaign, err := campaign.NewCampaign(name, content, emails)

	assert.NoError(t, err)
	assert.NotNil(t, newCampaign)
}

func TestNewCampaignInvalidName(t *testing.T) {
	name := "Shrt"
	content := "This is a valid campaign."
	emails := []string{"email@example.com"}

	newCampaign, err := campaign.NewCampaign(name, content, emails)

	assert.Error(t, err)
	assert.Equal(t,"name is required with min 5",  err.Error())
	assert.Nil(t, newCampaign)
}

func TestNewCampaignInvalidContent(t *testing.T) {
	name := "Valid Campaign"
	content := "Shrt"
	emails := []string{"email@example.com"}

	newCampaign, err := campaign.NewCampaign(name, content, emails)

	assert.Error(t, err)
	assert.Equal(t,"content is required with min 5",  err.Error())
	assert.Nil(t, newCampaign)
}

func TestNewCampaignInvalidEmail(t *testing.T) {
	name := "Valid Campaign"
	content := "This is a valid campaign."
	emails := []string{"invalid-email"}

	newCampaign, err := campaign.NewCampaign(name, content, emails)

	assert.Error(t, err)
	assert.Nil(t, newCampaign)
} 

func TestNewCampaignEmptyEmails(t *testing.T) {
	name := "Valid Campaign"
	content := "This is a valid campaign."
	emails := []string{}

	newCampaign, err := campaign.NewCampaign(name, content, emails)

	assert.Error(t, err)
	assert.Nil(t, newCampaign)
}
