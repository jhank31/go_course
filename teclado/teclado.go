package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func LeerDatos() {
	fmt.Println("Ingrese numero 1 : ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		numero1, _ = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingreado es incorrecto " + err.Error())
		}
	}
}
