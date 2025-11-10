package routes

import (
	"github.com/gorilla/mux"
	"github.com/luminous44/ReadSphere/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router){

	router.HandleFunc("/book", controllers.CreateNewBook).Methods("POST")
	router.HandleFunc("/book", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("//book/{id}", controllers.DeleteBook).Methods("DELETE")

}
