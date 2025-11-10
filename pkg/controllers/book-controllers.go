package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/luminous44/ReadSphere/pkg/models"
	"github.com/luminous44/ReadSphere/pkg/utils"
) 

var NewBook models.Book

func CreateBook(){

}
func GetBooks(w http.ResponseWriter, r *http.Request){
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request){

	vars :=mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId,0,0)
    if err != nil{
		fmt.Println("error while parsing")
	}
	bookDetails, _:= models.GetBookById(ID)

	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateNewBook(w http.ResponseWriter, r *http.Request){

	CreateB := &models.Book{}
	utils.ParseBody(r,CreateB)
	b := CreateB.CreateBook()
	res,_ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){

	vars :=mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId,0,0)
    if err != nil{
		fmt.Println("error while parsing")
	}
	bookDetails := models.DeleteBook(ID)

	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook (w http.ResponseWriter, r *http.Request){

	var upDateB = &models.Book{}
	utils.ParseBody(r,upDateB)
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId,0,0)
    if err != nil{
		fmt.Println("error while parsing")
	}
	booksDetails, db := models.GetBookById(ID)
	if upDateB.Name != ""{
		booksDetails.Name = upDateB.Name
	}
	if upDateB.Author != ""{
		booksDetails.Author = upDateB.Author
	}
	if upDateB.Publication != ""{
		booksDetails.Publication = upDateB.Publication
	}
	db.Save(&booksDetails)
	res, _ := json.Marshal(booksDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)


}