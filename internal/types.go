package internal

type Population struct {
	Type     string    `json:"type"`
	Name     string    `json:"name"`
	Features []Feature `json:"features"`
}

type Feature struct {
	Type       string `json:"type" bson:"type,omitempty"`
	Properties struct {
		GmlId                    string  `json:"gmlId,omitempty" bson:"gml_id,omitempty"`
		LocalId                  string  `json:"localId,omitempty" bson:"localId,omitempty"`
		Namespace                string  `json:"namespace,omitempty" bson:"namespace,omitempty"`
		VersionId                *string `json:"versionId,omitempty" bson:"versionId,omitempty"`
		LocalisedCharacterString string  `json:"localisedCharacterString,omitempty" bson:"LocalisedCharacterString,omitempty"`
		MeasurementUnitUom       *string `json:"measurementUnitUom,omitempty" bson:"measurementUnit_uom,omitempty"`
		NotCountedProportion     *string `json:"notCountedProportion,omitempty" bson:"notCountedProportion,omitempty"`
		BeginPosition            string  `json:"beginPosition,omitempty" bson:"beginPosition,omitempty"`
		EndPosition              string  `json:"endPosition,omitempty" bson:"endPosition,omitempty"`
		Duration                 string  `json:"duration,omitempty" bson:"duration,omitempty"`
	} `json:"properties" bson:"properties,omitempty"`
	Geometry struct {
		Type        string          `json:"type,omitempty" bson:"type,omitempty"`
		Coordinates [][][][]float64 `json:"coordinates,omitempty" bson:"coordinates,omitempty"`
	} `json:"geometry"`
}
