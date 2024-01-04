package internal

type CamionMatricula struct {
	camion Camion
	matricula string
}

type SuministroID struct {
	suministro Suministro
	id string
}

type Asignacion struct {
	matriculasCamiones []string
	idsuministro string
}

type BD struct {
	camiones []CamionMatricula
	suministros []SuministroID
	asignaciones []Asignacion
}

func getCamionMatricula(matricula string, bd BD) (CamionMatricula, bool) {
	for _, CamionMatricula := range bd.camiones {
		if CamionMatricula.matricula == matricula {
			return CamionMatricula, true
		}
	}
	return CamionMatricula{}, false
}

func getSuministroID(id string, bd BD) (SuministroID, bool) {
	for _, SuministroID := range bd.suministros {
		if SuministroID.id == id {
			return SuministroID, true
		}
	}
	return SuministroID{}, false
}

func deleteCamionMatricula(matricula string, bd BD) (CamionMatricula, bool) {
	for i, CamionMatricula := range bd.camiones {
		if CamionMatricula.matricula == matricula {
			bd.camiones = append(bd.camiones[:i], bd.camiones[i+1:]...)
			return CamionMatricula, true
		}
	}
	return CamionMatricula{}, false
}

func deleteSuministroID(id string, bd BD) (SuministroID, bool) {
	for i, SuministroID := range bd.suministros {
		if SuministroID.id == id {
			bd.suministros = append(bd.suministros[:i], bd.suministros[i+1:]...)
			return SuministroID, true
		}
	}
	return SuministroID{}, false
}

func putCamionMatricula(matricula string, NuevoCamion Camion, bd BD) (CamionMatricula, bool) {
	for i, CamionMatricula := range bd.camiones {
		if CamionMatricula.matricula == matricula {
			bd.camiones[i].camion = NuevoCamion
			return bd.camiones[i], true
		}
	}
	return CamionMatricula{}, false
}

func putSuministroID(id string, NuevoSuministro Suministro, bd BD) (SuministroID, bool) {
	for i, SuministroID := range bd.suministros {
		if SuministroID.id == id {
			bd.suministros[i].suministro = NuevoSuministro
			return bd.suministros[i], true
		}
	}
	return SuministroID{}, false
}

func getSuministroIDAsignacion(id string, bd BD) (Asignacion, bool) {
	for _, Asignacion := range bd.asignaciones {
		if Asignacion.idsuministro == id {
			return Asignacion, true
		}
	}
	return Asignacion{}, false
}

func postSuministroIDAsignacion(id string, matriculasCamiones []string, bd BD) (Asignacion, bool) {
	for _, Asignacion := range bd.asignaciones {
		if Asignacion.idsuministro == id {
			return Asignacion, false
		}
	}

	bd.asignaciones = append(bd.asignaciones, Asignacion{matriculasCamiones, id})
	return Asignacion{matriculasCamiones, id}, true
}





