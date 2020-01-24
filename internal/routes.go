package internal

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
	Db *Db
}

func New(db *Db) *controller {
	return &controller{Db: db}
}

func (contr *controller) GetPopulation(c *gin.Context) {
	var pop Population
	if e := json.Unmarshal(contr.Db.Population, &pop); e != nil {
		c.JSON(http.StatusInternalServerError, e.Error())
		return
	}
	c.JSON(http.StatusOK, pop)
}
