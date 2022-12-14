package clientlib

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type HTTPRequestOption func(req *http.Request)

func (c *Client) Get(ctx context.Context, pathQuery string, payload interface{}, httpRequestOptions ...HTTPRequestOption) error {
	req, err := c.NewRequest(ctx, http.MethodGet, pathQuery, nil)
	if err != nil {
		return fmt.Errorf(": %w", err)
	}
	//fmt.Println("get request before option", req)

	for _, option := range httpRequestOptions {
		option(req)
	}
	fmt.Println("get request", req)
	if err := c.Do(ctx, req, payload); err != nil {
		return fmt.Errorf(": %w", err)
	}
	return nil
}

func (c *Client) Post(ctx context.Context, endpoint string, body, payload interface{}) error {
	requestBody, err := json.Marshal(body)
	fmt.Println("post request Body:", string(requestBody))
	if err != nil {
		return fmt.Errorf(": %w", err)
	}

	req, err := c.NewRequest(ctx, http.MethodPost, endpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf(": %w", err)
	}
	fmt.Println("req : ", req)

	if err := c.Do(ctx, req, payload); err != nil {
		return fmt.Errorf(": %w", err)
	}

	return nil
}

func (c *Client) Delete(ctx context.Context, pathQuery string, payload interface{}, httpRequestOptions ...HTTPRequestOption) error {
	req, err := c.NewRequest(ctx, http.MethodDelete, pathQuery, nil)
	if err != nil {
		return fmt.Errorf(": %w", err)
	}

	for _, option := range httpRequestOptions {
		option(req)
	}

	if err := c.Do(ctx, req, payload); err != nil {
		return fmt.Errorf(": %w", err)
	}

	return nil
}
