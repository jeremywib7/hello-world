package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(writer, "This is the home page")
	if err != nil {
		return
	}
}

func About(writer http.ResponseWriter, request *http.Request) {
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
