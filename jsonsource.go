package main

import (
	"encoding/json"
	"os"
)

type jsonRoot struct {
	Stations []jsonStation `json:"stations"`
}

var paysVersISO = map[string]string{
	"France":    "FR",
	"Espagne":   "ES",
	"Portugal":  "PT",
	"Italie":    "IT",
	"Allemagne": "DE",
	"Belgique":  "BE",
	"Pays-Bas":  "NL",
	"Autriche":  "AT",
	"Suisse":    "CH",
	"Danemark":  "DK",
	"Suède":     "SE",
	"Norvège":   "NO",
	"Pologne":   "PL",
	"Tchéquie":  "CZ",
}

type jsonStation struct {
	ID           string            `json:"id"`
	Country      string            `json:"country"`
	AltitudeM    int               `json:"altitude_m"`
	Location     jsonLocation      `json:"location"`
	Device       jsonDevice        `json:"device"`
	Observations []jsonObservation `json:"observations"`
}

type jsonLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type jsonDevice struct {
	Type        string `json:"type"`
	InstalledOn string `json:"installed_on"`
}

type jsonObservation struct {
	Temperature float64  `json:"temperature_celsius"`
	Wind        jsonWind `json:"wind"`
	Conditions  string   `json:"conditions"`
	Notes       *string  `json:"notes"`
}

type jsonWind struct {
	Speed     float64 `json:"speed_kmh"`
	Direction int     `json:"direction_deg"`
}

func convertJSONObservation(jo jsonObservation) Observation {
	return Observation{
		Temperature: jo.Temperature,
		Conditions:  jo.Conditions,
		Wind:        Wind{Speed: jo.Wind.Speed, DirectionDeg: jo.Wind.Direction},
		Note:        jo.Notes,
	}
}

func convertJSONStation(js jsonStation) Station {

	//en Maj avant le map
	//iso := paysVersISO[strings.ToUpper(js.Country)]
	iso := paysVersISO[js.Country]

	var observations []Observation
	for _, jo := range js.Observations {
		observations = append(observations, convertJSONObservation(jo))
	}

	return Station{
		ID:           js.ID,
		Country:      iso,
		Altitude:     js.AltitudeM,
		Location:     Coordinates{Latitude: js.Location.Latitude, Longitude: js.Location.Longitude},
		DeviceModel:  js.Device.Type,
		Observations: observations,
	}
}

func LoadFromJSON(path string) ([]Station, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var root jsonRoot
	if err := json.Unmarshal(data, &root); err != nil {
		return nil, err
	}

	var stations []Station
	for _, js := range root.Stations {
		stations = append(stations, convertJSONStation(js))
	}
	return stations, nil
}
