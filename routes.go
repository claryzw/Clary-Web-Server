//Defined routes and handling of HTTP requests
package main

import "fmt"
import "net/http"


func initializeRoutes() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/about", aboutHandler)
}