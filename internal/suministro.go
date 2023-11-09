package internal

import "time"

type Suministro struct {
	direccion				string
	tiempo					time.Duration
	peso_kg				float32
	dimensiones_whd_cm			[3]float32
	tipo					TipoSuministro
}
