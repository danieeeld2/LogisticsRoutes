package internal

import (
	"testing"
)

func EsMasChico(camion1 []Camion, camion2 []Camion) bool {
	if len(camion1) > len(camion2) {
		return false
	}
	
	var suma_mma_1, suma_mma_2 uint
	var volumen_total_1, volumen_total_2 float32
	
	suma_mma_1 = 0
	suma_mma_2 = 0
	volumen_total_1 = 0
	volumen_total_2 = 0
	for _, camion := range camion1 {
		suma_mma_1 += camion.mma
		volumen_total_1 += camion.dimensiones_whd_cm[0] * camion.dimensiones_whd_cm[1] * camion.dimensiones_whd_cm[2]
	}
	for _, camion := range camion2 {
		suma_mma_2 += camion.mma
		volumen_total_2 += camion.dimensiones_whd_cm[0] * camion.dimensiones_whd_cm[1] * camion.dimensiones_whd_cm[2]
	}
	
	if suma_mma_2 < suma_mma_1 && volumen_total_2 < volumen_total_1 {
		return false
	}
	
	return true
}

func CalcularConjuntosDeCamiones(camiones []Camion, t int) [][]Camion{
	resultado := [][]Camion{{}}

	for _, elemento := range camiones {
		subconjuntosTemporales := [][]Camion{}

		for _, conjunto := range resultado {
			if len(conjunto) < t {
				nuevoSubconjunto := append([]Camion{}, conjunto...)
				nuevoSubconjunto = append(nuevoSubconjunto, elemento)
				subconjuntosTemporales = append(subconjuntosTemporales, nuevoSubconjunto)
			}
		}

		resultado = append(resultado, subconjuntosTemporales...)
	}

	return resultado
}

func ComprobarAsignacionOptima(CamionesDisponibles []Camion, suministro Suministro, CamionesAsignados []Camion) bool {
	if !PuedeTransportarSuministro(CamionesAsignados, suministro) && len(CamionesAsignados) > 0{
		return false
	}
	
	if len(CamionesDisponibles) == 0 {
		if len(CamionesAsignados) == 0 {
			return true
		}
	}
	
	if len(CamionesAsignados) > 0 {
		conjuntos_camiones_disponibles := CalcularConjuntosDeCamiones(CamionesDisponibles, len(CamionesAsignados))
		for _, conjunto := range conjuntos_camiones_disponibles {
			if PuedeTransportarSuministro(conjunto, suministro) && EsMasChico(conjunto, CamionesAsignados) {
				return false
			}
		}
	}
	
	
	return true
} 

func TestPlanificacion(t *testing.T) {
	t.Log("Comenzando test de planificacion")

	camionesDisponibles := []Camion{}
	CamionesAsignados := []Camion{}
	suministro := NuevoSuministro("Calle Falsa 123", 10, 100, [3]float32{10, 10, 10}, TipoSuministro(NORMAL))

	t.Log("No hay camiones disponibles ni asignados")
	AsigarCamiones(&camionesDisponibles, suministro, &CamionesAsignados)
	if !ComprobarAsignacionOptima(camionesDisponibles, suministro, CamionesAsignados) {
		t.Error("La asignacion no es optima")
	}

	camionesDisponibles = append(camionesDisponibles, NewCamion(TipoSuministro(NORMAL), [3]float32{10, 10, 10}, 100))
	camionesDisponibles = append(camionesDisponibles, NewCamion(TipoSuministro(QUIMICO), [3]float32{100, 100, 100}, 100))
	camionesDisponibles = append(camionesDisponibles, NewCamion(TipoSuministro(NORMAL), [3]float32{1000, 1000, 1000}, 100))
	camionesDisponibles = append(camionesDisponibles, NewCamion(TipoSuministro(NORMAL), [3]float32{1000, 1000, 1000}, 1000))
	CamionesAsignados = []Camion{}
	t.Log("Hay varios camiones disponibles que pueden transportar el vehículo")
	AsigarCamiones(&camionesDisponibles, suministro, &CamionesAsignados)
	if !ComprobarAsignacionOptima(camionesDisponibles, suministro, CamionesAsignados) {
		t.Error("La asignacion no es optima")
	}

	suministro = NuevoSuministro("Calle Falsa 123", 10, 10000, [3]float32{1000, 1000, 1000}, TipoSuministro(NORMAL))
	CamionesAsignados = []Camion{}
	t.Log("Hay varios camiones disponibles pero ninguno puede transportar el suministro")
	AsigarCamiones(&camionesDisponibles, suministro, &CamionesAsignados)
	if !ComprobarAsignacionOptima(camionesDisponibles, suministro, CamionesAsignados) {
		t.Error("La asignacion es optima")
	}

	suministro = NuevoSuministro("Calle Falsa 123", 10, 100, [3]float32{10, 10, 10}, TipoSuministro(QUIMICO))
	CamionesAsignados = []Camion{}
	t.Log("Comprobar que le asigna el de tipo correcto")
	AsigarCamiones(&camionesDisponibles, suministro, &CamionesAsignados)
	if !ComprobarAsignacionOptima(camionesDisponibles, suministro, CamionesAsignados) {
		t.Error("La asignacion no es optima")
	}

	suministro = NuevoSuministro("Calle Falsa 123", 10, 1050, [3]float32{10, 10, 10}, TipoSuministro(NORMAL))
	CamionesAsignados = []Camion{}
	t.Log("Necesita mas de un camión para ser transportado")
	AsigarCamiones(&camionesDisponibles, suministro, &CamionesAsignados)
	if !ComprobarAsignacionOptima(camionesDisponibles, suministro, CamionesAsignados) {
		t.Error("La asignacion no es optima")
	}
}

