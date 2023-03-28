package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error
var texto string

func PedirNumero() string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingreasa un numero : ")
	for iteracion := 1; iteracion <= 10; {

		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Haz ingresado un tipo de dato erroneo. intentalo de nuevo")
				continue
			} else {
				break
			}

		}

	}
	for iteracion := 1; iteracion <= 10; iteracion++ {
		texto += fmt.Sprintf("%d x %d = %d \n", numero, iteracion, numero*iteracion)
	}
	fmt.Println(texto)
	return texto
}
