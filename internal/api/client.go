package api

import (
	"encoding/json"
	"errors"
	"net/http"
)

type APIClient struct {
	baseURL string
	client  *http.Client
}

func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

func (c *APIClient) FetchData(endpoint string, result interface{}) error {
	resp, err := c.client.Get(c.baseURL + endpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to fetch data from API")
	}

	return json.NewDecoder(resp.Body).Decode(&result)
}
