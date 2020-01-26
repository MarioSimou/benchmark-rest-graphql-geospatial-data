package internal

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Feature struct {
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
	Geometry   Geometry   `json:"geometry"`
}

type Properties struct {
	GmlId                    string `json:"gmlId,omitempty"`
	LocalId                  string `json:"localId,omitempty"`
	Namespace                string `json:"namespace,omitempty"`
	VersionId                string `json:"versionId,omitempty"`
	LocalisedCharacterString string `json:"localisedCharacterString,omitempty"`
	MeasurementUnitUom       string `json:"measurementUnitUom,omitempty"`
	NotCountedProportion     string `json:"notCountedProportion,omitempty"`
	BeginPosition            string `json:"beginPosition,omitempty"`
	EndPosition              string `json:"endPosition,omitempty"`
	Duration                 string `json:"duration,omitempty"`
}

func (p *Properties) Value() (driver.Value, error) {
	fmt.Println("scanning properties...")
	return json.Marshal(p)
}
func (p *Properties) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &p)
}

type Geometry struct {
	Type        string          `json:"type,omitempty"`
	Coordinates [][][][]float64 `json:"coordinates,omitempty"`
}

func (g *Geometry) Value() (driver.Value, error) {
	fmt.Println("value geometry....")
	return json.Marshal(g)
}
func (g *Geometry) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &g)
}

type Population struct {
	Type string `json:"type,omitempty"`
	Name string `json:"name,omitempty"`
	Crs  struct {
		Type       string `json:"type"`
		Properties struct {
			Name string `json:"name"`
		} `json:"properties"`
	} `json:"crs,omitempty"`
	Features []Feature `json:"features"`
}
