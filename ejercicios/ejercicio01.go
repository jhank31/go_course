package ejercicios

import (
	"strconv"
)

func ConvertirAEntero(numero string) (int, string) {
	numeroReturn, error := strconv.Atoi(numero)
	if error != nil {
		return 0, "hubo un error" + error.Error()
	} else if numeroReturn >= 100 {
		return numeroReturn, "El numero es mayor a 100"
	} else {
		return numeroReturn, "El numero es menor a 100"
	}
}
