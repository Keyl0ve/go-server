package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

var Samples = map[string]*Sample{}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		var data Sample
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			fmt.Errorf("decode flow failed: %w", err)
		}

		//one := Sample{
		//	ID:           "a",
		//	Name:         "kyo",
		//	Email:        "kyo@gmail.com",
		//	Book:         Book{ID: "1", Name: "book1", ParentID: "parentIDa"},
		//	FavoriteBook: "anpanman",
		//}
		//
		//fmt.Println("sampleee: ", one)

		switch r.Method {
		case http.MethodGet:
			param := r.URL.Query().Get("id")

			if err := Get(w, param); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				_, err := w.Write([]byte("get method failed: %s"))
				if err != nil {
					panic(err)
				}
			} else {
				w.WriteHeader(http.StatusOK)
			}

		case http.MethodPost:
			if err := Post(w, data); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				_, err := w.Write([]byte("post method failed:"))
				if err != nil {
					panic(err)
				}
			} else {
				w.WriteHeader(http.StatusOK)
			}

		case http.MethodDelete:
			param := r.URL.Query().Get("id")
			if param == "" {
				fmt.Fprintf(w, "please provide id: ")
			}
			if err := Delete(w, param); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				_, err := w.Write([]byte("delete method failed:"))
				if err != nil {
					panic(err)
				}
			}
		default:
			w.Write([]byte("delete method failed:"))
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	http.ListenAndServe(":8080", nil)
}

func Get(w http.ResponseWriter, id string) error {
	if sample, ok := Samples[id]; ok {
		samplesMar, err := json.Marshal(sample)
		if err != nil {
			fmt.Fprintf(w, "not marshlized")
		}
		io.WriteString(w, string(samplesMar))
		return nil
	}

	sampleSlice := make([]*Sample, 0, len(Samples))

	for _, v := range Samples {
		sampleSlice = append(sampleSlice, v)
	}

	samplesMar, err := json.Marshal(sampleSlice)
	if err != nil {
		fmt.Fprintf(w, "not marshlized")
	}

	io.WriteString(w, string(samplesMar))

	return nil
}

func Post(w http.ResponseWriter, data Sample) error {
	if _, ok := Samples[data.ID]; ok {
		return fmt.Errorf("ID already exist")
	}
	if data.ID == "" {
		return fmt.Errorf("fill in ID")
	}
	if data.Name == "" {
		return fmt.Errorf("fill in NAME")
	}
	if data.Email == "" {
		return fmt.Errorf("fill in Email")
	}

	Samples[data.ID] = &data

	samplesMar, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintf(w, "not marshlized")
	}
	io.WriteString(w, string(samplesMar))

	return nil
}

func Delete(w http.ResponseWriter, id string) error {
	if sample, ok := Samples[id]; ok {
		delete(Samples, id)
		samplesMar, err := json.Marshal(sample)
		if err != nil {
			fmt.Errorf("not marshlized")
		}

		io.WriteString(w, string(samplesMar))
		return nil
	}

	return nil
}
