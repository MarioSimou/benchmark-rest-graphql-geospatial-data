package main

import (
	"database/sql"
	"log"

	"github.com/MarioSimou/gis-service-cy/internal"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	var router = gin.Default()
	var port = "3000"
	var e error
	var db *sql.DB

	if db, e = sql.Open("postgres", "postgresql://msimou:msimou@localhost:5432/gis?sslmode=disable"); e != nil {
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
