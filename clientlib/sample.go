package clientlib

import (
	"context"
	"fmt"
	"net/url"
)

type Sample struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Email        string  `json:"email"`
	Books        []*Book `json:"book"`
	FavoriteBook string  `json:"favoriteBook"`
}

type Book struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ParentID string `json:"parentID"`
}

type GetSampleRequest struct {
	id string
}

type CreateSampleRequest struct {
	Sample
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

	//for _, v := range input.Books {
	body := CreateSampleRequest{Sample{
		ID:    input.ID,
		Name:  input.Name,
		Email: input.Email,
		Books: []*Book{
			//{
			//	ID:       v.ID,
			//	Name:     v.Name,
			//	ParentID: v.ParentID,
			//},
		},
		FavoriteBook: input.FavoriteBook,
	}}
	//}
	for _, v := range input.Books {
		body.Books = append(body.Books, v)
	}

	fmt.Println("start:::")
	fmt.Println(body.ID)
	fmt.Println(body.Name)
	fmt.Println(body.Email)
	fmt.Println(body.Books[0].ID)
	fmt.Println(body.Books[1].ID)
	fmt.Println(body.Books[2].ID)

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
