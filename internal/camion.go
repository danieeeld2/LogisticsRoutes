package internal

type tipoCamion string

const (
	NORMAL       tipoCamion = "normal"
	FRIGORIFICO  			= "frigorifico"
	QUIMICO 				= "quimico"
)

type Camion struct {
	tipo      		tipoCamion
	capacidad 		[3]float32
	mma       		uint
	suministros		[]Suministro
}

func (c Camion) getSuministro(i uint) Suministro {
	return c.suministros[i]
}
