package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"net/http"

	"github.com/MarioSimou/gis-service-cy/internal"
	_ "github.com/lib/pq"
)

func main() {
	var port = os.Getenv("PORT")
	var e error
	var db *sql.DB

	if db, e = sql.Open("postgres", os.Getenv("POSTGRES_URI")); e != nil {
		log.Fatalln(e)
	}
	if e = db.Ping(); e != nil {
		log.Fatalln(e)
	}
	defer db.Close()

	var contr = internal.New(db)

	http.HandleFunc("/graphql", contr.GraphqlEndpoint)
	fmt.Println("The app listens on http://localhost:" + port + "/graphql")
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}
