# Clary-Simple-Web-Server-With-Go

This project is a simple web server built with the Go programming language. The goal of this project is to gain experience with building web servers using Go and the net/http package.

## Features

- **Clean Architecture**: Separation of concerns with routes, handlers, and middleware
- **RESTful API**: JSON response endpoints with proper status codes
- **Middleware Support**: Request logging with timing information
- **Error Handling**: Proper HTTP error responses

## Getting Started
To run the web server, you will need to have Go installed on your computer. You can download and install the latest version of Go from the official Go website.

### Prerequisites
- Go 1.16 or higher
- Visual Studio Code (recommended)

### Installation & Running the Server

1. Git clone https://github.com/claryzw/Clary-Web-Server.git
2. Navigate to the project directory in Visual Studio Code
3. Open the terminal in VS Code (Terminal > New Terminal) and run:
   ```
   go run main.go routes.go handler.go

   ```
4. The server is now running and accessible at http://localhost:8080

## File Directory

This project includes the following files:

* main.go - The main file that initializes and starts the server.
* routes.go - Defines the routes and handling of HTTP requests with middleware.
* handler.go - Handles specific requests and middleware implementation.

## Frameworks

This project does not use any external web frameworks.

## Built With:

* Go - The programming language used
* net/http - Go's built-in package for building HTTP servers


## Acknowledgments

* Go documentation

* Inspiration from other open-source Go web servers.
