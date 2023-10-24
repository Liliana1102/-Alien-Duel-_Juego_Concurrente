package models

import (
	// "fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Jugador1 struct {
	posX, posY float32
	status     bool
	j1         *canvas.Image
}

func NewJugador1(posx float32, posy float32, img *canvas.Image) *Jugador1 {
	return &Jugador1{
		posX:   posx,
		posY:   posy,
		status: true,
		j1:     img,
	}
}

func (n *Jugador1) Run(r rune) {


	if r == 'w' {
		// fmt.Println("j2 up")
		// Mover hacia arriba
		n.posY -= 10 // Puedes ajustar el valor según tus necesidades
	}
	if r == 's'{
		// fmt.Println("j2 dn")
		// Mover hacia abajo
		n.posY += 10 // Puedes ajustar el valor según tus necesidades
	}


	// Limitar la posición Y dentro de los límites de la ventana
	if n.posY <= 0 {
		n.posY = 0
	} else if n.posY >= 550 {
		n.posY = 550
	}

	// Actualizar la posición de la nave
	n.j1.Move(fyne.NewPos(n.posX, n.posY))


}
func (n *Jugador1) GetPosition() (float32,float32){
	return n.posX,n.posY
}
func (n *Jugador1) SetStatus(status bool) {
	n.status = status
}
