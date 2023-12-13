package repository

import "net/http"

type API interface {
	Get(string) (*http.Response, error)
	JSONParse(*http.Response, []string) ([]string, error)
}

type ContentfulAPI struct {
	api API
}

func NewContentfulAPI(api API) ContentfulAPI {
	return ContentfulAPI{
		api: api,
	}
}

func (c ContentfulAPI) FindByID(entryID string) ([]string, error) {
	res, err := c.api.Get(entryID)
	if err != nil {
		return []string{}, err
	}
	defer res.Body.Close()

	vecQuery := []string{
		".sys.id",
		".fields.name",
		".sys.createdAt",
	}
	return c.api.JSONParse(res, vecQuery)
}
