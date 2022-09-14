package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-server/sample/mypack"
	"io"
	"net/http"
	"net/url"
)

func main() {

	// URLを生成
	u, err := url.Parse("http://localhost:8080/")
	if err != nil {
		fmt.Errorf("invalid url: ")
	}
	u.Scheme = "https"

	mypack.
		//type Sample struct {
		//	ID    string `json:"id"`
		//	Name  string `json:"name"`
		//	Email string `json:"email"`
		//}

		// get query
		fmt.Println("GET query...")

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/?id=bb", nil)
	if err != nil {
		fmt.Errorf("Get failed: ")
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Errorf("Get failed: ")
	}
	println(resp.StatusCode)
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	// post
	fmt.Println("POST...")

	sample := Sample{ID: "aa", Name: "Kyo", Email: "kyo@gmail.com"}
	sampleJSON, err := json.Marshal(sample)
	req, err = http.NewRequest(http.MethodPost, "http://localhost:8080/", bytes.NewBuffer(sampleJSON))
	if err != nil {
		fmt.Errorf("Post failed: ")
	}
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Errorf("Post failed: ")
	}
	println(resp.StatusCode)
	body, _ = io.ReadAll(resp.Body)
	fmt.Println(string(body))

	// get
	fmt.Println("GET...")

	req, err = http.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	if err != nil {
		fmt.Errorf("Get failed: ")
	}
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Errorf("Get failed: ")
	}
	println(resp.StatusCode)
	body, _ = io.ReadAll(resp.Body)
	fmt.Println(string(body))

	// get query
	fmt.Println("DELETE query...")

	req, err = http.NewRequest(http.MethodDelete, "http://localhost:8080/?id=cc", nil)
	if err != nil {
		fmt.Errorf("Delete failed: ")
	}
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Errorf("Delete failed: ")
	}
	println(resp.StatusCode)
	body, _ = io.ReadAll(resp.Body)
	fmt.Println(string(body))

	defer resp.Body.Close()

}
