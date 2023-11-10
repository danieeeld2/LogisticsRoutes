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
	
	var suma_mma uint
	var volumen_total, volumen_suministro float32
	
	suma_mma = 0 
	volumen_total = 0 
	volumen_suministro = suministro.dimensiones_whd_cm[0]*suministro.dimensiones_whd_cm[1]*suministro.dimensiones_whd_cm[2] 
	for _, camion := range camiones {
		suma_mma += camion.mma
		volumen_total += camion.dimensiones_whd_cm[0]*camion.dimensiones_whd_cm[1]*camion.dimensiones_whd_cm[2]
	}
	
	if suma_mma <= suministro.mma || volumen_total <= volumen_suministro {
		return false
	}
	
	return true
}

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

func CalcularConjuntosDeCamiones(camiones []Camion, t uint) [][]Camion{
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

func FiltrarPorTipo(subconjuntos [][]Camion) [][]Camion {
	var subconjuntosFiltrados [][]Camion

	for _, conjunto := range subconjuntos {
		tipoAnterior := TipoSuministro("") 

		todosDelMismoTipo := true
		for _, camion := range conjunto {
			if tipoAnterior == "" {
				tipoAnterior = camion.tipo
			} else if camion.tipo != tipoAnterior {
				todosDelMismoTipo = false
				break
			}
			tipoAnterior = camion.tipo
		}

		if todosDelMismoTipo {
			subconjuntosFiltrados = append(subconjuntosFiltrados, conjunto)
		}
	}

	return subconjuntosFiltrados
}
func ComprobarAsignacionOptima(CamionesDisponibles []Camion, suministro Suministros, CamionesAsignados []Camion) bool {
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
			if PuedeTransporarSuministro(conjunto, suministro) && EsMasChico(conjunto, CamionesAsignados) {
				return false
			}
		}
	}
	
	
	return true
} 

func TestPlanificacion(t *testing.T) {
	t.log("Comenzando test de planificacion")

	camionesDisponibles := []Camion{}
	CamionesAsignados := []Camion{}
	suministro := NuevoSuministro("Calle Falsa 123", 10, 100, [3]float32{10, 10, 10}, TipoSuministro.NORMAL)

	t.log("No hay camiones disponibles ni asignados")
	AsigarCamiones(&camionesDisponibles, suministro, &CamionesAsignados)
	if !ComprobarAsignacionOptima(camionesDisponibles, suministro, CamionesAsignados) {
		t.Error("La asignacion no es optima")
	}

	camionesDisponibles.append(NewCamion(TipoSuministro.NORMAL, [3]float32{10, 10, 10}, 100))
	camionesDisponibles.append(NewCamion(TipoSuministro.QUIMICO, [3]float32{100, 100, 100}, 100))
	camionesDisponibles.append(NewCamion(TipoSuministro.NORMAL, [3]float32{1000, 1000, 1000}, 100))
	camionesDisponibles.append(NewCamion(TipoSuministro.NORMAL, [3]float32{1000, 1000, 1000}, 1000))
	CamionesAsignados = []Camion{}
	t.log("Hay varios camiones disponibles que pueden transportar el vehículo")
	AsigarCamiones(&camionesDisponibles, suministro, &CamionesAsignados)
	if !ComprobarAsignacionOptima(camionesDisponibles, suministro, CamionesAsignados) {
		t.Error("La asignacion no es optima")
	}

	suministro = NuevoSuministro("Calle Falsa 123", 10, 10000, [3]float32{1000, 1000, 1000}, TipoSuministro.NORMAL)
	CamionesAsignados = []Camion{}
	t.log("Hay varios camiones disponibles pero ninguno puede transportar el suministro")
	AsigarCamiones(&camionesDisponibles, suministro, &CamionesAsignados)
	if !ComprobarAsignacionOptima(camionesDisponibles, suministro, CamionesAsignados) {
		t.Error("La asignacion es optima")
	}

	suministro = NuevoSuministro("Calle Falsa 123", 10, 100, [3]float32{10, 10, 10}, TipoSuministro.QUIMICO)
	CamionesAsignados = []Camion{}
	t.log("Comprobar que le asigna el de tipo correcto")
	AsigarCamiones(&camionesDisponibles, suministro, &CamionesAsignados)
	if !ComprobarAsignacionOptima(camionesDisponibles, suministro, CamionesAsignados) {
		t.Error("La asignacion no es optima")
	}

	suministro = NuevoSuministro("Calle Falsa 123", 10, 1050, [3]float32{10, 10, 10}, TipoSuministro.NORMAL)
	CamionesAsignados = []Camion{}
	t.log("Necesita mas de un camión para ser transportado")
	AsigarCamiones(&camionesDisponibles, suministro, &CamionesAsignados)
	if !ComprobarAsignacionOptima(camionesDisponibles, suministro, CamionesAsignados) {
		t.Error("La asignacion no es optima")
	}
}

