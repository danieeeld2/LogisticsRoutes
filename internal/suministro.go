package internal

import "time"

type Suministro struct {
	direccion				string			`json:"direccion"`
	tiempo					time.Duration	`json:"tiempo"`
	peso_kg					float32			`json:"peso_kg"`
	volumen_cm3				float32			`json:"volumen_cm3"`
	tipo					TipoSuministro	`json:"tipo"`
}

func NuevoSuministro(direccion string, tiempo time.Duration, peso_kg float32, volumen float32, tipo TipoSuministro) Suministro {
	GetLogger().Info().Msg("Creando suministro de prueba")
	return Suministro{
		direccion:           direccion,
		tiempo:              tiempo,
		peso_kg:             peso_kg,
		volumen_cm3:         volumen,
		tipo:                tipo,
	}
}
