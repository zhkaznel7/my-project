package main 

import ("fmt"
		"net/http")

func home_page(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Go Hello")
}

func handleRequest(){
	http.HandleFunc("/", home_page)
	http.ListenAndServe(":8080", nil)
}

func main(){
	handleRequest()
}