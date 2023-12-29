package campaign

import (
	"Golang/internal/contract"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	service := Service{}
	newCampaing := contract.NewCampaign{
		Name:     "test",
		Content:  "test content",
		Contacts: []string{"teste@email.com"},
	}

	id, err := service.Create(newCampaing)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_SaveCampaign(t *testing.T) {
	assert := assert.New(t)

	newCampaing := contract.NewCampaign{
		Name:     "test",
		Content:  "test content",
		Contacts: []string{"teste@email.com"},
	}

	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaing.Name {
			return false
		} else if campaign.Content != newCampaing.Content {
			return false
		} else if len(campaign.Contacts) != len(newCampaing.Contacts) {
			return false
		}

		return true
	})).Return(nil)
	service := Service{repositoryMock}

	id, err := service.Create(newCampaing)

	assert.NotNil(id)
	assert.Nil(err)
}
