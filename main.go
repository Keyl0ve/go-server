package main

import (
	"context"
	"fmt"
	"github.com/Keyl0ve/go-server/clientlib"
	"net/http"
)

func main() {
	ctx := context.Background()
	client, err := clientlib.NewClient(http.DefaultClient, "http://localhost:8080/")
	if err != nil {
		panic(fmt.Sprintf("cannot create client: %v", err))
	}

	rawInput := clientlib.Sample{
		ID:    "a",
		Name:  "kyo",
		Email: "kyo@gmail.com",
		Books: []*clientlib.Book{
			{
				ID:       "1",
				Name:     "book1",
				ParentID: "a",
			},
			{
				ID:       "2",
				Name:     "book2",
				ParentID: "b",
			},
			{
				ID:       "3",
				Name:     "book3",
				ParentID: "c",
			},
		},
		FavoriteBook: "novel1",
	}

	input := &clientlib.CreateSampleRequest{Sample: rawInput}
	createdSample, err := client.CreateSample(ctx, input)
	if err != nil {
		fmt.Errorf("cannnot create sample: %w", err)
	}
	fmt.Println("CREATE...")
	fmt.Println("created id: ", createdSample.ID)
	fmt.Println("created name: ", createdSample.Name)
	fmt.Println("created email: ", createdSample.Email)
	for i, v := range createdSample.Books {
		fmt.Println("created book[", i, "] id: ", "...", v.ID)
		fmt.Println("created book[", i, "] name: ", "...", v.Name)
		fmt.Println("created book[", i, "] parentID: ", "...", v.ParentID)
	}
	fmt.Println("created favorite id: ", createdSample.FavoriteBook)

	fmt.Println("GET...")
	getSample, err := client.GetSample(ctx, createdSample.ID)
	if err != nil {
		fmt.Errorf("cannot get sample: %w", err)
	}
	fmt.Println("get sample id: ", getSample.ID)
	fmt.Println("get sample name: ", getSample.Name)
	fmt.Println("get sample email: ", getSample.Email)
	for i, v := range getSample.Books {
		fmt.Println("created book[", i, "] id: ", "...", v.ID)
		fmt.Println("created book[", i, "] name: ", "...", v.Name)
		fmt.Println("created book[", i, "] parentID: ", "...", v.ParentID)
	}
	fmt.Println("get favorite id: ", getSample.FavoriteBook)

	fmt.Println("LIST...")
	listSample, err := client.ListSample(ctx)
	if err != nil {
		fmt.Errorf("cannot list sample: %w", err)
	}
	for _, v := range listSample {
		fmt.Println("list sample id: ", v.ID)
		fmt.Println("list sample name: ", v.Name)
		fmt.Println("list sample email: ", v.Email)
		for i, vv := range v.Books {
			fmt.Println("created book[", i, "] id: ", "...", vv.ID)
			fmt.Println("created book[", i, "] name: ", "...", vv.Name)
			fmt.Println("created book[", i, "] parentID: ", "...", vv.ParentID)
		}
		fmt.Println("list favorite id: ", v.FavoriteBook)
	}

	fmt.Println("DELETE...")
	deleteSample, err := client.DeleteSample(ctx, getSample.ID)
	if err != nil {
		fmt.Errorf("cannot delete sample: %w", err)
	}

	fmt.Println("delete sample id: ", deleteSample.ID)
	fmt.Println("delete sample name: ", deleteSample.Name)
	fmt.Println("delete sample email: ", deleteSample.Email)
	for i, v := range deleteSample.Books {
		fmt.Println("created book[", i, "] id: ", "...", v.ID)
		fmt.Println("created book[", i, "] name: ", "...", v.Name)
		fmt.Println("created book[", i, "] parentID: ", "...", v.ParentID)
	}
	fmt.Println("delete favorite id: ", deleteSample.FavoriteBook)

	fmt.Println("LIST2...")
	listSample2, err := client.ListSample(ctx)
	if err != nil {
		fmt.Errorf("cannot list sample: %w", err)
	}
	for _, v := range listSample2 {
		fmt.Println("list sample id: ", v.ID)
		fmt.Println("list sample name: ", v.Name)
		fmt.Println("list sample email: ", v.Email)
		for i, vv := range v.Books {
			fmt.Println("created book[", i, "] id: ", "...", vv.ID)
			fmt.Println("created book[", i, "] id: ", "...", vv.Name)
			fmt.Println("created book[", i, "] id: ", "...", vv.ParentID)
		}
		fmt.Println("list favorite id: ", v.FavoriteBook)
	}
}
