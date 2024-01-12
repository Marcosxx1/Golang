package campaign

type Repository interface {
	Save(campaign *Campaign) error
	List() ([]*Campaign, error)
}
