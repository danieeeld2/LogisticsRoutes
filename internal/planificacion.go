package interal 

type Planificacion struct {
	camiones		[]Camion
}

func (p Planificacion) getCamion(i uint) Camion {
	return p.camiones[i]
}