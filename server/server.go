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

type CreateFavBookRequest struct {
	InputFavBookName string `json:"input_fav_book_name"`
	SampleID         string `json:"sample_id"`
}

var Samples = map[string]*Sample{}
var Books = map[string]*Book{}

var ArraySample []Sample

func main() {
	http.HandleFunc("/sample", func(w http.ResponseWriter, r *http.Request) {

		var sampleData Sample
		if err := json.NewDecoder(r.Body).Decode(&sampleData); err != nil {
			fmt.Errorf("decode flow failed: %w", err)
		}

		switch r.Method {
		case http.MethodGet:
			queryParamSampleID := r.URL.Query().Get("id")
			if queryParamSampleID != "" {
				if err := GetSample(w, queryParamSampleID); err != nil {
					w.WriteHeader(http.StatusBadRequest)
					_, err := w.Write([]byte("get method failed: %s"))
					if err != nil {
						panic(err)
					}
				} else {
					w.WriteHeader(http.StatusOK)
				}
			} else {
				if err := ListSample(w); err != nil {
					w.WriteHeader(http.StatusBadRequest)
					_, err := w.Write([]byte("get method failed: %s"))
					if err != nil {
						panic(err)
					}
				} else {
					w.WriteHeader(http.StatusOK)
				}
			}

		case http.MethodPost:
			if err := PostSample(w, sampleData); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				_, err := w.Write([]byte("post method failed:"))
				if err != nil {
					panic(err)
				}
			} else {
				w.WriteHeader(http.StatusOK)
			}

		case http.MethodDelete:
			queryParamSampleID := r.URL.Query().Get("id")
			if queryParamSampleID == "" {
				panic(fmt.Sprintf("please provide sample id"))
			}
			if err := DeleteSample(w, queryParamSampleID); err != nil {
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

	// ここから book の処理
	http.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {

		var book Book
		if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
			fmt.Errorf("decode flow failed: %w", err)
		}

		switch r.Method {
		case http.MethodGet:
			queryParamBookID := r.URL.Query().Get("book_id")
			if err := GetBook(w, queryParamBookID); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				_, err := w.Write([]byte("get method failed: %s"))
				if err != nil {
					panic(err)
				}
			} else {
				w.WriteHeader(http.StatusOK)
			}

		case http.MethodPost:
			if err := PostBook(w, book); err != nil {
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
				panic(fmt.Sprintf("please provid book id"))
			}
			if err := DeleteBook(w, param); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				_, err := w.Write([]byte("delete method failed:"))
				if err != nil {
					panic(err)
				}
			}
		default:
			w.Write([]byte("book method failed:"))
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	// ここから favoriteBook の処理
	http.HandleFunc("/favoritebook", func(w http.ResponseWriter, r *http.Request) {

		var inputFavBookRequest CreateFavBookRequest
		if err := json.NewDecoder(r.Body).Decode(&inputFavBookRequest); err != nil {
			fmt.Errorf("decode flow failed: %w", err)
		}

		fmt.Println("       fav book")

		switch r.Method {
		case http.MethodGet:
			queryParamSampleID := r.URL.Query().Get("id")
			if queryParamSampleID != "" {
				if err := GetFavBook(w, queryParamSampleID); err != nil {
					w.WriteHeader(http.StatusBadRequest)
					_, err := w.Write([]byte("get method failed: %s"))
					if err != nil {
						panic(err)
					}
				} else {
					w.WriteHeader(http.StatusOK)
				}
			} else {
				if err := ListSample(w); err != nil {
					w.WriteHeader(http.StatusBadRequest)
					_, err := w.Write([]byte("get method failed: %s"))
					if err != nil {
						panic(err)
					}
				} else {
					w.WriteHeader(http.StatusOK)
				}
			}

		case http.MethodPost:
			if err := PostFavBook(w, inputFavBookRequest); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				_, err := w.Write([]byte("post method failed:"))
				if err != nil {
					panic(err)
				}
			} else {
				w.WriteHeader(http.StatusOK)
			}

		case http.MethodDelete:
			queryParamSampleID := r.URL.Query().Get("id")
			if queryParamSampleID == "" {
				panic(fmt.Sprintf("please provide sample id"))
			}
			if err := DeleteFavBook(w, queryParamSampleID); err != nil {
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
			return fmt.Errorf("cannot marshal : %w", err)
		}
		w.Write(samplesMar)
		return nil
	}

	sampleSlice := make([]*Sample, 0, len(Samples))

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

	// TODO samplesMar
	samplesMar, err := json.Marshal(sampleSlice)
	if err != nil {
		fmt.Println(w, "not marshlized")
	}
	fmt.Println("samplesMar", string(samplesMar))

	w.Write(samplesMar)
	fmt.Println("this is multi get req")

	return nil
}

func GetBook(w http.ResponseWriter, bookID string) error {
	//if book, ok := Books[id]; ok {
	//	booksMar, err := json.Marshal(book)
	//	if err != nil {
	//		return fmt.Errorf("not marshlized : %w", err)
	//	}
	//	if _, err := w.Write(booksMar); err != nil {
	//		return fmt.Errorf("cannot write : %w", err)
	//	}
	//	return nil
	//}
	fmt.Println("start get book (server)")
	fmt.Println("books", Books)
	fmt.Println("samples", Samples)
	if book, ok := Books[bookID]; ok {
		bookMar, err := json.Marshal(book)
		if err != nil {
			fmt.Println(w, "not marshlized")
		}
		fmt.Println("boookMaraaaa", string(bookMar))

		if _, err := w.Write(bookMar); err != nil {
			return fmt.Errorf("cannot write : %w", err)
		}
		//w.Write()
		return nil
	}

	fmt.Println("start list book (server)", Books)

	bookSlice := make([]*Book, 0, len(Books))
	for _, v := range Books {
		bookSlice = append(bookSlice, v)
	}

	bookMar, err := json.Marshal(bookSlice)
	if err != nil {
		fmt.Println(w, "not marshlized")
	}
	fmt.Println("2 boookMaraaaa", string(bookMar))

	if _, err := w.Write(bookMar); err != nil {
		return fmt.Errorf("cannot write : %w", err)
	}

	return nil
}

func GetFavBook(w http.ResponseWriter, sampleID string) error {
	//for _, v := range
	if sample, ok := Samples[sampleID]; ok {
		samplesMar, err := json.Marshal(sample.FavoriteBook)
		if err != nil {
			return fmt.Errorf("cannot marshal : %w", err)
		}
		w.Write(samplesMar)
		return nil
	}

	return fmt.Errorf("omgggggg")
}

func PostSample(w http.ResponseWriter, data Sample) error {
	if _, ok := Samples[data.ID]; ok {
		return fmt.Errorf("ID already exist")
	} else if data.ID == "" {
		return fmt.Errorf("fill in ID")
	} else if data.Name == "" {
		return fmt.Errorf("fill in NAME")
	} else if data.Email == "" {
		return fmt.Errorf("fill in Email")
	}

	Samples[data.ID] = &data
	samplesMar, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("not marshlized : %w ", err)
	}

	if _, err := w.Write(samplesMar); err != nil {
		return fmt.Errorf("cannnot write : %w ", err)
	}

	return nil
}

func PostBook(w http.ResponseWriter, inputBook Book) error {
	Books[inputBook.BookID] = &inputBook

	Samples[inputBook.BookParentID].Books = append(Samples[inputBook.BookParentID].Books, &inputBook)
	sampleSlice := make([]*Sample, 0, len(Samples))

	for _, v := range Samples {
		sampleSlice = append(sampleSlice, v)
	}

	sampleMar, err := json.Marshal(sampleSlice)
	if err != nil {
		return fmt.Errorf("not marshlized")
	}

	if _, err := w.Write(sampleMar); err != nil {
		return fmt.Errorf("cannot write : %w", err)
	}

	return nil
}

func PostFavBook(w http.ResponseWriter, inputFavBookRequest CreateFavBookRequest) error {
	if sample, ok := Samples[inputFavBookRequest.SampleID]; ok {
		sample.FavoriteBook = inputFavBookRequest.InputFavBookName
		Samples[inputFavBookRequest.SampleID].FavoriteBook = inputFavBookRequest.InputFavBookName

		favBookMar, err := json.Marshal(inputFavBookRequest.InputFavBookName)
		if err != nil {
			return fmt.Errorf("not marshlized : %w ", err)
		}

		if _, err := w.Write(favBookMar); err != nil {
			return fmt.Errorf("cannnot write : %w ", err)
		}
		return nil
	}

	//Samples[data.ID] = &data
	//favBookMar, err := json.Marshal(inputFavBookRequest.FavBookName)
	//if err != nil {
	//	return fmt.Errorf("not marshlized : %w ", err)
	//}
	//
	//if _, err := w.Write(favBookMar); err != nil {
	//	return fmt.Errorf("cannnot write : %w ", err)
	//}

	return nil
}

func DeleteSample(w http.ResponseWriter, id string) error {
	if sample, ok := Samples[id]; ok {
		delete(Samples, id)
		deletedSample, err := json.Marshal(sample)
		if err != nil {
			fmt.Errorf("not marshlized")
		}

		w.Write(deletedSample)
		return nil
	}

	return nil
}

func DeleteBook(w http.ResponseWriter, id string) error {
	if book, ok := Books[id]; ok {
		delete(Books, id)
		deletedBook, err := json.Marshal(book)
		if err != nil {
			fmt.Errorf("not marshlized")
		}

		w.Write(deletedBook)
		return nil
	}

	return nil
}

func DeleteFavBook(w http.ResponseWriter, sampleID string) error {
	if sample, ok := Samples[sampleID]; ok {
		//delete(Samples, sampleID)
		targetFavBook, err := json.Marshal(sample.FavoriteBook)
		if err != nil {
			fmt.Errorf("not marshlized")
		}

		sample.FavoriteBook = ""

		Samples[sampleID].FavoriteBook = ""

		w.Write(targetFavBook)
		return nil
	}

	return nil
}
