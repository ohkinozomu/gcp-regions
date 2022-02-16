package regions

type Region struct {
	Name      string  `json:"name"`
	Flag      string  `json:"flag"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
