//This file will handle specific requests
package main

import "fmt"
import "net/http"
import "encoding/json"
import "time"

// This part will serve as a response structure for JSON responses
type Response struct { 
	Message string `json:"message"`
 	Status int `json:"status"`
	Time time.Time `json:"timestamp"`
}
	
func rootHandler(w http.ResponseWriter, r *http.Request) {
	// This only allows GET requests on root path
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	fmt.Fprint(w, "Welcome to Clarence's Web Server!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is an simple web server built with Go by Clarence Itai Msindo.")
}


func apiHandler(w http.ResponseWriter, r *http.Request) {
    // Set content type
    w.Header().Set("Content-Type", "application/json")
    
    // Create response
    resp := Response{
        Message: "API is working!",
        Status:  http.StatusOK,
        Time:    time.Now(),
    }
    
    // Convert to JSON and send
    json.NewEncoder(w).Encode(resp)
}

// Add middleware for logging
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        fmt.Printf("[%s] %s %s\n", start.Format(time.RFC3339), r.Method, r.URL.Path)
        next(w, r)
        fmt.Printf("[%s] Completed in %v\n", time.Now().Format(time.RFC3339), time.Since(start))
    }
}


