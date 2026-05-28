package main

func FilterByCountry(stations []Station, iso string) []Station {
	var resultat []Station
	for _, s := range stations {
		if s.Country == iso {
			resultat = append(resultat, s)
		}
	}
	return resultat
}

func AvgTemperature(s Station) float64 {
	if len(s.Observations) == 0 {
		return 0
	}
	var total float64
	for _, o := range s.Observations {
		total += o.Temperature
	}
	return total / float64(len(s.Observations))
}

func MaxWindGust(stations []Station) (Station, float64) {
	var stationMax Station
	var maxVent float64
	for _, s := range stations {
		for _, o := range s.Observations {
			if o.Wind.Speed > maxVent {
				maxVent = o.Wind.Speed
				stationMax = s
			}
		}
	}
	return stationMax, maxVent
}

func CountByCountry(stations []Station) map[string]int {
	comptage := make(map[string]int)
	for _, s := range stations {
		comptage[s.Country]++
	}
	return comptage
}
