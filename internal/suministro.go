package internal

import "time"

type TipoSuministro string

const (
	NORMAL       TipoSuministro = "normal"
	FRIGORIFICO  				= "frigorifico"
	QUIMICO 					= "quimico"
)

type Suministro struct {
	direccion				string
	tiempo					time.Duration
	peso_kg					float32
	dimensiones_whd_cm		[3]float32
	tipo					TipoSuministro
}