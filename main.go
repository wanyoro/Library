package main

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func main() {
	app := App{}
	app.Initialize()
	app.RunServer()
}
