package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// routes words
	r.HandleFunc("/words", CreateWord).Methods("POST")
	r.HandleFunc("/words", GetWords).Methods("GET")
	r.HandleFunc("/words/{id}", GetWord).Methods("GET")
	r.HandleFunc("/words/sentence/{id}", GetWordBySentenceId).Methods("GET")
	r.HandleFunc("/words/{id}", UpdateWord).Methods("PUT")
	r.HandleFunc("/words/{id}", DeleteWord).Methods("DELETE")

	// routes sentences
	r.HandleFunc("/sentences", CreateSentence).Methods("POST")
	r.HandleFunc("/sentences", GetSentences).Methods("GET")
	r.HandleFunc("/sentences/{id}", GetSentence).Methods("GET")
	r.HandleFunc("/sentences/file/{id}", GetSentenceByFileId).Methods("GET")
	r.HandleFunc("/sentences/{id}", UpdateSentence).Methods("PUT")
	r.HandleFunc("/sentences/{id}", DeleteSentence).Methods("DELETE")

	// routes files
	r.HandleFunc("/files", CreateFile).Methods("POST")
	r.HandleFunc("/files", GetFiles).Methods("GET")
	r.HandleFunc("/files/{id}", GetFileById).Methods("GET")
	r.HandleFunc("/files/categori/{id}", GetFileByCategoriId).Methods("GET")
	r.HandleFunc("/files/{id}", UpdateFile).Methods("PUT")
	r.HandleFunc("/files/{id}", DeleteFile).Methods("DELETE")

	// routes category
	r.HandleFunc("/category", CreateCategory).Methods("POST")
	r.HandleFunc("/category", GetCategories).Methods("GET")
	r.HandleFunc("/category/{id}", GetCategoryById).Methods("GET")
	r.HandleFunc("/category/{id}", UpdateCategory).Methods("PUT")
	r.HandleFunc("/category/{id}", DeleteCategory).Methods("DELETE")
	http.ListenAndServe(":8080", r)
}
