package catalog

//App an interpretation of features an metrics to describe and measure the app
type App struct {
	Global   Feature   `json:"global"`   // apps globals
	Features []Feature `json:"features"` // apps features and options
	Metrics  Metrics   `json:"metrics"`  // apps metrics
}

//Feature a generic feature object to describe functionality of apps
type Feature struct {
	Key           string `json:"key"`           // key string of the feature to provide descrptional context
	Label         string `json:"label"`         // label of the feature
	Description   string `json:"description"`   // description
	Prerequisites []Task `json:"prerequisites"` // prerequisites to do before implemention
}

//Task a task to do
type Task struct {
	Key            string        `json:"key"`            // key string of the feature to provide descrptional context
	Label          string        `json:"label"`          // label of the feature
	Description    string        `json:"description"`    // description
	Implementation DurationalKPI `json:"implementation"` // kpi for average consumption of time and personnel to implement this
}

//Metrics apps metrics to provide kpi's for work and financial measuring
type Metrics struct {
	Creation       DurationalKPI `json:"creation"`       // kpi about time consumed to build the first version
	Implementation DurationalKPI `json:"implementation"` // kpi for average consumption of time and personnel to implement this
}

//DurationalKPI a kpi providing dration realtet values
type DurationalKPI struct {
	Duration   uint64 `json:"duration"`         // duration in minutes
	Developers uint   `json:"developers_count"` // number of developers
}
