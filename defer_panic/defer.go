package defer_panic

import (
	"fmt"
)

func VerDefer() {
	fmt.Println("Este es un primer final ")
	defer fmt.Println("Este es el mensaje final ")

	fmt.Println("Este es el segundo mensaje")
}
