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

type camionPosicion struct {
	camion Camion
	posicion int
}

func ordenarCamiones(camiones []camionPosicion, ascendente bool) {
	sort.Slice(camiones, func(i, j int) bool {
		volumenI := camiones[i].camion.volumen_cm3
		volumenJ := camiones[j].camion.volumen_cm3
		masaI := camiones[i].camion.mma
		masaJ := camiones[j].camion.mma

		if ascendente {
			return volumenI < volumenJ && masaI < masaJ
		} else {
			return volumenI > volumenJ && masaI > masaJ
		}
	})
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

		camionesMismoTipo := []camionPosicion{}
		posicion := 0
		for _, camion := range *camionesDisponibles {
			if camion.tipo == suministro.tipo {
				camionesMismoTipo = append(camionesMismoTipo, camionPosicion{camion, posicion})
			}
			posicion++
		}

		ordenarCamiones(camionesMismoTipo, true)

		for _, c := range camionesMismoTipo {
			camion_array := []Camion{c.camion}
			if PuedeTransportarSuministro(camion_array, suministro) {
				*camionesAsignados = append(*camionesAsignados, c.camion)
				*camionesDisponibles = EliminarElemento(*camionesDisponibles, c.posicion)
				return
			}
		}

		ordenarCamiones(camionesMismoTipo, false)

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