// Clarence Itai Msindo's Web Server
package main

//The imports for the web server
import "fmt"
import "net/http"

// The functions for the web server
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "This is a Simple Web Server from Clarence Itai Msindo!")
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
