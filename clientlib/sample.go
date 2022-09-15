package clientlib

import (
	"context"
	"fmt"
	"net/url"
)

type GetSampleRequest struct {
	id string
}
type Sample struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateSampleRequest struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type DeleteSampleRequest struct {
	id string
}

func (i *GetSampleRequest) Query() url.Values {
	values := url.Values{
		"id": []string{i.id},
	}
	return values
}

func (c *Client) GetSample(ctx context.Context, id string) (*Sample, error) {
	query := c.url.Query()
	query.Add("id", id)
	var payload Sample
	if err := c.Get(ctx, fmt.Sprintf("%s?%s", c.url.String(), query.Encode()), &payload); err != nil {
		return nil, fmt.Errorf(": %w", err)
	}

	return &payload, nil
}

func (c *Client) ListSample(ctx context.Context) ([]*Sample, error) {
	var payload []*Sample
	if err := c.Get(ctx, c.url.String(), &payload); err != nil {
		return nil, fmt.Errorf(": %w", err)
	}
	return payload, nil
}
func (c *Client) CreateSample(ctx context.Context, input *CreateSampleRequest) (*Sample, error) {
	body := CreateSampleRequest{
		ID:    input.ID,
		Name:  input.Name,
		Email: input.Email,
	}

	var payload Sample
	if err := c.Post(ctx, c.url.String(), body, &payload); err != nil {
		return nil, fmt.Errorf(": %w", err)
	}

	return &payload, nil
}

func (c *Client) DeleteSample(ctx context.Context, id string) (*Sample, error) {
	query := c.url.Query()
	query.Add("id", id)
	var payload Sample
	if err := c.Delete(ctx, fmt.Sprintf("%s?%s", c.url.String(), query.Encode()), &payload); err != nil {
		return nil, fmt.Errorf(": %w", err)
	}

	return &payload, nil
}
