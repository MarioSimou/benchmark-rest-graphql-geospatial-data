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
			),
		 'geometry', st_asgeojson(st_transform(population."geom",4326), 4)::jsonb
	)
	 ) as features
	FROM population
	) as p
	`

	if e := contr.Db.QueryRow(sql).Scan(&pop); e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": e.Error()})
		return
	}
	c.JSON(http.StatusOK, pop)
}
