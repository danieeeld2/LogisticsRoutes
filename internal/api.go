package internal

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"encoding/json"
)

var router *chi.Mux

func createRouter() *chi.Mux{
	if router == nil {
		router = chi.NewRouter()
		router.Get("/camion/{matricula}", getCamiom)
		router.Put("/camion/{matricula}", putCamion)
		router.Delete("/camion/{matricula}", deleteCamion)
		router.Get("/suministro/{id}/asignacion", getAsignacion)
		router.Post("/suministro/{id}/asignacion", postAsignacion)
		router.Get("/suministro/{id}", getSuministro)
		router.Put("/suministro/{id}", putSuministro)
		router.Delete("/suministro/{id}", deleteSuministro)
	}

	return router
}

func getCamiom(w http.ResponseWriter, r *http.Request) {
	matricula := chi.URLParam(r, "matricula")
	camion, ok := getCamionMatricula(matricula, getBDPrueba())

	if !ok {
		http.Error(w, "Camion no encontrado", http.StatusNotFound)
		return
	}

	jsonResponse, err := json.Marshal(camion)
	if err != nil {
		http.Error(w, "Error al convertir a JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func putCamion(w http.ResponseWriter, r *http.Request) {
	matricula := chi.URLParam(r, "matricula")

	var nuevoCamion Camion
	err := json.NewDecoder(r.Body).Decode(&nuevoCamion)
	if err != nil {
		http.Error(w, "Error al leer el body", http.StatusBadRequest)
		return
	}

	camionActualizado, ok := putCamionMatricula(matricula, nuevoCamion, getBDPrueba())
	if !ok {
		http.Error(w, "Camion no encontrado", http.StatusNotFound)
		return
	}

	jsonResponse, err := json.Marshal(camionActualizado)
	if err != nil {
		http.Error(w, "Error al convertir a JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func deleteCamion(w http.ResponseWriter, r *http.Request) {
	matricula := chi.URLParam(r, "matricula")
	camion, ok := deleteCamionMatricula(matricula, getBDPrueba())

	if !ok {
		http.Error(w, "Camion no encontrado", http.StatusNotFound)
		return
	}

	jsonResponse, err := json.Marshal(camion)
	if err != nil {
		http.Error(w, "Error al convertir a JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func getAsignacion(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	asignacion, ok := getSuministroIDAsignacion(id, getBDPrueba())

	if !ok {
		http.Error(w, "Asignacion no encontrada", http.StatusNotFound)
		return
	}

	jsonResponse, err := json.Marshal(asignacion)
	if err != nil {
		http.Error(w, "Error al convertir a JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func postAsignacion(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var nuevaAsignacion []string
	err := json.NewDecoder(r.Body).Decode(&nuevaAsignacion)
	if err != nil {
		http.Error(w, "Error al leer el body", http.StatusBadRequest)
		return
	}

	asignacionCreada, ok := postSuministroIDAsignacion(id, nuevaAsignacion, getBDPrueba())
	if !ok {
		http.Error(w, "Asignacion no creada", http.StatusNotFound)
		return
	}

	jsonResponse, err := json.Marshal(asignacionCreada)
	if err != nil {
		http.Error(w, "Error al convertir a JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func getSuministro(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	suministro, ok := getSuministroID(id, getBDPrueba())

	if !ok {
		http.Error(w, "Suministro no encontrado", http.StatusNotFound)
		return
	}

	jsonResponse, err := json.Marshal(suministro)
	if err != nil {
		http.Error(w, "Error al convertir a JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func putSuministro(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var nuevoSuministro Suministro
	err := json.NewDecoder(r.Body).Decode(&nuevoSuministro)
	if err != nil {
		http.Error(w, "Error al leer el body", http.StatusBadRequest)
		return
	}

	suministroActualizado, ok := putSuministroID(id, nuevoSuministro, getBDPrueba())
	if !ok {
		http.Error(w, "Suministro no encontrado", http.StatusNotFound)
		return
	}

	jsonResponse, err := json.Marshal(suministroActualizado)
	if err != nil {
		http.Error(w, "Error al convertir a JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func deleteSuministro(w http.ResponseWriter, r *http.Request) {	
	id := chi.URLParam(r, "id")
	suministro, ok := deleteSuministroID(id, getBDPrueba())

	if !ok {
		http.Error(w, "Suministro no encontrado", http.StatusNotFound)
		return
	}
	
	jsonResponse, err := json.Marshal(suministro)
	if err != nil {
		http.Error(w, "Error al convertir a JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}	