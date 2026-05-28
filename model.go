package main

// station météo
type Station struct {
	Country      string
	Altitude     int
	Location     Coordinates
	DeviceModel  string
	Observations []Observation
}

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

// une mesure à un t donné
type Observation struct {
	Temperature float64
	Conditions  string
	Wind        Wind
	Note        *string
}
type Wind struct {
	Speed        float64
	DirectionDeg int
}
