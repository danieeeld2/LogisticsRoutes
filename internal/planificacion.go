package internal 

type Planificacion struct {
	camiones		[]Camion
	suministros 		[]Suministro
	asignacion		map[Camion][]Suministro
}


func AsigarCamiones(CamionesDisponibles *[]Camion, suministro Suministro, CamionesAsignados *[]Camion) {
}