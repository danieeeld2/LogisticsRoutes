package internal

type TipoSuministro string

const (
	NORMAL       TipoSuministro = "normal"
	FRIGORIFICO  				= "frigorifico"
	QUIMICO 					= "quimico"
)

type Suministro struct {
	peso		float32
	dimensiones	[3]float32
	tipo		TipoSuministro
}