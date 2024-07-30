package controllers

import {
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/JJ2820/Go_projects/go-bookstore/utils"
	"github.com/JJ2820/Go_projects/go-bookstore/utils/pkg"	
}

var NewBook models.Book // new book is of type book struct 

func GetBooks(w http.ResponseWriter, r *http.Request){
	newBooks:=models.GetAllBooks()
	res, _ :=json.Marshal(newBooks)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookByID(w http.ResponseWriter, r *http.Request){
	var := mux.Vars(r)
	bookID := vars["bookID"]
	ID, err:= strconv.ParseInt(bookID,0,0)
	if err != nil {
		fmt.Println("Error While Parsing")
	}
	book.Details, _:= modules.GetBookByID(ID)
	res, _:= json.Marshal(bookDetails)
	w.header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponceWriter, r *http.Request){
	createBook;= &models.Book{}
	utils.ParseBody(r, CreateBook)
	b:= CraeteBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteBook(w http.ResponceWriter, r *http.Request){
	var := mux.vars(r)
	bookId := vars["bookID"]
	ID, err := strconv.ParseInt(bookID,0,0)
	if err != nil {
		fmt.Println("Error While Parsing")
	}
	book := modules.DeleteBook(ID)
	res, +:= json.Marshal(book)
	w.Header().Set("Connect-Type","pkglication/json")
	w.WriterHead(http.StatusOK)
	w.Writer(res)
}

func UpdateBook(w http.ResponceWriter, r *http.Request){
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookID"]
	ID, err := strconv.ParseInt(bookID,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db:=models.GetBookById(ID)
	if updateBook.Name != ""{
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != ""{
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != ""{
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}