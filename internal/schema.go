package internal

import (
	"context"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var populationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Population",
		Fields: graphql.Fields{
			"type": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"features": &graphql.Field{
				Type: graphql.NewList(
					graphql.NewObject(
						graphql.ObjectConfig{
							Name: "Feature",
							Fields: graphql.Fields{
								"type": &graphql.Field{
									Type: graphql.String,
								},
								"properties": &graphql.Field{
									Type: graphql.NewObject(
										graphql.ObjectConfig{
											Name: "Feature",
											Fields: graphql.Fields{
												"gmlId": &graphql.Field{
													Type: graphql.String,
												},
												"localId": &graphql.Field{
													Type: graphql.String,
												},
												"namespace": &graphql.Field{
													Type: graphql.String,
												},
												"versionId": &graphql.Field{
													Type: graphql.String,
												},
												"localisedCharacterString": &graphql.Field{
													Type: graphql.String,
												},
												"measurementUnitUom": &graphql.Field{
													Type: graphql.String,
												},
												"notCountedProportion": &graphql.Field{
													Type: graphql.String,
												},
												"beginPosition": &graphql.Field{
													Type: graphql.String,
												},
												"endPosition": &graphql.Field{
													Type: graphql.String,
												},
												"duration": &graphql.Field{
													Type: graphql.String,
												},
											},
										},
									),
								},
								"geometry": &graphql.Field{
									Type: graphql.NewObject(
										graphql.ObjectConfig{
											Name: "Geometry",
											Fields: graphql.Fields{
												"type": &graphql.Field{
													Type: graphql.String,
												},
												"coordinates": &graphql.Field{
													Type: graphql.NewList(
														graphql.NewList(
															graphql.NewList(
																graphql.NewList(
																	graphql.Float,
																),
															),
														),
													),
												},
											},
										},
									),
								},
							},
						},
					),
				),
			},
		},
	},
)

func NewSchema(db *mongo.Database) (graphql.Schema, error) {
	return graphql.NewSchema(
		graphql.SchemaConfig{
			Query: graphql.NewObject(
				graphql.ObjectConfig{
					Name: "Query",
					Fields: graphql.Fields{
						"population": &graphql.Field{
							Type: populationType,
							Resolve: func(p graphql.ResolveParams) (interface{}, error) {
								var collection = db.Collection("population")
								var features []Feature

								var options = options.Find()
								options.Projection = bson.D{{"type", true}, {"properties", true}, {"geometry", true}}
								cur, e := collection.Find(context.Background(), bson.M{}, options)
								if e != nil {
									return nil, e
								}

								for cur.Next(context.Background()) {
									var feature Feature
									if cur.Decode(&feature); e != nil {
										return nil, e
									}
									features = append(features, feature)
								}

								return Population{
									Type:     "FeatureCollection",
									Name:     "pop-distribution",
									Features: features,
								}, nil
							},
						},
					},
				},
			),
		},
	)
}
