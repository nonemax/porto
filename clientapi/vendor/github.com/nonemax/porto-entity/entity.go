package entity

type Port struct {
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float32 `json:"coordinates"`
	Province    string    `json:"province"`
	TimeZone    string    `json:"timezone"`
	Unlocks     []string  `json:"unlocs"`
	Code        string    `json:"code"`
}
