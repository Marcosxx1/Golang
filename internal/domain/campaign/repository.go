package campaign

type Repository interface {
	Save(campagin *Campaign) error
}