package main

import{
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinhzu/gorm/dialects/mysql"
	"github.com/JJ2820/Go_projects/go-bookstore/pkg/routes"
}

func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handler("/", r)
	log.Fatal(http.ListenAndServer("localhost:9010", r))
	
}