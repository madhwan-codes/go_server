package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful. ")

	name := req.Form.Get("name")
	address := req.Form.Get("address")

	fmt.Printf("name: %s, address: %s\n", name, address)

}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if req.Method != "GET" {
		http.Error(w, "Meathod Not Allowed. ", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w, "Hello World!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
