package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ccutch/root-src/controllers"
)

var (
	port = os.Getenv("PORT")
	host = os.Getenv("HOST")
)

func main() {
	a := rs_controllers.Application()
	http.Handle("/", a)

	if port == "" {
		port = "4000"
	}
	addr := host + ":" + port
	fmt.Println("Server listening at", addr)
	http.ListenAndServe(addr, nil)
}
