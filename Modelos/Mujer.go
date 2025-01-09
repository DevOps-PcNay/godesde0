package Modelos

type Mujer struct {
	Hombre // Herencia !!!! todas la propiedade del Hombre .
}

func (m *Mujer) Respirar()    { m.respirando = true }
func (m *Mujer) Comer()       { m.comiendo = true }
func (m *Mujer) Pensar()      { m.pensando = true }
func (m *Mujer) EstaVivo()    { m.vivo = true }
func (m *Mujer) Sexo() string { return "Mujer" }
