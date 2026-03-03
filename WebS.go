package main

import (
	"fmt"
	"net/http"

)

// home route
func homeHandler(w http.Response, r *http.Request){
	fmt.Fprintf(w, "Welcome to my first Go web server!..")
}

// Aboute Route
func abouthandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is about page.")
}

func main(){
	// Register route
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/", abouthandler)

	fmt.Println("Server running on http://localhost:8080")
	// fmt.Println("Opens the last ")

	// start server
	http.ListenAndServe(":8080", nil)
}