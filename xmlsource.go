package main

import (
	"encoding/xml"
	"os"
	"strconv"
)

type xmlRoot struct {
	XMLName  xml.Name     `xml:"weather_dataset"`
	Stations []xmlStation `xml:"station"`
}

type xmlStation struct {
	Country      string           `xml:"country,attr"`
	Coordinates  xmlCoordinates   `xml:"coordinates"`
	Hardware     xmlHardware      `xml:"hardware"`
	Observations []xmlObservation `xml:"observations>observation"`
}

type xmlCoordinates struct {
	Lat      float64 `xml:"lat,attr"`
	Lon      float64 `xml:"lon,attr"`
	Altitude int     `xml:"altitude,attr"`
}

type xmlHardware struct {
	Model string `xml:"model,attr"`
}

type xmlObservation struct {
	Sky      string       `xml:"sky,attr"`
	Measures []xmlMeasure `xml:"measure"`
	Wind     xmlWind      `xml:"wind"`
	Note     *string      `xml:"note"`
}

type xmlMeasure struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

type xmlWind struct {
	Speed     float64 `xml:"speed,attr"`
	Direction int     `xml:"direction,attr"`
}

func LoadFromXML(path string) ([]Station, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var root xmlRoot
	err = xml.Unmarshal(data, &root)
	if err != nil {
		return nil, err
	}

	var stations []Station
	for _, xs := range root.Stations {
		station := convertXMLStation(xs)
		stations = append(stations, station)
	}

	return stations, nil
}

func convertXMLStation(xs xmlStation) Station {
	var observations []Observation
	for _, xo := range xs.Observations {
		observations = append(observations, convertXMLObservation(xo))
	}

	return Station{
		Country:      xs.Country,
		Altitude:     xs.Coordinates.Altitude,
		Location:     Coordinates{Latitude: xs.Coordinates.Lat, Longitude: xs.Coordinates.Lon},
		DeviceModel:  xs.Hardware.Model,
		Observations: observations,
	}
}

func convertXMLObservation(xo xmlObservation) Observation {
	var temperature float64
	for _, m := range xo.Measures {
		if m.Type == "temperature" {
			temperature, _ = strconv.ParseFloat(m.Value, 64)
		}
	}

	return Observation{
		Temperature: temperature,
		Conditions:  xo.Sky,
		Wind:        Wind{Speed: xo.Wind.Speed, DirectionDeg: xo.Wind.Direction},
		Note:        xo.Note,
	}
}
