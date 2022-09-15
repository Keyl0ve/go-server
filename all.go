package main

import (
	"fmt"
	"github.com/Keyl0ve/go-server/clientlib"
	"net/http"
)

var url = "http://localhost:8080/"

func check(cli *http.Client, url string) {
	_, err := clientlib.NewClient(cli, url)
	if err != nil {
		panic("can not create client")
	}

	fmt.Println("success!")
}
