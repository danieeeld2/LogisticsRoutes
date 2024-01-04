package internal

var bd_prueba BD = BD{
	camiones: []CamionMatricula{
		{
			camion: NewCamion(TipoSuministro(NORMAL), 1000, 100),
			matricula: "1234ABC",
		},
		{
			camion: NewCamion(TipoSuministro(QUIMICO), 1000000, 100),
			matricula: "1234ABD",
		},
		{
			camion: NewCamion(TipoSuministro(NORMAL), 1000000, 100),
			matricula: "1234ABE",
		},
		{
			camion: NewCamion(TipoSuministro(NORMAL), 1000000, 1000),
			matricula: "1234ABF",
		},
	},
	suministros: []SuministroID{
		{
			suministro: NuevoSuministro("Calle Falsa 123", 10, 100, 1000, TipoSuministro(NORMAL)),
			id: "1",
		},
		{
			suministro: NuevoSuministro("Calle Falsa 123", 10, 10000, 1000000000, TipoSuministro(NORMAL)),
			id: "2",
		},
		{
			suministro: NuevoSuministro("Calle Falsa 123", 10, 10000, 1000000000, TipoSuministro(QUIMICO)),
			id: "3",
		},
	},
	asignaciones: []Asignacion{},
}

func getBDPrueba() BD{
	return bd_prueba
}