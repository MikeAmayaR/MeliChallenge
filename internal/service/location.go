package service

import "errors"

var satellites = map[string][2]float32{
	"Kenobi":    {-500, -200},
	"Skywalker": {100, -100},
	"Sato":      {500, 100},
}

func GetLocation(distances ...float32) (x, y float32, err error) {
	if len(distances) != 3 {
		return 0, 0, errors.New("se requieren exactamente tres distancias")
	}

	// Coordenadas de los satélites
	kenobi := satellites["Kenobi"]
	skywalker := satellites["Skywalker"]
	sato := satellites["Sato"]

	d1, d2, d3 := distances[0], distances[1], distances[2]

	A := 2*skywalker[0] - 2*kenobi[0]
	B := 2*skywalker[1] - 2*kenobi[1]
	C := d1*d1 - d2*d2 + kenobi[0]*kenobi[0] - skywalker[0]*skywalker[0] + kenobi[1]*kenobi[1] - skywalker[1]*skywalker[1]
	D := 2*sato[0] - 2*skywalker[0]
	E := 2*sato[1] - 2*skywalker[1]
	F := d2*d2 - d3*d3 + skywalker[0]*skywalker[0] - sato[0]*sato[0] + skywalker[1]*skywalker[1] - sato[1]*sato[1]

	denominator := A*E - B*D
	if denominator == 0 {
		return 0, 0, errors.New("error matemático: divisor cero")
	}

	x = (C*E - F*B) / denominator
	y = (C*D - A*F) / (B*D - A*E)

	return x, y, nil
}
