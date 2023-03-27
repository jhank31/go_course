package main

import "src/go_desde_0/teclado"

//"src/go_desde_0/ejercicios"
//"runtime"
//"src/go_desde_0/variables"

func main() {

	teclado.LeerDatos()

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
