package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type EmailCheckerClient struct {
	BaseURL *url.URL

	httpClient *http.Client
}

type EmailCheckRequest struct {
	Email string `json:"email"`
}

type EmailCheckResponse struct {
	IsValid bool `json:"is_valid"`
}

func NewEmailCheckerClient(baseURL *url.URL) *EmailCheckerClient {
	return &EmailCheckerClient{
		BaseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

func (c *EmailCheckerClient) CheckEmail(email string) (bool, error) {
	request := EmailCheckRequest{email}
	req, err := c.newRequest("GET", "/check", request)
	if err != nil {
		return false, err
	}
	var response EmailCheckResponse
	_, err = c.do(req, &response)
	if err != nil {
		return false, err
	}
	return response.IsValid, nil
}

func (c *EmailCheckerClient) newRequest(method, path string, body interface{}) (*http.Request, error) {
	if c.BaseURL == nil {
		return nil, errors.New("Email checker BaseURL is nil")
	}
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}
func (c *EmailCheckerClient) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Could not perform request to email checker service. Error: %w", err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	if err != nil {
		err = fmt.Errorf("Could not decode email checker response. Error: %w", err)
	}
	return resp, err
}
