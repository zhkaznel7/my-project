package main 

import ("fmt"; "net/http"; "html/template"; "database/sql";
	_ "github.com/go-sql-driver/mysql")

type Article struct {
	Id uint16
	Title, Anons, FullText string
}

var posts = []Article{

}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	res, err := db.Query("SELECT * From `articles` ")
	if err != nil {
		panic(err)
	}
	defer res.Close() // Ресурсты жабуды ұмытпаймыз
	posts = []Article{}

	for res.Next() {
		var post Article
		// Scan арқылы базадағы деректі айнымалыға меншіктейміз
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.FullText )
		if err != nil {
			panic(err)
		}
		// Айнымалылар fmt.Sprintf-тің тырнақшасынан кейін үтірмен жазылады
		// fmt.Println(fmt.Sprintf("Post: %s with id: %d", post.Title, post.Id ))
		posts = append(posts, post)

	}

	t.ExecuteTemplate(w, "index", posts)
} 

func create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "create", nil)
} 

func save_article(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	full_text := r.FormValue("full_text")

	if title == "" || anons == "" || full_text == ""{
		fmt.Fprintf(w, "Не все данные заполнены")
	} else {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	res, err := db.Query(fmt.Sprintf("INSERT INTO `articles` (`title`, `anons`, `full_text`) VALUES ('%s', '%s', '%s')", title, anons, full_text))
	if err != nil {
		panic(err)
	}
	defer res.Close() // Ресурсты жабуды ұмытпаймыз

	http.Redirect(w, r, "/", http.StatusSeeOther)
	}

} 

func handleFunc() {
	http.HandleFunc("/", index)
	http.HandleFunc("/create/", create)
	http.HandleFunc("/save_article/", save_article)
	http.ListenAndServe(":8080", nil)
} 

func main(){
	
	handleFunc()
}