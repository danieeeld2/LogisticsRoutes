package internal

type CamionMatricula struct {
	Camion    Camion `json:"camion"`
	Matricula string `json:"matricula"`
}

type SuministroID struct {
	Suministro Suministro `json:"suministro"`
	ID         string     `json:"id"`
}

type Asignacion struct {
	MatriculasCamiones []string `json:"matriculasCamiones"`
	IDSuministro       string   `json:"idSuministro"`
}

type BD struct {
	Camiones     []CamionMatricula `json:"camiones"`
	Suministros  []SuministroID    `json:"suministros"`
	Asignaciones []Asignacion      `json:"asignaciones"`
}

func getCamionMatricula(matricula string, bd BD) (CamionMatricula, bool) {
	for _, CamionMatricula := range bd.Camiones {
		if CamionMatricula.Matricula == matricula {
			return CamionMatricula, true
		}
	}
	return CamionMatricula{}, false
}

func getSuministroID(id string, bd BD) (SuministroID, bool) {
	for _, SuministroID := range bd.Suministros {
		if SuministroID.ID == id {
			return SuministroID, true
		}
	}
	return SuministroID{}, false
}

func deleteCamionMatricula(matricula string, bd BD) (CamionMatricula, bool) {
	for i, CamionMatricula := range bd.Camiones {
		if CamionMatricula.Matricula == matricula {
			bd.Camiones = append(bd.Camiones[:i], bd.Camiones[i+1:]...)
			return CamionMatricula, true
		}
	}
	return CamionMatricula{}, false
}

func deleteSuministroID(id string, bd BD) (SuministroID, bool) {
	for i, SuministroID := range bd.Suministros {
		if SuministroID.ID == id {
			bd.Suministros = append(bd.Suministros[:i], bd.Suministros[i+1:]...)
			return SuministroID, true
		}
	}
	return SuministroID{}, false
}

func putCamionMatricula(matricula string, NuevoCamion Camion, bd BD) (CamionMatricula, bool) {
	for i, CamionMatricula := range bd.Camiones {
		if CamionMatricula.Matricula == matricula {
			bd.Camiones[i].Camion = NuevoCamion
			return bd.Camiones[i], true
		}
	}
	return CamionMatricula{}, false
}

func putSuministroID(id string, NuevoSuministro Suministro, bd BD) (SuministroID, bool) {
	for i, SuministroID := range bd.Suministros {
		if SuministroID.ID == id {
			bd.Suministros[i].Suministro = NuevoSuministro
			return bd.Suministros[i], true
		}
	}
	return SuministroID{}, false
}

func getSuministroIDAsignacion(id string, bd BD) (Asignacion, bool) {
	for _, Asignacion := range bd.Asignaciones {
		if Asignacion.IDSuministro == id {
			return Asignacion, true
		}
	}
	return Asignacion{}, false
}

func postSuministroIDAsignacion(id string, matriculasCamiones []string, bd BD) (Asignacion, bool) {
	for _, Asignacion := range bd.Asignaciones {
		if Asignacion.IDSuministro == id {
			return Asignacion, false
		}
	}

	bd.Asignaciones = append(bd.Asignaciones, Asignacion{matriculasCamiones, id})
	return Asignacion{matriculasCamiones, id}, true
}





