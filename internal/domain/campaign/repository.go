package campaign

type Repository interface {
	Save(campaign *Campaign) error
	List() ([]*Campaign, error)
	Update(id string, campaign *Campaign) (*Campaign, error)
	FindById(id string) (*Campaign, error)
}
