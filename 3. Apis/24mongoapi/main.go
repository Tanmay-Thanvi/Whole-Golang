package main

import (
	"fmt"
	"net/http"

	"github.com/Tanmay-Thanvi/mongoapi/router"
)

func main() {
	fmt.Println("MongoDB Api")
	r := router.Router()
	fmt.Println("Server is getting started....")

	http.ListenAndServe(":4000", r)
	fmt.Println("Listening at port 4000...")
}
