package main

import (
	"encoding/json"
	"net/http"
)

type App struct {
	store *Store
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

func (a *App) listStations(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, a.store.All())
}

func (a *App) getStation(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	st, ok := a.store.Get(id)
	if !ok {
		writeError(w, http.StatusNotFound, "station introuvable")
		return
	}
	writeJSON(w, http.StatusOK, st)
}

func (a *App) createStation(w http.ResponseWriter, r *http.Request) {
	var st Station
	err := json.NewDecoder(r.Body).Decode(&st)
	if err != nil {
		writeError(w, http.StatusBadRequest, "JSON invalide")
		return
	}
	if st.ID == "" {
		writeError(w, http.StatusBadRequest, "id manquant")
		return
	}
	if a.store.Has(st.ID) {
		writeError(w, http.StatusConflict, "id déjà utilisé")
		return
	}
	a.store.Put(st)
	writeJSON(w, http.StatusCreated, st)
}

func (a *App) updateStation(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var st Station
	err := json.NewDecoder(r.Body).Decode(&st)
	if err != nil {
		writeError(w, http.StatusBadRequest, "JSON invalide")
		return
	}

	st.ID = id
	existe := a.store.Has(id)
	a.store.Put(st)

	if existe {
		writeJSON(w, http.StatusOK, st)
	} else {
		writeJSON(w, http.StatusCreated, st)
	}
}
