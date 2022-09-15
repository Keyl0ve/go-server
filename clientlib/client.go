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
	// URLを生成
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

//func (c *Client) Get(id string) (*Sample, error) {
//	reqUrl := c.url + "?id=" + id + ""
//	fmt.Println("reqUrl: ", reqUrl)
//	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
//	if err != nil {
//		return nil, fmt.Errorf(": %w", err)
//	}
//	fmt.Println("get req: ", req)
//	resp, err := c.Do(req)
//	if err != nil {
//		return nil, fmt.Errorf(": %w", err)
//	}
//	println(resp.StatusCode)
//
//	return resp, nil
//}
//
//func (c *Client) List() (*[]Sample, error) {
//	fmt.Println("c.url: ", c.url)
//	req, err := http.NewRequest(http.MethodGet, c.url, nil)
//	if err != nil {
//		return nil, fmt.Errorf(": %w", err)
//	}
//	fmt.Println("list req: ", req)
//	resp, err := c.Do(req)
//	if err != nil {
//		return nil, fmt.Errorf(": %w", err)
//	}
//	fmt.Println("list resp: ", resp)
//	println(resp.StatusCode)
//
//	return resp, nil
//}
//func (c *Client) Create(body io.Reader) (*Sample, error) {
//	sample := Sample{
//		ID:    "aa",
//		Name:  "Kyo",
//		Email: "kyo@gmail.com",
//	}
//	sampleJSON, err := json.Marshal(sample)
//	req, err := http.NewRequest(http.MethodPost, c.url, bytes.NewBuffer(sampleJSON))
//	if err != nil {
//		return nil, fmt.Errorf(": %w", err)
//	}
//
//	if body != nil {
//		req.Header.Set("Content-Type", "application/json")
//	}
//	resp, err := c.Do(req)
//	if err != nil {
//		return nil, fmt.Errorf(": %w", err)
//	}
//	println(resp.StatusCode)
//
//	return req, nil
//}
//func (c *Client) Delete(ctx context.Context, urlStr string, body io.Reader) (*Sample, error) {
//	req, err := http.NewRequest(http.MethodDelete, urlStr, body)
//	if err != nil {
//		return nil, fmt.Errorf(": %w", err)
//	}
//
//	if body != nil {
//		req.Header.Set("Content-Type", "application/json")
//	}
//
//	return req, nil
//}
//
