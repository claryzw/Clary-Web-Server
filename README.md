# Clary-Simple-Web-Server-With-Go

This project is a simple web server built with the Go programming language. The goal of this project is to gain experience with building web servers using Go and the net/http package.

## Features

* Basic server that listens on a specific port and responds to incoming requests.
* Routing to handle different URL paths.
* Handling different HTTP methods like GET, POST, PUT, DELETE etc.
* Handling different types of requests like JSON, XML etc.

## Getting Started

To run the web server, you will need to have Go installed on your computer. You can download and install the latest version of Go from the official Go website.

Once you have Go installed, you can run the web server using the following command:

    go run main.go

This will start the web server on port 8080. You can visit the web server in your web browser by navigating to http://localhost:8080.

## File Directory

This project includes the following files:

* main.go - The main file that initializes and starts the server.
* routes.go - Defines the routes and handling of HTTP requests.
* handler.go - Handles specific requests.

## Frameworks

This project also uses popular web frameworks such as Gin, Echo, Revel etc. that are built on top of net/http package and abstract some of the low-level details of the package and make it easy to build web applications.
Built With:

* Go - The programming language used
* net/http - Go's built-in package for building HTTP servers


## Acknowledgments

Go documentation

Inspiration from other open-source Go web servers.
