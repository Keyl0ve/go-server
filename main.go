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

	inputSample := &clientlib.Sample{
		ID:           "a",
		Name:         "kyo",
		Email:        "kyo@gmail.com",
		Books:        []*clientlib.Book{},
		FavoriteBook: "novel1",
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
		FavoriteBook: "novel2",
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
	fmt.Println("createdBook1  id:  ", createdBook1.BookID)
	fmt.Println("createdBook1  name:  ", createdBook1.BookName)
	fmt.Println("createdBook1  parentID:  ", createdBook1.BookParentID)
	if err != nil {
		fmt.Println("cannot create book: ", err)
	}

	inputBook2 := &clientlib.Book{
		BookID:       "2",
		BookName:     "book2",
		BookParentID: inputSample2.ID,
	}
	createdBook2, err := client.CreateBook(ctx, inputBook2)
	fmt.Println("createdBook2 id:  ", createdBook2.BookID)
	fmt.Println("createdBook2 name:  ", createdBook2.BookName)
	fmt.Println("createdBook2 parentID:  ", createdBook2.BookParentID)

	//inputBook3 := &clientlib.Book{
	//	BookID:       "3",
	//	BookName:     "book3",
	//	BookParentID: createdSample.ID,
	//}
	//createdBook3, err := client.CreateBook(ctx, *inputBook3)
	//fmt.Println("createdBook1 createdBook3 id:  ", createdBook3.BookID)
	//fmt.Println("createdBook1 createdBook3 name:  ", createdBook3.BookName)
	//fmt.Println("createdBook1 createdBook3 parentID:  ", createdBook3.BookParentID)

	//createdSample.Books = append(createdSample.Books, createdBook1)

	//createdBook2, err := client.CreateBook(ctx, inputBook2)
	//if err != nil {
	//	fmt.Println("cannot create book: ", err)
	//}
	//createdSample.Books = append(createdSample.Books, createdBook2)
	//createdBook3, err := client.CreateBook(ctx, inputBook3)
	//if err != nil {
	//	fmt.Println("cannot create book: ", err)
	//}
	//createdSample.Books = append(createdSample.Books, createdBook3)

	fmt.Println("\nGET Book...")
	getBook, err := client.GetBook(ctx, createdBook1.BookID)
	if err != nil {
		fmt.Errorf("cannot get book: %w", err)
	}
	fmt.Println("get book: ", getBook)
	fmt.Println("get book id: ", getBook.BookID)
	fmt.Println("get book name: ", getBook.BookName)
	fmt.Println("get book email: ", getBook.BookParentID)

	fmt.Println("\nLIST sample...")
	listSample, err := client.ListSample(ctx)
	if err != nil {
		fmt.Errorf("cannot list sample: %w", err)
	}
	for _, v := range listSample {
		fmt.Println("list sample id: ", v.ID)
		fmt.Println("list sample name: ", v.Name)
		fmt.Println("list sample email: ", v.Email)
		for i, vv := range v.Books {
			fmt.Println("created book[", i, "] id: ", "...", vv.BookID)
			fmt.Println("created book[", i, "] name: ", "...", vv.BookName)
			fmt.Println("created book[", i, "] parentID: ", "...", vv.BookParentID)
		}
		fmt.Println("list favorite id: ", v.FavoriteBook)
	}

	fmt.Println("\nLIST book...")
	listBook, err := client.ListBook(ctx)
	if err != nil {
		fmt.Errorf("cannot list book: %w", err)
	}
	for _, v := range listBook {
		fmt.Println("list book id: ", v.BookID)
		fmt.Println("list book name: ", v.BookName)
		fmt.Println("list book email: ", v.BookParentID)
		//for i, vv := range v.BookID {
		//	fmt.Println("created book[", i, "] id: ", "...", vv)
		//	fmt.Println("created book[", i, "] name: ", "...", vv.BookName)
		//	fmt.Println("created book[", i, "] parentID: ", "...", vv.BookParentID)
		//}
		//fmt.Println("list favorite id: ", v.FavoriteBook)
	}

	fmt.Println("\nDELETE...")
	deleteSample, err := client.DeleteSample(ctx, getSample.ID)
	if err != nil {
		fmt.Errorf("cannot delete sample: %w", err)
	}

	fmt.Println("delete sample id: ", deleteSample.ID)
	fmt.Println("delete sample name: ", deleteSample.Name)
	fmt.Println("delete sample email: ", deleteSample.Email)
	for i, v := range deleteSample.Books {
		fmt.Println("created book[", i, "] id: ", "...", v.BookID)
		fmt.Println("created book[", i, "] name: ", "...", v.BookName)
		fmt.Println("created book[", i, "] parentID: ", "...", v.BookParentID)
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
			fmt.Println("created book[", i, "] id: ", "...", vv.BookID)
			fmt.Println("created book[", i, "] id: ", "...", vv.BookName)
			fmt.Println("created book[", i, "] id: ", "...", vv.BookParentID)
		}
		fmt.Println("list favorite id: ", v.FavoriteBook)
	}
}
