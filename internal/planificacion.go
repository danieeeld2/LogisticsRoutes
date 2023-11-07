package interal 

type Planificacion struct {
	camiones		[]Camion
	suministros 	[]Suministro
	asignacion		map[camiones][]Suministro
}