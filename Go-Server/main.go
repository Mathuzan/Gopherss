//Begins package main
package main

import (
	"fmt"
	"log"
	"net/http"
)

func formhandler(w http.ResponseWriter, r "http.Request") {
	if err := r.ParseForm(); err != nil {
		fmt.Printf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Post request Successful")
	name := r.formValue("name")
	address := r.formValue("address")
	fmt.Fprintf(w, "Name = %s\n", &name)
	fmt.Fprintf(w, "Address = %s\n", address)

}

func helloHandler(w http.ResponseWriter, r "http.Request") {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.statusNotFound)
		return
	} 

	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

//Main function begins
func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formhandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err !=nil {
		log.Fatal(err)
	}

}