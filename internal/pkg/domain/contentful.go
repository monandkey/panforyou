package domain

type ContentfulRepository interface {
	FindByID(string) ([]string, error)
}
