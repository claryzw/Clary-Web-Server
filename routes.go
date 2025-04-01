//Defined routes and handling of HTTP requests
package main

import "fmt"
import "net/http"

func initializeRoutes() {
	http.HandleFunc("/", loggingMiddleware(rootHandler))
	http.HandleFunc("/about", loggingMiddleware(aboutHandler))
	http.HandleFunc("/api", loggingMiddleware(apiHandler))
}
