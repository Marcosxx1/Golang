package campaign

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCampaign(t *testing.T) {

  assert := assert.New(t)

  expectedName := "Nome Campanha"
  expectedContent := "Body"
  expectedContacts := []string{"test1@example.com", "test2@example.com"}

  campaign := NewCampaign(expectedName, expectedContent, expectedContacts)

  assert.Equal(campaign.ID, "1")
  assert.Equal(expectedName, campaign.Name)
  assert.Equal(expectedContent, campaign.Content)

  for _, actual := range campaign.Contacts {
    for _, expected := range expectedContacts {
      if actual.Email == expected {
        fmt.Println(actual.Email)
        fmt.Println(expected)
        } 
    }
  }
  
}




/* package campaign

import "testing"


func TestNewCampaign(t *testing.T){
	expectedName := "New Campaign"
	expectedContent := "Content"
	expectedContacts := []string{"first@email.com", "second@email.com"}


	campapaign := NewCampaign(expectedName, expectedContent, expectedContacts)

	if campapaign.Name != expectedName{
		t.Errorf("got %s, wanted %s", campapaign.Name, expectedName)
	}
} */