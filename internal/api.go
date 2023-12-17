package internal

import (
	"net/http"
	"github.com/go-chi/chi/v5"
)

var router *chi.Mux

func createRouter() *chi.Mux{
	if router == nil {
		router = chi.NewRouter()
		router.Get("/config", getConfig)
		router.Get("/camiones", getCamiones)
		router.Post("/camiones", createCamion)
		router.Get("/suministros", getSuministros)
		router.Post("/suministros", createSuministro)
		router.Post("/planificacion", planificarEntrega)
		router.Get("/planificacion", getEntrega)
	}

	return router
}

func getConfig(w http.ResponseWriter, r *http.Request){

}

func getCamiones(w http.ResponseWriter, r *http.Request){

}

func createCamion(w http.ResponseWriter, r *http.Request){
	
}

func getSuministros(w http.ResponseWriter, r *http.Request){
	
}

func createSuministro(w http.ResponseWriter, r *http.Request){
	
}

func planificarEntrega(w http.ResponseWriter, r *http.Request){
	
}

func getEntrega(w http.ResponseWriter, r *http.Request){
	
}