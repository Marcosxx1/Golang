package campaign

import (
	"Golang/internal/contract"
	internalerrors "Golang/internal/internal-errors"
	"errors"
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
	
func (r *repositoryMock) List() ([]*Campaign, error) {
	args := r.Called()
	return args.Get(0).([]*Campaign), args.Error(1)
}

func Test_ListCampaigns(t *testing.T){
	assert := assert.New(t)

	expectedCampaigns := []*Campaign{
		{Name: "Campaign1", Content: "Content1", Contacts: []Contact{{Email: "contact1@test.com"}}},
		{Name: "Campaign2", Content: "Content2", Contacts: []Contact{{Email: "contact2@test.com"}}},
	}

	repositoryMock := new(repositoryMock)
	repositoryMock.On("List").Return(expectedCampaigns, nil)

	service := Service{Repository: repositoryMock}

	campaigns, err := service.listAll()

	assert.Nil(err)

	assert.Equal(expectedCampaigns, campaigns)

	repositoryMock.AssertExpectations(t)
}

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(nil)

	service := Service{Repository: repositoryMock}

	id, err := service.Create(contract.NewCampaign{
		Name:     "Test Y",
		Content:  "Body Hi!",
		Contacts: []string{"teste1@test.com"},
	})

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(repositoryMock)
	service := Service{Repository: repositoryMock}

	newCampaign := contract.NewCampaign{
		Name:     "",
		Content:  "Body Hi!",
		Contacts: []string{"teste1@test.com"},
	}

	_, err := service.Create(newCampaign)

	assert.NotNil(err)
	assert.Equal("name is required with min 5", err.Error())
}

func Test_Create_SaveCampaign(t *testing.T) {
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		return campaign.Name == "Test Y" &&
			campaign.Content == "Body Hi!" &&
			len(campaign.Contacts) == 1 &&
			campaign.Contacts[0].Email == "teste1@test.com"

	})).Return(nil)

	service := Service{Repository: repositoryMock}
	service.Create(contract.NewCampaign{
		Name:     "Test Y",
		Content:  "Body Hi!",
		Contacts: []string{"teste1@test.com"},
	})

	repositoryMock.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(repositoryMock)

	repositoryMock.On("Save", mock.Anything).Return(internalerrors.ProcessErrorToReturn(errors.New("error to save on database")))

	service := Service{Repository: repositoryMock}

	_, err := service.Create(contract.NewCampaign{
		Name:     "Test Y",
		Content:  "Body Hi!",
		Contacts: []string{"teste1@test.com"},
	})
	

	assert.True(errors.Is(internalerrors.ErrInternal, err))
}
