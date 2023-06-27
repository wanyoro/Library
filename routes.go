package main

import (
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gorilla/mux"
)

func (a *App) Initialize() {
	var err error
	const DNS = "postgres://postgres@localhost/Lib_DB?sslmode=disable"
	a.DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect")
	}

	a.DB.Debug().AutoMigrate(&Person{}, &Book{})
	a.Router = mux.NewRouter().StrictSlash(true)
	a.InitializeRoutes()
}

func (a *App) InitializeRoutes() {
	a.Router.Use(SetContentTypeMiddleware)
	a.Router.HandleFunc("/", a.SignUp).Methods("POST")
	a.Router.HandleFunc("/login", a.Login).Methods("POST")
	a.Router.HandleFunc("/home", Home).Methods("GET")
	a.Router.HandleFunc("/user/{id}", a.GetUserWithId).Methods("GET")
	a.Router.HandleFunc("/people/", AuthJWTVerify(a.GetAllPersons)).Methods("GET")
	a.Router.HandleFunc("/users/{id}", a.UpdateUser).Methods("PUT")
	a.Router.HandleFunc("/book/", AuthJWTVerify(a.CreateBook)).Methods("POST")
	a.Router.HandleFunc("/books/{id}", a.GetBookWithId).Methods("GET")
	a.Router.HandleFunc("/books/", AuthJWTVerify(a.GetAllBooks)).Methods("GET")
	a.Router.HandleFunc("/deleteBook/{id}", AuthJWTVerify(a.DeleteBook)).Methods("DELETE")
	a.Router.HandleFunc("/userswithbks", a.GetAllUsersWithBks).Methods("GET")
	a.Router.HandleFunc("/userswithoutbooks", a.GetAllUsersWithoutBks).Methods("GET")
	a.Router.HandleFunc("/availablebooks", a.GetUnAssignedBooks).Methods("GET")
	a.Router.HandleFunc("/book/{person_id}", GetUserWithBkId).Methods("GET")
	a.Router.HandleFunc("/peoplebooks", a.GetPeopleAndBooks).Methods("GET")
	a.Router.HandleFunc("/assignedBks", GetAllAssignedBooks).Methods("GET")

}

func (a *App) RunServer() {
	log.Printf("\nStarting on port 8000")
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}
