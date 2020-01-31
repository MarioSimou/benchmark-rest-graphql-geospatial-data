package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/MarioSimou/gis-service-cy/internal"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	var router = gin.New()
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

	router.GET("/api/v1/cy/population", contr.GetPopulation)
	log.Fatalln(router.Run(":" + port))
}
