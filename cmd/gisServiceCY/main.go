package main

import (
	"context"
	"log"

	"github.com/MarioSimou/gis-service-cy/internal"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	var router = gin.Default()
	var port = "3000"
	var e error
	var client *mongo.Client

	if client, e = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017")); e != nil {
		log.Fatalln(e)
	}
	if e = client.Ping(context.Background(), nil); e != nil {
		log.Fatalln(e)
	}

	var contr = internal.New(client.Database("gis", nil))
	router.GET("/api/v1/cy/population", contr.GetPopulation)
	log.Fatalln(router.Run(":" + port))
}
