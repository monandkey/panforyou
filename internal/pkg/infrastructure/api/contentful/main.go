package contentful

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/itchyny/gojq"
	"github.com/monandkey/panforyou/internal/pkg/adapter/repository"
)

type contentfulAPI struct {
}

func NewContentfulAPIFactory() repository.API {
	return contentfulAPI{}
}

func (a contentfulAPI) urlFactory(entryID string) (string, error) {
	endpint := os.Getenv("CONTENTFUL_ENDPINT")
	spaces := os.Getenv("CONTENTFUL_SPACES")
	accessToken := os.Getenv("CONTENTFUL_ACCESS_TOKEN")

	u, err := url.Parse(endpint)
	if err != nil {
		return "", err
	}
	u.Path = path.Join(
		u.Path,
		"spaces",
		spaces,
		"entries",
		entryID,
	)
	q := u.Query()
	q.Set("access_token", accessToken)
	u.RawQuery = q.Encode()
	return u.String(), nil
}

func (a contentfulAPI) Get(entryID string) (*http.Response, error) {
	url, err := a.urlFactory(entryID)
	if err != nil {
		return nil, err
	}
	return http.Get(url)
}

func (a contentfulAPI) JSONParse(res *http.Response, jqPattern []string) ([]string, error) {
	var (
		vecStr     []string
		decodeData interface{}
	)

	if err := json.NewDecoder(res.Body).Decode(&decodeData); err != nil {
		return vecStr, err
	}

	for _, v := range jqPattern {
		query, err := gojq.Parse(v)
		if err != nil {
			return vecStr, err
		}
		iter := query.Run(decodeData)
		for {
			w, ok := iter.Next()
			if !ok {
				break
			}
			if err, ok := w.(error); ok {
				return vecStr, err
			}
			if w == nil {
				return vecStr, errors.New("the specified ID may not exist or the structure of Response may have changed")
			}
			vecStr = append(vecStr, w.(string))
		}
	}
	return vecStr, nil
}
