package internal

type Camion struct {
	tipo      				TipoSuministro
	volumen_cm3				float32
	mma       				uint
}

func NewCamion(tipo TipoSuministro, volumen float32, mma uint) Camion {
	return Camion{
		tipo:			tipo,
		volumen_cm3:	volumen,
		mma:			mma,
	}
}
