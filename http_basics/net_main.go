// package main

// import (
// 	"fmt"
// 	"net/http"
// )


// func main(){
// 	http.HandleFunc("/", helloHandler)
// 	http.HandleFunc("/about", aboutHandler)
// 	fmt.Println("🚀 Server started at http://localhost:8080")
// 	http.ListenAndServe(":8080", nil)
// }

// func helloHandler(w http.ResponseWriter, r *http.Request){
// 	fmt.Println(w, "Hello from Go HTTP")
// }

// func aboutHandler(w http.ResponseWriter, r *http.Request){
// 	fmt.Println(w, "About Page, Powered by go")
// }

// // POST requests can be handled similarly by defining a handler function
// // and using http.HandleFunc to route the requests.

// func postHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodPost {
// 		fmt.Fprintln(w, "Handling POST request")
// 	} else {
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 	}
// }