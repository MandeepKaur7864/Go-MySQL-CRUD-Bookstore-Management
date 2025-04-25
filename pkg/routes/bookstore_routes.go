package routes

import (
	"bookstore_management/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterBookstoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/getBooks", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/getBookById/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/createBook", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/updateBook/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/deleteBook/{id}", controllers.DeleteBook).Methods("DELETE")

}
