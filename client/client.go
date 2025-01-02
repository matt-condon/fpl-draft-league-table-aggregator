package client

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/matt-condon/fpl-draft-league-table-aggregator/core/models"
)

type Client struct {
	httpClient *http.Client
	UserAgent  string
}

func (c *Client) NewRequest(method string, path string, body []byte) (*http.Request, error) {

	req, err := http.NewRequest(method, path, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, models.HttpError{
			Status:  resp.StatusCode,
			Message: resp.Status,
		}
	}

	if resp.ContentLength == 0 {
		return resp, nil
	}

	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	c := &Client{
		httpClient: httpClient,
	}

	return c
}
