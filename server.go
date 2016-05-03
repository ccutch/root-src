package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ccutch/root-src/controllers"
	"github.com/gorilla/mux"
)

var (
	port = os.Getenv("PORT")
	host = os.Getenv("HOST")
)

func main() {
	a := rs_controllers.Application()

	router := mux.NewRouter()
	router.Handle("/styles/", http.FileServer(http.Dir("/styles")))
	a.Mount(router)
	http.Handle("/", router)

	if port == "" {
		port = "4000"
	}
	addr := host + ":" + port
	fmt.Println("Server listening at", addr)
	http.ListenAndServe(addr, nil)
}
