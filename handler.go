//This file will handle specific requests
package main

import "fmt"
import "net/http"


func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is an example web server built with Go.")
}