package main

import (
	"fmt"
	"log"
)

func main() {
	stationsJSON, err := LoadFromJSON("weather_data.json")
	if err != nil {
		log.Fatal(err)
	}

	stationsXML, err := LoadFromXML("weather_data.xml")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("JSON : %d stations, %d observations\n", len(stationsJSON), compterObs(stationsJSON))
	fmt.Printf("XML  : %d stations, %d observations\n", len(stationsXML), compterObs(stationsXML))

	if len(stationsJSON) == len(stationsXML) && compterObs(stationsJSON) == compterObs(stationsXML) {
		fmt.Println("Cohérence : OK")
	} else {
		fmt.Println("Cohérence : ÉCART")
	}

	_, rafale := MaxWindGust(stationsJSON)
	fmt.Printf("Rafale la plus forte : %.1f km/h\n", rafale)

	fmt.Printf("Temp. moyenne Bordeaux : %.1f °C\n", AvgTemperature(stationsJSON[0]))

	fmt.Println("Stations par pays :", CountByCountry(stationsJSON))

	fmt.Println("Stations FR :", len(FilterByCountry(stationsJSON, "FR")))
}

func compterObs(stations []Station) int {
	total := 0
	for _, s := range stations {
		total += len(s.Observations)
	}
	return total
}
