package files

import (
	"bufio"
	"fmt"
	"os"
	"src/go_desde_0/ejercicios"
)

var filename string = "./files/txt/tabla.txt"

func Grabartabla() {
	var texto string = ejercicios.PedirNumero()
	archivo, err := os.Create(filename)
	if err != nil {
		fmt.Println("error al crear el archivo " + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.PedirNumero()
	if !Append(filename, texto) {
		fmt.Println("error al concatenar contenido")
	}
}

func Append(file string, texto string) bool {
	arch, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el append " + err.Error())
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el WriteString " + err.Error())
		return false
	}

	arch.Close()

	return true
}

func LeoArchivo() {
	archivo, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error al leer el archivo " + err.Error())
		return
	}
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> ", registro)
	}
	archivo.Close()
}
