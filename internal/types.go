package internal

type Db struct {
	Population []byte
}

type Population struct {
	Type string `json:"type"`
	Name string `json:"pop-distribution"`
	Crs  struct {
		Type       string `json:"type"`
		Properties struct {
			Name string `jsonm:"name"`
		} `json:"properties"`
	} `json:"crs"`
	Features []struct {
		Type       string `json:"type"`
		Properties struct {
			GmlId                    string `json:"gml_id"`
			LocalId                  string `json:"localId"`
			Namespace                string `json:"namespace"`
			VersionId                string `json:"versionId"`
			LocalisedCharacterString string `json:"LocalisedCharacterString"`
			MeasurementUnitUom       string `json:"measurementUnit_uom"`
			NotCountedProportion     string `json:"notCountedProportion"`
			BeginPosition            string `json:"beginPosition"`
			EndPosition              string `json:"endPosition"`
			Duration                 string `json:"duration"`
		} `json:"properties"`
		Geometry struct {
			Type        string          `json:"type"`
			Coordiantes [][][][]float64 `json:"coordinates"`
		} `json:"geometry"`
	} `json:"features"`
}
