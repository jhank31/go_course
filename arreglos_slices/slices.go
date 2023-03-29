package arreglos_slices

import "fmt"

var tabla3 []int = []int{2, 5, 4}

var arreglo [10]int = [10]int{3, 4, 5, 6, 1, 3, 2, 4}

func MostrarSlice() {
	fmt.Println(tabla3)

	porcion := arreglo[3:]   //slice creado desde un vector desde la posicion 3
	porcion2 := arreglo[:8]  //slice creado desde un vector desde la posicion 0 hasta la 8
	porcion3 := arreglo[3:9] //slice creado desde un vector desde la posicion 3 hasta la 9

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	//elementos := make([]int, 5, 20)
	//fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("Largo %d, Capacidad %d. ", len(nums), cap(nums))
}
