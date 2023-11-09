package internal

import (
	"testing"
)

func PuedeTransportarSuministro(camiones []Camion, suministro Suministros) bool {
	for _, camion := range camiones {
		if camion.tipo != suministro.tipo {
			return false
		}
	}
	
	uint suma_mma = 0
	float32 volumen_total = 0
	float32 volumen_suministro = suministro.dimensiones_whd_cm[0]*suministro.dimensiones_whd_cm[1]*suministro.dimensiones_whd_cm[2]
	for _, camion := range camiones {
		suma_mma += camion.mma
		volumen_total += camion.dimensiones_whd_cm[0]*camion.dimensiones_whd_cm[1]*camion.dimensiones_whd_cm[2]
	}
	
	if suma_mma < suministro.mma || volumen_total < volumen_suministro {
		return false
	}
	
	return true
}

func TestPlanificacion(t *testing.T) {
		
}
