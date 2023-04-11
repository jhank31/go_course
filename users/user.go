package users

import (
	"fmt"
	"src/go_desde_0/modelos"
	"time"
)

func AltaUsuario() {
	user := new(modelos.User)
	user.AddUser(10, "jhancarlos", time.Now(), true)
	fmt.Println("Usuario creado ", user.Name)
}
