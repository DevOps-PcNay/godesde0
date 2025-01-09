package Modelos

type Hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	vivo       bool
}

// Solo se si se ejecuta una linea.
// (h *Hombre) = Para que la funcion pertenezca a una estructura.
// Respirar() se definio en la Interfaz "Humano", y se debe colocar para hacer referencia.

func (h *Hombre) Respirar() {
	h.respirando = true
}

// Puede ser en una sola linea, ya que es solo una instruccion.
func (h *Hombre) Comer()         { h.comiendo = true }
func (h *Hombre) Pensar()        { h.pensando = true }
func (h *Hombre) EstaVivo() bool { h.vivo = true }
func (h *Hombre) Sexo() string   { return "Hombre" }
