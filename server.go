package main

import (
	"fmt"
	"net/http"

	"github.com/ccutch/root-src/controllers"
)

func main() {
	a := rs_controllers.Application()
	http.Handle("/", a)

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
