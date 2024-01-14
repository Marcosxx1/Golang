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

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*Campaign), args.Error(1)
}

func (r *repositoryMock) FindById(id string) (*Campaign, error) {
	args := r.Called(id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Campaign), args.Error(1)
}

func (r *repositoryMock) Update(id string, campaign *Campaign) (*Campaign, error) {
	args := r.Called(id, campaign)
	return args.Get(0).(*Campaign), args.Error(1)
}

func Test_ListCampaigns(t *testing.T) {
	assert := assert.New(t)

	expectedCampaigns := []*Campaign{
		{Name: "Campaign1", Content: "Content1", Contacts: []Contact{{Email: "contact1@test.com"}}},
		{Name: "Campaign2", Content: "Content2", Contacts: []Contact{{Email: "contact2@test.com"}}},
	}

	repositoryMock := new(repositoryMock)
	repositoryMock.On("List").Return(expectedCampaigns, nil)

	service := Service{Repository: repositoryMock}

	campaigns, err := service.ListAll()

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

func TestListCampaigns(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(repositoryMock)

	expectedCampaigns := []*Campaign{
		{Name: "Campaign1", Content: "Content1", Contacts: []Contact{{Email: "contact1@test.com"}}},
		{Name: "Campaign2", Content: "Content2", Contacts: []Contact{{Email: "contact2@test.com"}}},
	}
	repositoryMock.On("List").Return(expectedCampaigns, nil)

	service := Service{Repository: repositoryMock}

	campaigns, err := service.ListAll()

	assert.Nil(err)
	assert.Equal(expectedCampaigns, campaigns)

	repositoryMock.AssertExpectations(t)
}

func TestListCampaignsWithError(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(repositoryMock)

	expectedError := errors.New("error fetching campaigns")
	repositoryMock.On("List").Return(nil, expectedError)

	service := Service{Repository: repositoryMock}

	campaigns, err := service.ListAll()

	assert.NotNil(err)
	assert.Equal(expectedError, err)
	assert.Nil(campaigns)

	repositoryMock.AssertExpectations(t)
}

func TestUpdateCampaign(t *testing.T) {
	assert := assert.New(t)

	// Create a repositoryMock
	repositoryMock := new(repositoryMock)

	// Set up an expectation for the Update method
	updatedCampaign := contract.UpdateCampaign{ Name: "UpdatedCampaign", Content: "UpdatedContent"}
	repositoryMock.On("Update", "1", updatedCampaign).Return(updatedCampaign, nil)

	// Create a service using the repositoryMock
	service := Service{Repository: repositoryMock}

	// Act
	returnedCampaign, err := service.Update("1", updatedCampaign)

	// Assert
	assert.Nil(err)
	assert.Equal(updatedCampaign, returnedCampaign)

	// Verify that the expected method was called on the repositoryMock
	repositoryMock.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	// Create a new instance of the repository mock
	mockRepo := new(repositoryMock)

	// Create an instance of the service or struct that uses the repository
	// Replace 'Service' and 'YourRepository' with the actual names in your code
	service := Service{Repository: mockRepo}

	// Define a sample campaign and ID
	id := "123"
	campaignToUpdate := &Campaign{
		ID:   id,
		Name: "Updated Campaign",
		// other fields...
	}

	// Expectations: Assuming the Update method returns a campaign and no error
	mockRepo.On("Update", id, campaignToUpdate).Return(campaignToUpdate, nil)

	// Call the actual method you want to test
	updatedCampaign, err := service.Update(id, contract.UpdateCampaign{
		Name:    "Updated Campaign",
		Content: "Updated Content",
		// other fields...
	})

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, updatedCampaign)
	assert.Equal(t, campaignToUpdate.Name, updatedCampaign.Name)
	// assert other fields...

	// Optionally, you can assert that the expected method was called
	mockRepo.AssertExpectations(t)
}

// Test for the Update method with an error case
func TestUpdate_ErrorCase(t *testing.T) {
	// Create a new instance of the repository mock
	mockRepo := new(repositoryMock)

	// Create an instance of the service or struct that uses the repository
	// Replace 'Service' and 'YourRepository' with the actual names in your code
	service := Service{Repository: mockRepo}

	// Define a sample campaign and ID
	id := "123"
	campaignToUpdate := &Campaign{
		ID:   id,
		Name: "Updated Campaign",
		// other fields...
	}

	// Expectations: Assuming the Update method returns an error
	expectedError := errors.New("update failed")
	mockRepo.On("Update", id, campaignToUpdate).Return(nil, expectedError)

	// Call the actual method you want to test
	updatedCampaign, err := service.Update(id, contract.UpdateCampaign{
		Name:    "Updated Campaign",
		Content: "Updated Content",
		// other fields...
	})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, updatedCampaign)
	assert.EqualError(t, err, expectedError.Error())

	// Optionally, you can assert that the expected method was called
	mockRepo.AssertExpectations(t)
}