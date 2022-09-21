package main

import (
	"encoding/json"
	"go-news/fun"
	"go-news/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateWord(w http.ResponseWriter, r *http.Request) {
	// get request formvalue
	sentenceId := r.FormValue("sentenceId")
	pos := r.FormValue("pos")
	word := r.FormValue("word")

	// convert string to int
	sentenceIdInt, _ := strconv.Atoi(sentenceId)

	// parsing data
	res := fun.CreateWord(sentenceIdInt, pos, word)

	// status code
	w.WriteHeader(http.StatusCreated)
	// response
	w.Header().Set("Content-Type", "application/json")
	// return word
	json.NewEncoder(w).Encode(res)
}

func GetWords(w http.ResponseWriter, r *http.Request) {
	// get all words
	words := fun.GetWords()

	// status code
	w.WriteHeader(http.StatusOK)
	// response
	w.Header().Set("Content-Type", "application/json")
	// return words
	json.NewEncoder(w).Encode(words)
}

func GetWord(w http.ResponseWriter, r *http.Request) {
	// get request vars
	vars := mux.Vars(r)
	id := vars["id"]

	// convert string to int
	idInt, _ := strconv.Atoi(id)

	// get word
	word := fun.GetWord(idInt)

	// status code
	w.WriteHeader(http.StatusOK)
	// response
	w.Header().Set("Content-Type", "application/json")
	// return word
	json.NewEncoder(w).Encode(word)
}

func GetWordBySentenceId(w http.ResponseWriter, r *http.Request) {
	// get request vars
	vars := mux.Vars(r)
	id := vars["id"]

	// convert string to int
	idInt, _ := strconv.Atoi(id)

	// get word
	word := fun.GetWordBySentenceId(idInt)

	// status code
	w.WriteHeader(http.StatusOK)
	// response
	w.Header().Set("Content-Type", "application/json")
	// return word
	json.NewEncoder(w).Encode(word)
}

func UpdateWord(w http.ResponseWriter, r *http.Request) {
	// get request vars
	vars := mux.Vars(r)
	id := vars["id"]

	// get request formvalue
	sentenceId := r.FormValue("sentenceId")
	pos := r.FormValue("pos")
	word := r.FormValue("word")

	// convert string to int
	idInt, _ := strconv.Atoi(id)
	sentenceIdInt, _ := strconv.Atoi(sentenceId)

	// parsing data
	res := fun.UpdateWord(idInt, sentenceIdInt, pos, word)

	// status code
	w.WriteHeader(http.StatusOK)
	// response
	w.Header().Set("Content-Type", "application/json")
	// return word
	json.NewEncoder(w).Encode(res)
}

func DeleteWord(w http.ResponseWriter, r *http.Request) {
	// get request vars
	vars := mux.Vars(r)
	id := vars["id"]

	// convert string to int
	idInt, _ := strconv.Atoi(id)

	// parsing data
	res := fun.DeleteWord(idInt)

	// status code
	w.WriteHeader(http.StatusOK)
	// response
	w.Header().Set("Content-Type", "application/json")
	// return word
	json.NewEncoder(w).Encode(res)
}

func CreateSentence(w http.ResponseWriter, r *http.Request) {
	// get request formvalue
	fileId := r.FormValue("fileId")
	sentence := r.FormValue("sentence")
	posSentence := r.FormValue("posSentence")
	pos := r.FormValue("pos")

	// convert string to int
	fileIdInt, _ := strconv.Atoi(fileId)

	// parsing data
	res := fun.CreateSentence(fileIdInt, sentence, posSentence, pos)

	// status code
	w.WriteHeader(http.StatusCreated)
	// response
	w.Header().Set("Content-Type", "application/json")
	// return sentence
	json.NewEncoder(w).Encode(res)
}

func GetSentences(w http.ResponseWriter, r *http.Request) {
	var sentences []model.Sentence
	// get all sentences
	fun.GetSentences().Scan(&sentences)

	// status code
	w.WriteHeader(http.StatusOK)
	// response
	w.Header().Set("Content-Type", "application/json")
	// return sentences
	json.NewEncoder(w).Encode(sentences)
}

func GetSentence(w http.ResponseWriter, r *http.Request) {
	// get request vars
	vars := mux.Vars(r)
	id := vars["id"]

	// convert string to int
	idInt, _ := strconv.Atoi(id)

	// get sentence
	sentence := fun.GetSentence(idInt)

	// status code
	w.WriteHeader(http.StatusOK)
	// response
	w.Header().Set("Content-Type", "application/json")
	// return sentence
	json.NewEncoder(w).Encode(sentence)
}

func GetSentenceByFileId(w http.ResponseWriter, r *http.Request) {
	// get request vars
	vars := mux.Vars(r)
	id := vars["id"]

	// convert string to int
	idInt, _ := strconv.Atoi(id)

	// get sentence
	sentence := fun.GetSentenceByFileId(idInt)

	// status code
	w.WriteHeader(http.StatusOK)
	// response
	w.Header().Set("Content-Type", "application/json")
	// return sentence
	json.NewEncoder(w).Encode(sentence)
}

func UpdateSentence(w http.ResponseWriter, r *http.Request) {
	// get request vars
	vars := mux.Vars(r)
	id := vars["id"]

	// get request formvalue
	fileId := r.FormValue("fileId")
	sentence := r.FormValue("sentence")
	posSentence := r.FormValue("posSentence")
	pos := r.FormValue("pos")

	// convert string to int
	idInt, _ := strconv.Atoi(id)
	fileIdInt, _ := strconv.Atoi(fileId)

	// parsing data
	res := fun.UpdateSentence(idInt, fileIdInt, sentence, posSentence, pos)

	// status code
	w.WriteHeader(http.StatusOK)
	// response
	w.Header().Set("Content-Type", "application/json")
	// return sentence
	json.NewEncoder(w).Encode(res)
}

func DeleteSentence(w http.ResponseWriter, r *http.Request) {
	// get request vars
	vars := mux.Vars(r)
	id := vars["id"]

	// convert string to int
	idInt, _ := strconv.Atoi(id)

	// parsing data
	res := fun.DeleteSentence(idInt)

	// status code
	w.WriteHeader(http.StatusOK)
	// response
	w.Header().Set("Content-Type", "application/json")
	// return sentence
	json.NewEncoder(w).Encode(res)
}

func CreateFile(w http.ResponseWriter, r *http.Request) {
	// get request formvalue
	categoriId := r.FormValue("categoriId")
	fileName := r.FormValue("fileName")
	content := r.FormValue("content")

	// convert string to int
	categoriIdInt, _ := strconv.Atoi(categoriId)

	// parsing data
	res := fun.CreateFile(categoriIdInt, fileName, content)

	// status code
	w.WriteHeader(http.StatusCreated)
	// response
	w.Header().Set("Content-Type", "application/json")
	// return file
	json.NewEncoder(w).Encode(res)
}

func GetFiles(w http.ResponseWriter, r *http.Request) {
	// get all files
	files := fun.GetFiles()

	// status code
	w.WriteHeader(http.StatusOK)
	// response
	w.Header().Set("Content-Type", "application/json")
	// return files
	json.NewEncoder(w).Encode(files)
}

func GetFileById(w http.ResponseWriter, r *http.Request) {
	// get request vars
	vars := mux.Vars(r)
	id := vars["id"]

	// convert string to int
	idInt, _ := strconv.Atoi(id)

	// get file
	file := fun.GetFileById(idInt)

	// status code
	w.WriteHeader(http.StatusOK)
	// response
	w.Header().Set("Content-Type", "application/json")
	// return file
	json.NewEncoder(w).Encode(file)
}

func GetFileByCategoriId(w http.ResponseWriter, r *http.Request) {
	// get request vars
	vars := mux.Vars(r)
	id := vars["id"]

	// convert string to int
	idInt, _ := strconv.Atoi(id)

	// get file
	file := fun.GetFileByCategoriId(idInt)

	// status code
	w.WriteHeader(http.StatusOK)
	// response
	w.Header().Set("Content-Type", "application/json")
	// return file
	json.NewEncoder(w).Encode(file)
}

func UpdateFile(w http.ResponseWriter, r *http.Request) {
	// get request vars
	vars := mux.Vars(r)
	id := vars["id"]

	// get request formvalue
	categoriId := r.FormValue("categoriId")
	fileName := r.FormValue("fileName")
	content := r.FormValue("content")

	// convert string to int
	idInt, _ := strconv.Atoi(id)
	categoriIdInt, _ := strconv.Atoi(categoriId)

	// parsing data
	res := fun.UpdateFile(idInt, categoriIdInt, fileName, content)

	// status code
	w.WriteHeader(http.StatusOK)
	// response
	w.Header().Set("Content-Type", "application/json")
	// return file
	json.NewEncoder(w).Encode(res)
}

func DeleteFile(w http.ResponseWriter, r *http.Request) {
	// get request vars
	vars := mux.Vars(r)
	id := vars["id"]

	// convert string to int
	idInt, _ := strconv.Atoi(id)

	// parsing data
	res := fun.DeleteFile(idInt)

	// status code
	w.WriteHeader(http.StatusOK)
	// response
	w.Header().Set("Content-Type", "application/json")
	// return file
	json.NewEncoder(w).Encode(res)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	// get request formvalue
	categoryName := r.FormValue("categoryName")

	// parsing data
	res := fun.CreateCategory(categoryName)

	// status code
	w.WriteHeader(http.StatusCreated)
	// response
	w.Header().Set("Content-Type", "application/json")
	// return category
	json.NewEncoder(w).Encode(res)
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	// get all categories
	categories := fun.GetCategories()

	// status code
	w.WriteHeader(http.StatusOK)
	// response
	w.Header().Set("Content-Type", "application/json")
	// return categories
	json.NewEncoder(w).Encode(categories)
}

func GetCategoryById(w http.ResponseWriter, r *http.Request) {
	// get request vars
	vars := mux.Vars(r)
	id := vars["id"]

	// convert string to int
	idInt, _ := strconv.Atoi(id)

	// get category
	category := fun.GetCategoryById(idInt)

	// status code
	w.WriteHeader(http.StatusOK)
	// response
	w.Header().Set("Content-Type", "application/json")
	// return category
	json.NewEncoder(w).Encode(category)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	// get request vars
	vars := mux.Vars(r)
	id := vars["id"]

	// get request formvalue
	categoryName := r.FormValue("categoryName")

	// convert string to int
	idInt, _ := strconv.Atoi(id)

	// parsing data
	res := fun.UpdateCategory(idInt, categoryName)

	// status code
	w.WriteHeader(http.StatusOK)
	// response
	w.Header().Set("Content-Type", "application/json")
	// return category
	json.NewEncoder(w).Encode(res)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	// get request vars
	vars := mux.Vars(r)
	id := vars["id"]

	// convert string to int
	idInt, _ := strconv.Atoi(id)

	// parsing data
	res := fun.DeleteCategory(idInt)

	// status code
	w.WriteHeader(http.StatusOK)
	// response
	w.Header().Set("Content-Type", "application/json")
	// return category
	json.NewEncoder(w).Encode(res)
}