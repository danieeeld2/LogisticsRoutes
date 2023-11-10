package internal 

type Planificacion struct {
	camiones		[]Camion
	suministros 		[]Suministro
	asignacion		map[Camion][]Suministro
}


func AsigarCamiones(CamionesDisponibles *[]Camion, suministro Suministro, CamionesAsignados *[]Camion) {
	
}// ALGORITMO DE ASIGNACION
// FILEPATH: /home/daniel/Git/LogisticsRoutes/internal/planificacion.go

func AsignarCamiones(CamionesDisponibles *[]Camion, suministro Suministro, CamionesAsignados *[]Camion) {
	// Si no hay camiones disponibles, terminamos
	if len(*CamionesDisponibles) == 0 {
		return
	}

	// Buscamos si hay algún camión individual que pueda cargar el suministro
	camionIndividual := BuscarCamionIndividual(CamionesDisponibles, suministro)
	if camionIndividual != nil {
		*CamionesAsignados = append(*CamionesAsignados, *camionIndividual)
		eliminarCamion(CamionesDisponibles, camionIndividual)
		return
	}

	// Si no hay camiones individuales, buscamos en los subconjuntos de camiones disponibles
	for i := 1; i <= len(*CamionesDisponibles); i++ {
		subconjuntos := Combinaciones(*CamionesDisponibles, i)
		for _, subconjunto := range subconjuntos {
			if PuedeCargar(subconjunto, suministro) {
				*CamionesAsignados = append(*CamionesAsignados, subconjunto...)
				eliminarCamiones(CamionesDisponibles, subconjunto)
				return
			}
		}
	}
}
