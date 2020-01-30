package internal

import (
	"database/sql"

	"github.com/graphql-go/graphql"
)

var PopulationType = graphql.NewObject(
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
											Name: "Properties",
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
												"notCountedProportion": &graphql.Field{
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
													Type: graphql.NewList(graphql.NewList(graphql.NewList(graphql.NewList(graphql.Float)))),
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

func NewSchema(db *sql.DB) (graphql.Schema, error) {
	return graphql.NewSchema(
		graphql.SchemaConfig{
			Query: graphql.NewObject(
				graphql.ObjectConfig{
					Name: "Query",
					Fields: graphql.Fields{
						"population": &graphql.Field{
							Type: PopulationType,
							Resolve: func(p graphql.ResolveParams) (interface{}, error) {
								var pop Population
								var sql = `
								SELECT row_to_json(p)
								FROM (
									SELECT
										'FeatureCollection' as type,
										'pop-distribution' as name,
										json_agg(
											jsonb_build_object(
												'type', 'Feature',
												'properties', jsonb_build_object(
													'gmlId',population."gml_id",
													'localId', population."localId",
													'namespace', population."namespace",
													'versionId', population."versionId",
													'localisedCharacterString', population."LocalisedCharacterString",
													'notCountedProportion', population."notCountedProportion",
													'endPosition',population."endPosition",
													'duration', population."duration"
												),
												'geometry', st_asgeojson(population."geom")::jsonb
											)
										) as features 
									FROM population
								) p
								`

								if e := db.QueryRow(sql).Scan(&pop); e != nil {
									return nil, e
								}
								return pop, nil
							},
						},
					},
				},
			),
		},
	)
}
