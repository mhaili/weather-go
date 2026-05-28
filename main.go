package main

import (
	"fmt"
	"log"
)

func main() {
	stations, err := LoadFromJSON("weather_data.json")
	if err != nil {
		log.Fatal(err)
	}

	stationsXML, err := LoadFromXML("weather_data.xml")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Stations :", len(stations))
	fmt.Println("Obs :", len(stations[0].Observations))
	fmt.Println("XML  :", len(stationsXML), "stations")

	s := stations[0]
	fmt.Println(s.Country, s.Altitude, s.Location, s.DeviceModel)

	x := stationsXML[0]
	fmt.Println("\n--- XML ---")
	fmt.Println(x.Country, x.Altitude, x.Location, x.DeviceModel)
	fmt.Println(x.Observations[0].Temperature, x.Observations[0].Conditions, x.Observations[0].Wind, x.Observations[0].Note)

	o := s.Observations[0]
	fmt.Println(o.Temperature, o.Conditions, o.Wind, o.Note)

	_, err = LoadFromXML("weather_data.xml")
	if err != nil {
		log.Fatal(err)
	}
}
