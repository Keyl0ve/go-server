package main

import (
	"encoding/json"
	"fmt"
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
	BookID       string `json:"book_id"`
	BookName     string `json:"book_name"`
	BookParentID string `json:"book_parent_id"`
}

var Samples = map[string]*Sample{}
var Books = map[string]*Book{}

func main() {
	http.HandleFunc("/sample", func(w http.ResponseWriter, r *http.Request) {

		var data Sample
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			fmt.Errorf("decode flow failed: %w", err)
		}

		switch r.Method {
		case http.MethodGet:
			param := r.URL.Query().Get("id")
			if param != "" {
				fmt.Println("this is sample get req")
				if err := GetSample(w, param); err != nil {
					fmt.Println("bad request get")
					w.WriteHeader(http.StatusBadRequest)
					_, err := w.Write([]byte("get method failed: %s"))
					if err != nil {
						panic(err)
					}
				} else {
					fmt.Println("this is get req")
					w.WriteHeader(http.StatusOK)
				}
			} else {
				fmt.Println("this is sample list req")
				if err := ListSample(w); err != nil {
					fmt.Println("bad request get")
					w.WriteHeader(http.StatusBadRequest)
					_, err := w.Write([]byte("get method failed: %s"))
					if err != nil {
						panic(err)
					}
				} else {
					fmt.Println("this is list req")
					w.WriteHeader(http.StatusOK)
				}
			}

		case http.MethodPost:
			if err := PostSample(w, data); err != nil {
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
			if err := DeleteSample(w, param); err != nil {
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

	// これから book の処理
	http.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {

		var data *Book
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			fmt.Errorf("decode flow failed: %w", err)
		}

		switch r.Method {
		case http.MethodGet:
			param := r.URL.Query().Get("book_id")

			if err := GetBook(w, param); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				_, err := w.Write([]byte("get method failed: %s"))
				if err != nil {
					panic(err)
				}
			} else {
				w.WriteHeader(http.StatusOK)
			}

		case http.MethodPost:
			fmt.Println("koko", data)
			if err := PostBook(w, *data); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				_, err := w.Write([]byte("post method failed:"))
				if err != nil {
					panic(err)
				}
			} else {
				w.WriteHeader(http.StatusOK)
			}

		case http.MethodDelete:
			param := r.URL.Query().Get("book_id")
			if param == "" {
				fmt.Fprintf(w, "please provide id: ")
			}
			if err := DeleteBook(w, param); err != nil {
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

func GetSample(w http.ResponseWriter, id string) error {
	if sample, ok := Samples[id]; ok {
		samplesMar, err := json.Marshal(sample)
		if err != nil {
			fmt.Println(w, "not marshlized")
		}
		fmt.Println("samplesMaraaaa", string(samplesMar))
		w.Write(samplesMar)
		fmt.Println("this is 1 get req")
		return nil
	}

	sampleSlice := make([]*Sample, 0, len(Samples))

	//for _, v := range Samples {
	//	sampleSlice = append(sampleSlice, v)
	//}

	samplesMar, err := json.Marshal(sampleSlice)
	if err != nil {
		fmt.Println(w, "not marshlized")
	}

	w.Write(samplesMar)
	fmt.Println("this is multi get req")

	return nil
}

func ListSample(w http.ResponseWriter) error {
	sampleSlice := make([]*Sample, 0, len(Samples))

	for _, v := range Samples {
		sampleSlice = append(sampleSlice, v)
	}

	fmt.Println("sample slice:", sampleSlice)

	samplesMar, err := json.Marshal(sampleSlice)

	//samplesMar, err := json.Marshal(Samples)
	if err != nil {
		fmt.Println(w, "not marshlized")
	}
	fmt.Println("samplesMar", string(samplesMar))

	w.Write(samplesMar)
	fmt.Println("this is multi get req")

	return nil
}

func GetBook(w http.ResponseWriter, id string) error {
	if book, ok := Books[id]; ok {
		booksMar, err := json.Marshal(book)
		if err != nil {
			fmt.Println(w, "not marshlized")
		}
		w.Write(booksMar)
		return nil
	}

	bookSlice := make([]*Book, 0, len(Books))

	for _, v := range Books {
		bookSlice = append(bookSlice, v)
	}

	booksMar, err := json.Marshal(bookSlice)
	if err != nil {
		fmt.Println(w, "not marshlized")
	}

	w.Write(booksMar)

	return nil
}

func PostSample(w http.ResponseWriter, data Sample) error {
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
		fmt.Println(w, "not marshlized")
	}
	w.Write(samplesMar)

	return nil
}

func PostBook(w http.ResponseWriter, data Book) error {
	//if _, ok := Samples[data.ID]; ok {
	//	return fmt.Errorf("ID already exist")
	//}
	//if data == nil {
	//	return fmt.Errorf("fill in ID")
	//}
	//bookSlice := make([]*Book, 0, len(Books))
	//
	//bookSlice = append(bookSlice, data)
	Books[data.BookID] = &data

	samplesMar, err := json.Marshal(data)
	if err != nil {
		fmt.Println(w, "not marshlized")
	}
	w.Write(samplesMar)

	return nil
}

func DeleteSample(w http.ResponseWriter, id string) error {
	if sample, ok := Samples[id]; ok {
		delete(Samples, id)
		samplesMar, err := json.Marshal(sample)
		if err != nil {
			fmt.Errorf("not marshlized")
		}

		w.Write(samplesMar)
		return nil
	}

	return nil
}

func DeleteBook(w http.ResponseWriter, id string) error {
	if book, ok := Books[id]; ok {
		delete(Books, id)
		booksMar, err := json.Marshal(book)
		if err != nil {
			fmt.Errorf("not marshlized")
		}

		w.Write(booksMar)
		return nil
	}

	return nil
}
