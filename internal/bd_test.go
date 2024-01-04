package internal

import (
	"testing"
)

func TestBD(t *testing.T) {
	resetearBD()
	bd := getBDPrueba()

	camion, _ := getCamionMatricula("AAAAAA", bd)
	if camion != (CamionMatricula{}) {
		t.Error("Camion encontrado y no existe")
	}
	camion, _ = getCamionMatricula("1234ABC", bd)
	if camion == (CamionMatricula{}) {
		t.Error("Camion no encontrado y existe")
	}
	camion, _ = deleteCamionMatricula("AAAAAA", bd)
	if camion != (CamionMatricula{}) {
		t.Error("Camion encontrado y no existe")
	}
	camion, _ = deleteCamionMatricula("1234ABC", bd)
	if camion == (CamionMatricula{}) {
		t.Error("Camion no encontrado y existe")
	}
	camion, _ = getCamionMatricula("1234ABC", bd)
	if camion != (CamionMatricula{}) {
		t.Error("Camion encontrado y fue eliminado antes")
	}
	camion, _ = putCamionMatricula("AAAAAA", Camion{}, bd)
	if camion != (CamionMatricula{}) {
		t.Error("Camion encontrado y no existe")
	}
	camion, _ = putCamionMatricula("1234ABD", Camion{}, bd)
	if camion != (CamionMatricula{Camion: Camion{}, Matricula: "1234ABD"}) {
		t.Error("Camion no actualizado")
	}
	suministro, _ := getSuministroID("AAAAAA", bd)
	if suministro != (SuministroID{}) {
		t.Error("Suministro encontrado y no existe")
	}
	suministro, _ = getSuministroID("1", bd)
	if suministro == (SuministroID{}) {
		t.Error("Suministro no encontrado y existe")
	}
	suministro, _ = deleteSuministroID("AAAAAA", bd)
	if suministro != (SuministroID{}) {
		t.Error("Suministro encontrado y no existe")
	}
	suministro, _ = deleteSuministroID("1", bd)
	if suministro == (SuministroID{}) {
		t.Error("Suministro no encontrado y existe")
	}
	suministro, _ = getSuministroID("1", bd)
	if suministro != (SuministroID{}) {
		t.Error("Suministro encontrado y fue eliminado antes")
	}
	suministro, _ = putSuministroID("AAAAAA", Suministro{}, bd)
	if suministro != (SuministroID{}) {
		t.Error("Suministro encontrado y no existe")
	}
	suministro, _ = putSuministroID("2", Suministro{}, bd)
	if suministro != (SuministroID{Suministro: Suministro{}, ID: "2"}) {
		t.Error("Suministro no actualizado")
	}
	asignacion, _ := postSuministroIDAsignacion("2", []string{"1234ABD"}, bd)
	if asignacion.IDSuministro != "2" && asignacion.MatriculasCamiones[0] != "1234ABD" {
		t.Error("Asignacion no actualizado")
	}
	_, status := getSuministroIDAsignacion("AAAAAA", bd)
	if status != false {
		t.Error("Asignacion encontrado y no existe")
	}
	asignacion, status = getSuministroIDAsignacion("3", bd)
	if status != true {
		t.Error("Asignacion no encontrado y existe")
	}
	if asignacion.IDSuministro != "3" && asignacion.MatriculasCamiones[0] != "1234ABF" {
		t.Error("Asignacion no encontrado correctamente")
	}
}