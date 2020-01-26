package internal

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
	Db *sql.DB
}

func New(db *sql.DB) *controller {
	return &controller{Db: db}
}

func (contr *controller) GetPopulation(c *gin.Context) {
	var features []Feature
	var sql = `
	SELECT
	'Feature' as type,
	jsonb_build_object(
		'gmlId', population."gml_id",
		'localId', population."localId",
		'namespace', population."namespace",
		'versionId', population."versionId",
		'localisedCharacterString', population."LocalisedCharacterString",
		'measurementUnitUom',population."measurementUnit_uom",
		'notCountedProportion', population."notCountedProportion",
		'beginPosition', population."beginPosition",
		'endPosition', population."endPosition",
		'duration', population."duration"
	) as properties,
	st_asgeojson(population."geom")::jsonb as geometry
	FROM population;
	`

	rows, e := contr.Db.Query(sql)
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": e.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var feature Feature
		if e := rows.Scan(
			&feature.Type,
			&feature.Properties,
			&feature.Geometry,
		); e != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": e.Error()})
			return
		}
		features = append(features, feature)
	}

	c.JSON(http.StatusOK, gin.H{
		"type":     "FeatureCollection",
		"name":     "pop-distrubution",
		"features": features,
	})
}
