package modelos

type Mujer struct {
	Hombre
	EsMadre bool
}

func (this *Mujer) Respirar() {
	this.respirando = true
}
func (this *Mujer) Comer() {
	this.comiendo = true
}
func (this *Mujer) Pensar() {
	this.pensando = true
}
func (this *Mujer) Sexo() string {
	return "Mujer"
}