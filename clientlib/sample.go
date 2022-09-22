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
	BookID       string `json:"book_id"`
	BookName     string `json:"book_name"`
	BookParentID string `json:"book_parent_id"`
}

type GetSampleRequest struct {
	id string
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
	path := "sample"
	query := c.url.Query()
	query.Add("id", id)
	var payload Sample
	fmt.Println("get sample url: ", fmt.Sprintf("%s%s?%s", c.url.String(), path, query.Encode()))

	if err := c.Get(ctx, fmt.Sprintf("%s%s?%s", c.url.String(), path, query.Encode()), &payload); err != nil {
		return nil, fmt.Errorf(": %w", err)
	}

	return &payload, nil
}

func (c *Client) GetBook(ctx context.Context, bookID string) (*Book, error) {
	path := "book"
	query := c.url.Query()
	query.Add("book_id", bookID)
	var payload Book
	fmt.Println("get book url: ", fmt.Sprintf("%s%s?%s", c.url.String(), path, query.Encode()))
	if err := c.Get(ctx, fmt.Sprintf("%s%s?%s", c.url.String(), path, query.Encode()), &payload); err != nil {
		return nil, fmt.Errorf(": %w", err)
	}

	return &payload, nil
}

func (c *Client) ListSample(ctx context.Context) ([]*Sample, error) {
	path := "sample"
	var payload []*Sample
	fmt.Println("this is list sample url", fmt.Sprintf("%s%s", c.url.String(), path))

	if err := c.Get(ctx, fmt.Sprintf("%s%s", c.url.String(), path), &payload); err != nil {
		return nil, fmt.Errorf(": %w", err)
	}
	return payload, nil
}

func (c *Client) ListBook(ctx context.Context) ([]*Book, error) {
	path := "book"
	var payload []*Book
	fmt.Println("this is list book url", fmt.Sprintf("%s%s", c.url.String(), path))

	if err := c.Get(ctx, fmt.Sprintf("%s%s", c.url.String(), path), &payload); err != nil {
		fmt.Println("li", err)
		return nil, fmt.Errorf(": %w", err)
	}
	fmt.Println("pay", payload)
	return payload, nil
}

func (c *Client) CreateSample(ctx context.Context, input *Sample) (*Sample, error) {
	path := "sample"
	//for _, v := range input.Books {
	body := Sample{
		ID:           input.ID,
		Name:         input.Name,
		Email:        input.Email,
		Books:        []*Book{},
		FavoriteBook: input.FavoriteBook,
	}

	fmt.Println("start:::")
	fmt.Println(body.ID)
	fmt.Println(body.Name)
	fmt.Println(body.Email)

	var payload Sample
	fmt.Println("sample post", fmt.Sprintf("%s%s", c.url.String(), path))
	if err := c.Post(ctx, fmt.Sprintf("%s%s", c.url.String(), path), body, &payload); err != nil {
		return nil, fmt.Errorf(": %w", err)
	}
	fmt.Println("create sample payload: ", payload)

	return &payload, nil
}

func (c *Client) CreateBook(ctx context.Context, inputBook *Book) (*Sample, error) {
	path := "book"
	body := Book{
		BookID:       inputBook.BookID,
		BookName:     inputBook.BookName,
		BookParentID: inputBook.BookParentID,
	}

	//Sample[input.ParentID]Book = append(Sample[input.ParentID]Book, body)
	//fmt.Println("this is body id", body.ID)

	//query := c.url.Query()
	//query.Add("id", input.BookParentID)

	fmt.Println("book post: ", fmt.Sprintf("%s%s", c.url.String(), path))
	var payload Sample
	if err := c.Post(ctx, fmt.Sprintf("%s%s", c.url.String(), path), body, &payload); err != nil {
		fmt.Println("this is error in Post create book")
		return nil, fmt.Errorf(": %w", err)
	}
	fmt.Println("create book payload: ", payload)
	return &payload, nil
}

func (c *Client) DeleteSample(ctx context.Context, id string) (*Sample, error) {
	path := "sample"
	query := c.url.Query()
	query.Add("id", id)
	var payload Sample
	fmt.Println("sample delete url: ", fmt.Sprintf("%s%s?%s", c.url.String(), path, query.Encode()))

	if err := c.Delete(ctx, fmt.Sprintf("%s%s?%s", c.url.String(), path, query.Encode()), &payload); err != nil {
		return nil, fmt.Errorf(": %w", err)
	}

	return &payload, nil
}

func (c *Client) DeleteBook(ctx context.Context, bookID string) (*Book, error) {
	path := "book"
	query := c.url.Query()
	query.Add("book_id", bookID)
	var payload Book
	fmt.Println("book delete url: ", fmt.Sprintf("%s%s?%s", c.url.String(), path, query.Encode()))
	if err := c.Delete(ctx, fmt.Sprintf("%s%s?%s", c.url.String(), path, query.Encode()), &payload); err != nil {
		return nil, fmt.Errorf(": %w", err)
	}

	return &payload, nil
}
