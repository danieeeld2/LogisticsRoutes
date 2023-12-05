package internal

type Camion struct {
	tipo      				TipoSuministro
	volumen_cm3				float32
	mma       				uint
}

func NewCamion(tipo TipoSuministro, volumen float32, mma uint) Camion {
	GetLogger().Info().Msg("Creando camion de prueba")
	return Camion{
		tipo:			tipo,
		volumen_cm3:	volumen,
		mma:			mma,
	}
}
