package arreglos_slices

import "fmt"

// asi se crea un array

var tabla [10]int = [10]int{10, 43}

// asi se crea una matriz

var matriz [10][12]int

func MuestrArreglos() {
	tabla[7] = 33
	tabla[2] = 4

	tabla2 := [10]string{"hola", "carlitos", "andy"}

	for ite := 0; ite < len(tabla2); ite++ {
		fmt.Println(tabla2[ite])
	}

	matriz[7][6] = 15

	fmt.Println(matriz)
}
