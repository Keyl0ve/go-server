package main

import (
	"context"
	"fmt"
	"github.com/Keyl0ve/go-server/REST/clientlib"
	"net/http"
)

func main() {
	ctx := context.Background()
	client, err := clientlib.NewClient(http.DefaultClient, "http://localhost:8080/")
	if err != nil {
		panic(fmt.Sprintf("cannot create client: %v", err))
	}

	inputSample := &clientlib.Sample{
		ID:           "a",
		Name:         "kyo",
		Email:        "kyo@gmail.com",
		Books:        []*clientlib.Book{},
		FavoriteBook: "default book",
	}
	createdSample, err := client.CreateSample(ctx, inputSample)

	if err != nil {
		fmt.Errorf("cannnot create sample: %w", err)
	}

	fmt.Println("\nCREATE...")
	fmt.Println("created id: ", createdSample.ID)
	fmt.Println("created name: ", createdSample.Name)
	fmt.Println("created email: ", createdSample.Email)
	for i, v := range createdSample.Books {
		fmt.Println("created book[", i, "] id: ", "...", v.BookID)
		fmt.Println("created book[", i, "] name: ", "...", v.BookName)
		fmt.Println("created book[", i, "] parentID: ", "...", v.BookParentID)
	}
	fmt.Println("created favorite id: ", createdSample.FavoriteBook)

	inputSample2 := &clientlib.Sample{
		ID:           "aa",
		Name:         "kyo2",
		Email:        "kyo2@gmail.com",
		Books:        []*clientlib.Book{},
		FavoriteBook: "default book 2",
	}
	_, err = client.CreateSample(ctx, inputSample2)

	fmt.Println("created id: ", inputSample2.ID)
	fmt.Println("created name: ", inputSample2.Name)
	fmt.Println("created email: ", inputSample2.Email)
	for i, v := range inputSample2.Books {
		fmt.Println("created book[", i, "] id: ", "...", v.BookID)
		fmt.Println("created book[", i, "] name: ", "...", v.BookName)
		fmt.Println("created book[", i, "] parentID: ", "...", v.BookParentID)
	}
	fmt.Println("created favorite id: ", createdSample.FavoriteBook)

	fmt.Println("\nGET...")
	getSample, err := client.GetSample(ctx, createdSample.ID)
	if err != nil {
		fmt.Errorf("cannot get sample: %w", err)
	}
	fmt.Println("get sample id: ", getSample.ID)
	fmt.Println("get sample name: ", getSample.Name)
	fmt.Println("get sample email: ", getSample.Email)
	for i, v := range getSample.Books {
		fmt.Println("created book[", i, "] id: ", "...", v.BookID)
		fmt.Println("created book[", i, "] name: ", "...", v.BookName)
		fmt.Println("created book[", i, "] parentID: ", "...", v.BookParentID)
	}
	fmt.Println("get favorite id: ", getSample.FavoriteBook)

	fmt.Println("\nCREATE BOOK...")

	inputBook1 := &clientlib.Book{
		BookID:       "1",
		BookName:     "book1",
		BookParentID: inputSample.ID,
	}
	createdBook1, err := client.CreateBook(ctx, inputBook1)

	inputBook2 := &clientlib.Book{
		BookID:       "2",
		BookName:     "book2",
		BookParentID: inputSample.ID,
	}
	createdBook1, err = client.CreateBook(ctx, inputBook2)

	inputBook3 := &clientlib.Book{
		BookID:       "3",
		BookName:     "book3",
		BookParentID: inputSample2.ID,
	}
	createdBook1, err = client.CreateBook(ctx, inputBook3)
	for i, v := range createdBook1 {
		fmt.Println("created book[", i, "] sample id: ", "...", v.ID)
		fmt.Println("created book[", i, "] sample name: ", "...", v.Name)
		fmt.Println("created book[", i, "] sample email: ", "...", v.Email)
		for _, vv := range v.Books {
			fmt.Println("created book[", i, "] sample book id: ", "...", vv.BookID)
			fmt.Println("created book[", i, "] sample book name: ", "...", vv.BookName)
			fmt.Println("created book[", i, "] sample book parent id: ", "...", vv.BookParentID)
		}
		fmt.Println("created book[", i, "] sample favorite book: ", "...", v.FavoriteBook)
	}
	if err != nil {
		fmt.Println("cannot create book: ", err)
	}

	fmt.Println("\nGET Book...")
	getBook, err := client.GetBook(ctx, inputBook1.BookID)
	if err != nil {
		fmt.Errorf("cannot get book: %w", err)
	}
	fmt.Println("get book: ", getBook)
	fmt.Println("get book id: ", getBook.BookID)
	fmt.Println("get book name: ", getBook.BookName)
	fmt.Println("get book email: ", getBook.BookParentID)

	fmt.Println("\nLIST book...")
	listBook, err := client.ListBook(ctx)
	if err != nil {
		fmt.Errorf("cannot list book: %w", err)
	}
	for _, v := range listBook {
		fmt.Println("list book id: ", v.BookID)
		fmt.Println("list book name: ", v.BookName)
		fmt.Println("list book email: ", v.BookParentID)
	}

	fmt.Println("\nDELETE sample...")
	deleteSample, err := client.DeleteSample(ctx, inputSample2.ID)
	if err != nil {
		fmt.Errorf("cannot delete sample: %w", err)
	}
	fmt.Println("delete sample id: ", deleteSample.ID)
	fmt.Println("delete sample name: ", deleteSample.Name)
	fmt.Println("delete sample email: ", deleteSample.Email)
	for i, v := range deleteSample.Books {
		fmt.Println("delete book[", i, "] id: ", "...", v.BookID)
		fmt.Println("delete book[", i, "] name: ", "...", v.BookName)
		fmt.Println("delete book[", i, "] parentID: ", "...", v.BookParentID)
	}
	fmt.Println("delete favorite id: ", deleteSample.FavoriteBook)

	fmt.Println("\nDELETE book...")
	deleteBook, err := client.DeleteBook(ctx, inputBook1.BookID)
	if err != nil {
		fmt.Errorf("cannot delete book: %w", err)
	}
	fmt.Println("delete book id: ", deleteBook.BookID)
	fmt.Println("delete book name: ", deleteBook.BookName)
	fmt.Println("delete book parentID: ", deleteBook.BookParentID)

	fmt.Println("\nLIST2 sample...")
	listSample2, err := client.ListSample(ctx)
	if err != nil {
		fmt.Errorf("cannot list sample: %w", err)
	}
	for _, v := range listSample2 {
		fmt.Println("list sample id: ", v.ID)
		fmt.Println("list sample name: ", v.Name)
		fmt.Println("list sample email: ", v.Email)
		for i, vv := range v.Books {
			fmt.Println("list book[", i, "] id: ", "...", vv.BookID)
			fmt.Println("list book[", i, "] id: ", "...", vv.BookName)
			fmt.Println("list book[", i, "] id: ", "...", vv.BookParentID)
		}
		fmt.Println("list favorite id: ", v.FavoriteBook)
	}

	fmt.Println("\nLIST2 book...")
	listBook2, err := client.ListBook(ctx)
	if err != nil {
		fmt.Errorf("cannot list book: %w", err)
	}
	for _, v := range listBook2 {
		fmt.Println("list book id: ", v.BookID)
		fmt.Println("list book name: ", v.BookName)
		fmt.Println("list book email: ", v.BookParentID)
	}

	fmt.Println("\nGET favorite book before create...")
	getFavBook, err := client.GetFavBook(ctx, createdSample.ID)
	fmt.Println("get fav book name: ", getFavBook)

	fmt.Println("\nCREATE favorite book...")
	inputCreateFavRequest := &clientlib.CreateFavBookRequest{
		InputFavBookName: "hahaha book",
		SampleID:         createdSample.ID,
	}
	createdFavBook, err := client.CreateFavBook(ctx, inputCreateFavRequest)
	fmt.Println("created fav book name: ", createdFavBook)

	fmt.Println("\nGET favorite book after create...")
	getFavBook, err = client.GetFavBook(ctx, createdSample.ID)
	fmt.Println("get fav book name: ", getFavBook)

	fmt.Println("\nDELETE favorite book...")
	deletedFavBook, err := client.DeleteFavBook(ctx, createdSample.ID)
	fmt.Println("deleted fav book name: ", deletedFavBook)

	fmt.Println("\nGET favorite book after delete...")
	getFavBook, err = client.GetFavBook(ctx, createdSample.ID)
	fmt.Println("get fav book name: ", getFavBook)
}
