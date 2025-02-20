package main

import "net/http"

func main() {
	setUpApi()
	http.ListenAndServe(":8080", nil) // start the server on port 8080
}
func setUpApi() {
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
}
