package internal

import "time"

type Suministro struct {
	direccion	string
	tiempo		time.Duration
	peso		float32
	dimensiones	[3]float32
	tipo		tipoCamion
}