package internal

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Body struct {
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
	Query         string                 `json:"query"`
}

type Population struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Features []struct {
		Type       string `json:"type"`
		Properties struct {
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
		} `json:"properties"`
		Geometry struct {
			Type        string          `json:"type,omitempty"`
			Coordinates [][][][]float64 `json:"coordinates,omitempty"`
		} `json:"geometry"`
	} `json:"features"`
}

func (p *Population) Value() (driver.Value, error) {
	return json.Marshal(p)
}
func (p *Population) Scan(val interface{}) error {
	b, ok := val.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &p)
}
