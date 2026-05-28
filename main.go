package main

import (
	"log"
	"net/http"
)

func main() {
	stations, err := LoadFromJSON("weather_data.json")
	if err != nil {
		log.Fatal(err)
	}

	store := NewStore()
	for _, s := range stations {
		store.Put(s)
	}
	log.Printf("bootstrap : %d stations", len(stations))

	mux := http.NewServeMux()
	http.ListenAndServe(":8080", mux)
}
