package main

import (
	"fmt"
	"src/go_desde_0/goroutines"
)

//"src/go_desde_0/ejercicios"
//"runtime"
//"src/go_desde_0/variables"

func main() {

	//arreglos_slices.MostrarSlice()
	//arreglos_slices.Capacidad()
	//users.AltaUsuario()
	// Valentina := new(modelos.Mujer)
	// ejer_interfaces.HumanosRespirando(Valentina)
	//defer_panic.EjemploPanic()
	go goroutines.MiNombreLento("Jhancarlos")

	fmt.Println("estoy aqui")

	var x string

	fmt.Scanln(&x)
	//teclado.LeerDatos()
	//iteraciones.Iterar()
	//files.LeoArchivo()
	//funciones.Calculos()
	//funciones.LlamarClosure()
	//funciones.Exponencia(4)

	// _, texto := ejercicios.ConvertirAEntero("20")

	// println(texto)

	// variables.MuestroEnteros()
	// variables.MostrarRestoDatos()
	// fmt.Print(variables.MostrarRestoDatos())

	//estado, texto := variables.ConvertiraText(1234)
	//fmt.Println(estado, texto)

	// ----------------------------------------------------------------------------------

	//se puede hacer la asignacion de una variable directamente en un condicional
	// if os := runtime.GOOS; os == "Linux." || os == "OS X." {
	// 	fmt.Println("Es Linux")
	// } else {
	// 	fmt.Println("Es Windows")
	// }

	//-------------------------------------------------------------------------------------
	// switch funciona como en los otros lenguajes practicamente

	// switch os := runtime.GOOS; os {
	// case "linux":
	// 	fmt.Println("Es linux")
	// 	break
	// case "darwin":
	// 	fmt.Println("esto es darwin")
	// 	break
	// default:
	// 	fmt.Printf("%s \n", os)
	// }

}
