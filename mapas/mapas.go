package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)

	paises["mexico"] = "D.F"
	paises["argentina"] = "Buenos aires"
	paises["colombia"] = "Bogota"

	fmt.Println(paises)

	campeonato := map[string]int{
		"Barcelona":   39,
		"Real Madrid": 100,
		"America":     40,
	}
	//fmt.Println(campeonato)

	// for equipo, puntaje := range campeonato {
	// 	fmt.Printf("Equipo %s, tiene un puntaje de %d \n", equipo, puntaje)
	// }

	delete(campeonato, "Real Madrid")

	for equipo, puntaje := range campeonato {
		fmt.Printf("Equipo %s, tiene un puntaje de %d \n", equipo, puntaje)
	}
}
