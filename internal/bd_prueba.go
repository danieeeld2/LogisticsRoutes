package internal

var bd_prueba BD = BD{}

func getBDPrueba() BD{
	return bd_prueba
}

func resetearBD() {
	bd_prueba = BD{
		Camiones: []CamionMatricula{
			{
				Camion: NewCamion(TipoSuministro(NORMAL), 1000, 100),
				Matricula: "1234ABC",
			},
			{
				Camion: NewCamion(TipoSuministro(QUIMICO), 1000000, 100),
				Matricula: "1234ABD",
			},
			{
				Camion: NewCamion(TipoSuministro(NORMAL), 1000000, 100),
				Matricula: "1234ABE",
			},
			{
				Camion: NewCamion(TipoSuministro(NORMAL), 1000000, 1000),
				Matricula: "1234ABF",
			},
		},
		Suministros: []SuministroID{
			{
				Suministro: NuevoSuministro("Calle Falsa 123", 10, 100, 1000, TipoSuministro(NORMAL)),
				ID: "1",
			},
			{
				Suministro: NuevoSuministro("Calle Falsa 123", 10, 10, 10, TipoSuministro(NORMAL)),
				ID: "2",
			},
			{
				Suministro: NuevoSuministro("Calle Falsa 123", 10, 10000, 1000000000, TipoSuministro(QUIMICO)),
				ID: "3",
			},
		},
		Asignaciones: []Asignacion{
			{
				MatriculasCamiones: []string{"1234ABF"},
				IDSuministro: "3",
			},
		},			
	}
}