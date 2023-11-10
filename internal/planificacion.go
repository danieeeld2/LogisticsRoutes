package internal 

import (
	"sort"
)

type Planificacion struct {
	camiones		[]Camion
	suministros 		[]Suministro
	asignacion		map[Camion][]Suministro
}

func EliminarElemento(slice []Camion, indice int) []Camion {
	return append(slice[:indice], slice[indice+1:]...)
}

func AsigarCamiones(CamionesDisponibles *[]Camion, suministro Suministro, CamionesAsignados *[]Camion) {
		if len(*CamionesDisponibles) == 0 {
			return
		}

		type camionPosicion struct {
			camion Camion
			posicion int
		}

		camionesMismoTipo := []camionPosicion{}
		posicion := 0
		for _, camion := range *CamionesDisponibles {
			if camion.tipo == suministro.tipo {
				camionesMismoTipo = append(camionesMismoTipo, camionPosicion{camion, posicion})
			}
			posicion++
		}

		sort.Slice(camionesMismoTipo, func(i, j int) bool {
			volumenI := calcularVolumen(camionesMismoTipo[i].camion)
			volumenJ := calcularVolumen(camionesMismoTipo[j].camion)
			masaI := camionesMismoTipo[i].camion.mma
			masaJ := camionesMismoTipo[j].camion.mma
			return volumenI < volumenJ && masaI < masaJ
		})

		for _, c := range camionesMismoTipo {
			camion_array := []Camion{c.camion}
			if PuedeTransportarSuministro(camion_array, suministro) {
				*CamionesAsignados = append(*CamionesAsignados, c.camion)
				*CamionesDisponibles = EliminarElemento(*CamionesDisponibles, c.posicion)
				return
			}
		}

		sort.Slice(camionesMismoTipo, func(i, j int) bool {
			volumenI := calcularVolumen(camionesMismoTipo[i].camion)
			volumenJ := calcularVolumen(camionesMismoTipo[j].camion)
			masaI := camionesMismoTipo[i].camion.mma
			masaJ := camionesMismoTipo[j].camion.mma
			return volumenI > volumenJ && masaI > masaJ
		})

		copia_temporal_disponibles := *CamionesDisponibles
		var asignados []Camion

		for _, c := range camionesMismoTipo {
			if PuedeTransportarSuministro(asignados, suministro) {
				*CamionesAsignados = append(*CamionesAsignados, asignados...)
				return
			}
			asignados = append(asignados, c.camion)
			copia_temporal_disponibles = EliminarElemento(*CamionesDisponibles, c.posicion)
		}

		*CamionesDisponibles = copia_temporal_disponibles

		return		
}