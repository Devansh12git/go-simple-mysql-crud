package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sakibn55/go-simple-mysql-crud/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/", r)
	log.Println("server started at port 9010")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
