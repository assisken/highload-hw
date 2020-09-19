package app

type Forecast struct {
	City        string 	`json:"city"`
	Unit        string 	`json:"unit"`
	Temperature float64	`json:"temperature"`
}
