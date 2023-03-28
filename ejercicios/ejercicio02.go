package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error

func PedirNumero() {
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
		fmt.Printf("%d x %d = %d \n", numero, iteracion, numero*iteracion)
	}

}
