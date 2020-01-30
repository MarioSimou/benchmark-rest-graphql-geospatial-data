package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/MarioSimou/gis-service-cy/internal"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	var port = os.Getenv("PORT")
	var e error
	var client *mongo.Client

	if client, e = mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_URI"))); e != nil {
		log.Fatalln(e)
	}
	if e = client.Ping(context.Background(), nil); e != nil {
		log.Fatalln(e)
	}

	var contr = internal.New(client.Database(os.Getenv("MONGO_DB"), nil))
	http.HandleFunc("/graphql", contr.GraphqlEndpoint)
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}
