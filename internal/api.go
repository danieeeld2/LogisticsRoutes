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

// Aquí leería del fichero de configuración una conexión a la base de datos
// Creamos una de prueba

var bd BD = getBDPrueba()

func getCamiom(w http.ResponseWriter, r *http.Request) {
	matricula := chi.URLParam(r, "matricula")
	camion, ok := getCamionMatricula(matricula, bd)

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
	
}

func deleteCamion(w http.ResponseWriter, r *http.Request) {
	
}

func getAsignacion(w http.ResponseWriter, r *http.Request) {
	
}

func postAsignacion(w http.ResponseWriter, r *http.Request) {
	
}

func getSuministro(w http.ResponseWriter, r *http.Request) {
	
}

func putSuministro(w http.ResponseWriter, r *http.Request) {
	
}

func deleteSuministro(w http.ResponseWriter, r *http.Request) {	
	
}	