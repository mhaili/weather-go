package main

import (
	"encoding/json"
	"net/http"
)

type App struct {
	store *Store
}

type ErrorResponse struct {
	Error string `json:"error"`
	Code  string `json:"code,omitempty"`
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

//func writeError(w http.ResponseWriter, status int, msg string) {
//	writeJSON(w, status, map[string]string{"error": msg})
//}

func writeError(w http.ResponseWriter, status int, code, msg string) {
	writeJSON(w, status, ErrorResponse{Error: msg, Code: code})
}

func (a *App) listStations(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, a.store.All())
}

func (a *App) getStation(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	st, ok := a.store.Get(id)
	if !ok {
		//writeError(w, http.StatusNotFound, "station introuvable")
		writeError(w, http.StatusNotFound, "NOT_FOUND", "station introuvable")
		return
	}
	writeJSON(w, http.StatusOK, st)
}

func (a *App) createStation(w http.ResponseWriter, r *http.Request) {
	var st Station
	err := json.NewDecoder(r.Body).Decode(&st)
	if err != nil {
		//writeError(w, http.StatusBadRequest, "JSON invalide")
		writeError(w, http.StatusBadRequest, "BAD_JSON", "JSON invalide")
		return
	}
	if st.ID == "" {
		//writeError(w, http.StatusBadRequest, "id manquant")
		writeError(w, http.StatusBadRequest, "MISSING_ID", "id manquant")
		return
	}
	if a.store.Has(st.ID) {
		//writeError(w, http.StatusConflict, "id déjà utilisé")
		writeError(w, http.StatusConflict, "ID_TAKEN", "id déjà utilisé")
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
		//writeError(w, http.StatusBadRequest, "JSON invalide")
		writeError(w, http.StatusBadRequest, "BAD_JSON", "JSON invalide")
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

func (a *App) deleteStation(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if !a.store.Delete(id) {
		writeError(w, http.StatusNotFound, "NOT_FOUND", "station introuvable")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (a *App) listObservations(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	st, ok := a.store.Get(id)
	if !ok {
		writeError(w, http.StatusNotFound, "NOT_FOUND", "station introuvable")
		return
	}
	writeJSON(w, http.StatusOK, st.Observations)
}
