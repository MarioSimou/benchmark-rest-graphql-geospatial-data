package internal

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type controller struct {
	Db *mongo.Database
}

func New(db *mongo.Database) *controller {
	return &controller{db}
}

func (contr *controller) GetPopulation(c *gin.Context) {
	var collection = contr.Db.Collection("population")
	var features []Feature

	cur, e := collection.Find(context.Background(), bson.M{}, nil)
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": e.Error()})
		return
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var feature Feature
		if e := cur.Decode(&feature); e != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": e.Error()})
			return
		}

		features = append(features, feature)
	}
	c.JSON(http.StatusOK, Population{
		Type:     "FeatureCollection",
		Name:     "pop-distribution",
		Features: features,
	})
}
