package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const baseURL = "https://slack.com/api"

type Client struct {
	token      string
	httpClient *http.Client
}

func NewClient(token string) *Client {
	return &Client{
		token:      token,
		httpClient: &http.Client{},
	}
}

func (c *Client) Get(method string, params url.Values) (json.RawMessage, error) {
	reqURL := fmt.Sprintf("%s/%s", baseURL, method)
	if len(params) > 0 {
		reqURL += "?" + params.Encode()
	}

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	var result map[string]json.RawMessage
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("parsing response: %w", err)
	}

	okRaw, exists := result["ok"]
	if exists {
		var ok bool
		if err := json.Unmarshal(okRaw, &ok); err == nil && !ok {
			errMsg := "unknown error"
			if e, ok := result["error"]; ok {
				json.Unmarshal(e, &errMsg)
			}
			return nil, fmt.Errorf("slack API error: %s", errMsg)
		}
	}

	return body, nil
}

func (c *Client) Post(method string, params url.Values) (json.RawMessage, error) {
	reqURL := fmt.Sprintf("%s/%s", baseURL, method)

	req, err := http.NewRequest("POST", reqURL, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	var result map[string]json.RawMessage
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("parsing response: %w", err)
	}

	okRaw, exists := result["ok"]
	if exists {
		var ok bool
		if err := json.Unmarshal(okRaw, &ok); err == nil && !ok {
			errMsg := "unknown error"
			if e, ok := result["error"]; ok {
				json.Unmarshal(e, &errMsg)
			}
			return nil, fmt.Errorf("slack API error: %s", errMsg)
		}
	}

	return body, nil
}

func (c *Client) PostJSON(method string, jsonBody []byte) (json.RawMessage, error) {
	reqURL := fmt.Sprintf("%s/%s", baseURL, method)

	req, err := http.NewRequest("POST", reqURL, strings.NewReader(string(jsonBody)))
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	var result map[string]json.RawMessage
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("parsing response: %w", err)
	}

	okRaw, exists := result["ok"]
	if exists {
		var ok bool
		if err := json.Unmarshal(okRaw, &ok); err == nil && !ok {
			errMsg := "unknown error"
			if e, ok := result["error"]; ok {
				json.Unmarshal(e, &errMsg)
			}
			return nil, fmt.Errorf("slack API error: %s", errMsg)
		}
	}

	return body, nil
}
