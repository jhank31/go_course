package main

import (
	"fmt"
	"src/go_desde_0/variables"
)

func main() {

	// variables.MuestroEnteros()
	// variables.MostrarRestoDatos()
	// fmt.Print(variables.MostrarRestoDatos())

	estado, texto := variables.ConvertiraText(1234)
	fmt.Println(estado, texto)
}
