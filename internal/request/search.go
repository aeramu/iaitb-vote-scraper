package request

import (
	"encoding/json"
	"errors"
	"github.com/aeramu/ia-itb-scraper/internal/entity"
	"io/ioutil"
	"net/http"
)

const (
	url = "https://ivoting.iaitb.or.id/api/open/alumnee/simple/all"

	apiKey = "api-key"
	apiSecret = "bsgcyfgveyujeygfefc387r34ybr39brnr3r3"

	queryPage = "page"
	queryPerPage = "perPage"
	queryName = "name"
	queryMajor = "studyprogram"
	queryGeneration = "generation"
)

func Search(name, major, generation string) (*entity.Alumnee, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.New("failed create request")
	}

	req.Header.Add(apiKey, apiSecret)

	query := req.URL.Query()
	query.Add(queryPage, "1")
	query.Add(queryPerPage, "10")
	query.Add(queryName, name)
	query.Add(queryMajor, major)
	query.Add(queryGeneration, generation)
	req.URL.RawQuery = query.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.New("failed send request to server")
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("failed read response body")
	}

	var resp entity.Response
	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}

	if resp.Meta.TotalItems < 1 {
		return &entity.Alumnee{
			KeywordName: name,
			Generation:  generation,
			Major:       major,
			Error:       "not found",
		}, nil
	}

	if resp.Meta.TotalItems > 1 {
		resp.Data[0].Error = "multiple result"
	}

	resp.Data[0].KeywordName = name
	resp.Data[0].Major = major
	resp.Data[0].Generation = generation

	return &resp.Data[0], nil
}

