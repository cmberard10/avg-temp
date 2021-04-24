package models

import "time"

//City object returned from https://public.opendatasoft.com/api/records/1.0/search/?dataset=geonames-all-cities-with-a-population-1000&q=&refine.country=United+States&sort=population&rows=100
type City struct {
	NHits      int `json:"nhits"`
	Parameters struct {
		Dataset string `json:"dataset"`
		Refine  struct {
			Country string `json:"country"`
		} `json:"refine"`
		Timezone string   `json:"timezone"`
		Rows     int      `json:"rows"`
		Start    int      `json:"start"`
		Sort     []string `json:"sort"`
		Format   string   `json:"format"`
	} `json:"parameters"`
	Records     []Record `json:"records"`
	FacetGroups []struct {
		Facets []struct {
			Count int    `json:"count"`
			Path  string `json:"path"`
			State string `json:"state"`
			Name  string `json:"name"`
		} `json:"facets"`
		Name string `json:"name"`
	} `json:"facet_groups"`
}

//Record City Record item
type Record struct {
	DatasetID string `json:"datasetid"`
	RecordID  string `json:"recordid"`
	Fields    struct {
		Elevation        string    `json:"elevation"`
		Name             string    `json:"name"`
		ModificationDate string    `json:"modification_date"`
		Country          string    `json:"country"`
		FeatureClass     string    `json:"feature_class"`
		AlternateNames   string    `json:"alternate_names"`
		FeatureCode      string    `json:"feature_code"`
		Longitude        string    `json:"longitude"`
		GeonameID        string    `json:"geoname_id"`
		Timezone         string    `json:"timezone"`
		Dem              int       `json:"dem"`
		CountryCode      string    `json:"country_code"`
		ASCIIName        string    `json:"ascii_name"`
		Latitude         string    `json:"latitude"`
		Admin1Code       string    `json:"admin1_code"`
		Coordinates      []float64 `json:"coordinates"`
		Population       int       `json:"population"`
	} `json:"fields,omitempty"`
	Geometry struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"geometry"`
	RecordTimestamp time.Time `json:"record_timestamp"`
}
