package internal

import "time"

type TipoSuministro string

const (
	NORMAL       TipoSuministro = "normal"
	FRIGORIFICO  				= "frigorifico"
	QUIMICO 					= "quimico"
)

type Suministro struct {
	direccion	string
	tiempo		time.Duration
	peso		float32
	dimensiones	[3]float32
	tipo		TipoSuministro
}