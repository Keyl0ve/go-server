package clientlib

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	client *http.Client
	url    *url.URL
}

func NewClient(client *http.Client, urlStr string) (*Client, error) {
	URL, err := url.Parse(urlStr)
	if err != nil {
		return nil, fmt.Errorf("url parse error")
	}

	return &Client{
		client: client,
		url:    URL,
	}, nil
}

func (c *Client) PathQuery(urlStr string, query url.Values) (string, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", fmt.Errorf(": %w", err)
	}
	u.RawQuery = query.Encode()

	return u.String(), nil
}

func (c *Client) NewRequest(ctx context.Context, method, urlStr string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, urlStr, body)
	if err != nil {
		return nil, fmt.Errorf(": %w", err)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

func (c *Client) Do(ctx context.Context, req *http.Request, payload interface{}) error {
	if ctx == nil {
		return fmt.Errorf("context must not be nil")
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf(": %w", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(payload); err != nil {
		return fmt.Errorf(": %w", err)
	}

	return nil
}