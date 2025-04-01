// Clarence Itai Msindo's Web Server
package main

//The imports for the web server
import "fmt"
import "net/http"
import "log"

// The functions for the web server
func main() {

	// Initialize routes
	initializeRoutes()
	
	// Start server with better handling
	fmt.Println("Server strating on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
