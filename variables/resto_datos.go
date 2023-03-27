package variables

import (
	"fmt"
	"time"
	//"strconv"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func MostrarRestoDatos() float32 {
	Nombre := "Jhancarlos Lenis Arango"
	Estado := true
	Sueldo := 1566.44
	Fecha := time.Now()

	fmt.Println("nombre =", Nombre, "\nestado =", Estado, "\nSueldo =", Sueldo, "\nFecha =", Fecha)

	return float32(Sueldo)

}

func ConvertiraText(numero int) (bool, string) {
	//esta es una manera de convertir de int a string
	//texto := strconv.Itoa(numero)

	//TODO esta es otra manera de hacer lo mismo
	texto := fmt.Sprint(numero)
	return true, texto
}
