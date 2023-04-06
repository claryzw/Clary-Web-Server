//This file will handle specific requests
package main

import "fmt"
import "net/http"


func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Simple, Web Server!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is an simple web server built with Go by Clarence Itai Msindo.")
}
