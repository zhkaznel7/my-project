package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name"` // stringv емес, string болуы керек
	Age  uint16 `json:"age"`
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// SELECT сұранысын міндетті түрде тырнақшаға (" ") аламыз
	res, err := db.Query("SELECT `name`, `age` FROM `users` ")
	if err != nil {
		panic(err)
	}
	defer res.Close() // Ресурсты жабуды ұмытпаймыз

	for res.Next() {
		var user User
		// Scan арқылы базадағы деректі айнымалыға меншіктейміз
		err = res.Scan(&user.Name, &user.Age)
		if err != nil {
			panic(err)
		}
		// Айнымалылар fmt.Sprintf-тің тырнақшасынан кейін үтірмен жазылады
		fmt.Println(fmt.Sprintf("User: %s with age %d", user.Name, user.Age))
	}
}