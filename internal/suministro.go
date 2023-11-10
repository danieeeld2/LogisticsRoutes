package internal

import "time"

type Suministro struct {
	direccion				string
	tiempo					time.Duration
	peso_kg					float32
	dimensiones_whd_cm		[3]float32
	tipo					TipoSuministro
}

func NuevoSuministro(direccion string, tiempo time.Duration, peso_kg float32, dimensiones [3]float32, tipo TipoSuministro) Suministro {
	return Suministro{
		direccion:           direccion,
		tiempo:              tiempo,
		peso_kg:             peso_kg,
		dimensiones_whd_cm: dimensiones,
		tipo:                tipo,
	}
}
