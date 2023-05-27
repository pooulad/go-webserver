package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "form error :%v \n", err)
		return
	}
	fmt.Fprintf(w, "post request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "name is %s\n", name)
	fmt.Fprintf(w, "address is %s\n", address)
}

func helloHanler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 error", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not allowed", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello here")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHanler)

	fmt.Printf("starting server ar port 8080 \n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
