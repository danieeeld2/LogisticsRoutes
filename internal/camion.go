package internal

type Camion struct {
	tipo      				TipoSuministro
	dimensiones_whd_cm		[3]float32
	mma       				uint
}

func NewCamion(tipo TipoSuministro, dimensiones [3]float32, mma uint) Camion {
	return Camion{
		tipo:			tipo,
		dimensiones_whd_cm:	dimensiones,
		mma:			mma,
	}
}

func calcularVolumen(camion Camion) float32 {
	return camion.dimensiones_whd_cm[0] * camion.dimensiones_whd_cm[1] * camion.dimensiones_whd_cm[2]
}
