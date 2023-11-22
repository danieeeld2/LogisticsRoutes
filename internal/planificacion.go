package internal 

import (
	"sort"
)

type Planificacion struct {
	camiones		[]Camion
	suministros 	[]Suministro
	asignacion		map[Camion][]Suministro
}

func EliminarElemento(slice []Camion, indice int) []Camion {
	return append(slice[:indice], slice[indice+1:]...)
}

func PuedeTransportarSuministro(camiones []Camion, suministro Suministro) bool {
	for _, camion := range camiones {
		if camion.tipo != suministro.tipo {
			return false
		}
	}
	
	var suma_mma uint
	var volumen_total float32
	
	suma_mma = 0 
	volumen_total = 0 
	for _, camion := range camiones {
		suma_mma += camion.mma
		volumen_total += camion.volumen_cm3
	}
	
	if float32(suma_mma) <= suministro.peso_kg || volumen_total <= suministro.volumen_cm3 {
		return false
	}
	
	return true
}

func AsigarCamiones(camionesDisponibles *[]Camion, suministro Suministro, camionesAsignados *[]Camion) {
		if len(*camionesDisponibles) == 0 {
			return
		}

		type camionPosicion struct {
			camion Camion
			posicion int
		}

		camionesMismoTipo := []camionPosicion{}
		posicion := 0
		for _, camion := range *camionesDisponibles {
			if camion.tipo == suministro.tipo {
				camionesMismoTipo = append(camionesMismoTipo, camionPosicion{camion, posicion})
			}
			posicion++
		}

		sort.Slice(camionesMismoTipo, func(i, j int) bool {
			volumenI := camionesMismoTipo[i].camion.volumen_cm3
			volumenJ := camionesMismoTipo[j].camion.volumen_cm3
			masaI := camionesMismoTipo[i].camion.mma
			masaJ := camionesMismoTipo[j].camion.mma
			return volumenI < volumenJ && masaI < masaJ
		})

		for _, c := range camionesMismoTipo {
			camion_array := []Camion{c.camion}
			if PuedeTransportarSuministro(camion_array, suministro) {
				*camionesAsignados = append(*camionesAsignados, c.camion)
				*camionesDisponibles = EliminarElemento(*camionesDisponibles, c.posicion)
				return
			}
		}

		sort.Slice(camionesMismoTipo, func(i, j int) bool {
			volumenI := camionesMismoTipo[i].camion.volumen_cm3
			volumenJ := camionesMismoTipo[j].camion.volumen_cm3
			masaI := camionesMismoTipo[i].camion.mma
			masaJ := camionesMismoTipo[j].camion.mma
			return volumenI > volumenJ && masaI > masaJ
		})

		copia_temporal_disponibles := *camionesDisponibles
		var asignados []Camion

		for _, c := range camionesMismoTipo {
			if PuedeTransportarSuministro(asignados, suministro) {
				*camionesAsignados = append(*camionesAsignados, asignados...)
				return
			}
			asignados = append(asignados, c.camion)
			copia_temporal_disponibles = EliminarElemento(*camionesDisponibles, c.posicion)
		}

		*camionesDisponibles = copia_temporal_disponibles

		return		
}