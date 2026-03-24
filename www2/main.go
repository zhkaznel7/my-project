package main 

import ("fmt"; "net/http"; "html/template")

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", nil)
} 

func handleFunc() {
	http.HandleFune("/", index)
	http.ListenAndServe(":8080", nil)
} 

func main(){
	
	handleFunc()
}