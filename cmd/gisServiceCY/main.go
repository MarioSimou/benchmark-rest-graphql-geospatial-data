package main

import (
	"log"

	"io/ioutil"
	"path/filepath"

	"github.com/MarioSimou/gis-service-cy/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	var router = gin.Default()
	var port = "3000"
	var pop []byte
	var e error

	if pop, e = ioutil.ReadFile(filepath.Join("data", "pop-distribution.geojson")); e != nil {
		log.Fatalln(e)
	}
	var db = internal.Db{Population: pop}
	var contr = internal.New(&db)

	router.GET("/api/v1/cy/population", contr.GetPopulation)
	log.Fatalln(router.Run(":" + port))
}
