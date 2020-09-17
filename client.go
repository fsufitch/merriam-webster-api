package mwapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/fsufitch/merriam-webster-api/types"
)

// Client describes access to the Merriam-Webster dictionary at dictionaryapi.com
type Client interface {
	SearchCollegiate(word string) ([]types.CollegiateResult, []string, error)
}

// basicClient is a basic HTTP-based client for querying the M-W dictionary
type basicClient struct {
	APIKey    string
	BaseURLs  *BaseURLs
	UserAgent string
	Client    *http.Client
}

// BaseURLs is a configuration struct for passing in custom base URLs
type BaseURLs struct {
	Collegiate string
}

func (u *BaseURLs) update(other *BaseURLs) *BaseURLs {
	copy := *u
	if other != nil {
		if other.Collegiate != "" {
			copy.Collegiate = other.Collegiate
		}

		// more fields
	}
	return &copy
}

var defaultURLs = &BaseURLs{
	Collegiate: "https://www.dictionaryapi.com/api/v3/references/collegiate/json",
}

// NewClient creates a client based on a given API key
func NewClient(apiKey string, userAgent string, baseURLs *BaseURLs) Client {
	return &basicClient{
		APIKey:    apiKey,
		BaseURLs:  defaultURLs.update(baseURLs),
		UserAgent: userAgent,
		Client:    http.DefaultClient,
	}
}

// SearchCollegiate implements a search of the collegiate dictionary
func (bc basicClient) SearchCollegiate(word string) ([]types.CollegiateResult, []string, error) {
	word = strings.TrimSpace(strings.ToLower(word))

	queryURL, err := url.Parse(bc.BaseURLs.Collegiate)
	if err != nil {
		return nil, nil, err
	}
	queryURL.Path = path.Join(queryURL.Path, word)

	q, _ := url.ParseQuery(queryURL.RawQuery)
	q.Add("key", bc.APIKey)

	queryURL.RawQuery = q.Encode()

	response, err := bc.Client.Do(&http.Request{
		Method: "GET",
		URL:    queryURL,
		Header: http.Header{
			"User-Agent": {bc.UserAgent},
		},
	})

	if err != nil {
		return nil, nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("Non-zero status %d; body: %s", response.StatusCode, string(body))
	}

	result := []types.CollegiateResult{}
	suggestions := []string{}

	var err1, err2 error
	if err1 = json.Unmarshal(body, &result); err1 == nil {
		return result, nil, nil
	} else if err2 = json.Unmarshal(body, &suggestions); err2 == nil {
		return nil, suggestions, nil
	}

	return nil, nil, fmt.Errorf("could not parse response as results or suggestions; response was: %s", string(body))
}
