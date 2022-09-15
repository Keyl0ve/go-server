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

	input := &clientlib.CreateSampleRequest{ID: "b", Name: "kyo", Email: "kyo@email.com"}
	createdSample, err := client.CreateSample(ctx, input)
	if err != nil {
		fmt.Errorf("cannnot create sample: %w", err)
	}

	fmt.Println("created id: ", createdSample.ID)
	fmt.Println("created name: ", createdSample.Name)
	fmt.Println("created email: ", createdSample.Email)

	getSample, err := client.GetSample(ctx, createdSample.ID)
	if err != nil {
		fmt.Errorf("cannot get sample: %w", err)
	}
	fmt.Println("get sample id: ", getSample.ID)
	fmt.Println("get sample name: ", getSample.Name)
	fmt.Println("get sample email: ", getSample.Email)

	listSample, err := client.ListSample(ctx)
	if err != nil {
		fmt.Errorf("cannot list sample: %w", err)
	}
	for _, v := range listSample {
		fmt.Println("list sample id: ", v.ID)
		fmt.Println("list sample name: ", v.Name)
		fmt.Println("list sample email: ", v.Email)
	}

	deleteSample, err := client.DeleteSample(ctx, getSample.ID)
	if err != nil {
		fmt.Errorf("cannot delete sample: %w", err)
	}

	fmt.Println("delete sample id: ", deleteSample.ID)
	fmt.Println("delete sample name: ", deleteSample.Name)
	fmt.Println("delete sample email: ", deleteSample.Email)

	listSample2, err := client.ListSample(ctx)
	if err != nil {
		fmt.Errorf("cannot list sample: %w", err)
	}
	for _, v := range listSample2 {
		fmt.Println("list sample id: ", v.ID)
		fmt.Println("list sample name: ", v.Name)
		fmt.Println("list sample email: ", v.Email)
	}
}

//Client 作成
//client := clientlib.NewClient(http.DefaultClient)
//
//fmt.Println("this is get")
//resp, err := client.Get("bb")
//if err != nil {
//	panic("get panic")
//}
//body, err := io.ReadAll(resp.Body)
//if err != nil {
//	fmt.Errorf("get body error")
//}
//fmt.Println(string(body))
//
//fmt.Println("this is list")
//resp, _ := client.List()
//body, _ := io.ReadAll(resp.Body)
//fmt.Println(string(body))
//
//fmt.Println("this is create")
//sample := clientlib.Sample{
//	ID:    "aa",
//	Name:  "Kyo",
//	Email: "kyo@gmail.com",
//}
//sampleJSON, err := json.Marshal(sample)
//client.Create(nil)
//body, _ = io.ReadAll(resp.Body)
//fmt.Println(string(body))
//
//client.Delete("http://localhost:8080/?id=bb", nil)
//
//resp, err := http.DefaultClient.Do(req)
//if err != nil {
//	fmt.Errorf("Get failed: ")
//}
//println(resp.StatusCode)
//body, _ := io.ReadAll(resp.Body)
//fmt.Println(string(body))
//
//post
//fmt.Println("POST...")
//
//sample := clientlib.Sample{
//	ID:    "bb",
//	Name:  "Kyo",
//	Email: "kyo@gmail.com",
//}
//
//sampleJSON, err := json.Marshal(sample)
//req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/", bytes.NewBuffer(sampleJSON))
//if err != nil {
//	fmt.Errorf("Post failed: ")
//}
//resp, err = http.DefaultClient.Do(req)
//fmt.Println("default resp", resp)
//if err != nil {
//	fmt.Errorf("Post failed: ")
//}
//println(resp.StatusCode)
//body, _ = io.ReadAll(resp.Body)
//fmt.Println(string(body))

//// get
//fmt.Println("GET...")
//
//req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
//if err != nil {
//	fmt.Errorf("Get failed: ")
//}
//fmt.Println("ddd req: ", req)
//resp, err = http.DefaultClient.Do(req)
//if err != nil {
//	fmt.Errorf("Get failed: ")
//}
//fmt.Println("ddd resp: ", resp)
//println(resp.StatusCode)
//body, _ = io.ReadAll(resp.Body)
//fmt.Println(string(body))
//
//// get query
//fmt.Println("DELETE query...")
//
//req, err = http.NewRequest(http.MethodDelete, "http://localhost:8080/?id=cc", nil)
//if err != nil {
//	fmt.Errorf("Delete failed: ")
//}
//resp, err = http.DefaultClient.Do(req)
//if err != nil {
//	fmt.Errorf("Delete failed: ")
//}
//println(resp.StatusCode)
//body, _ = io.ReadAll(resp.Body)
//fmt.Println(string(body))

//}
