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

	fmt.Printf("JSON : %d stations, %d observations\n", len(stationsJSON), compterObs(stationsJSON))

}

func compterObs(stations []Station) int {
	total := 0
	for _, s := range stations {
		total += len(s.Observations)
	}
	return total
}
