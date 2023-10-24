package models


import (
	// "fmt"


	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"

)

type Jugador2 struct {
	posX, posY float32
	status bool	
	j2 *canvas.Image
}

func NewJugador2(posx float32, posy float32, img *canvas.Image) *Jugador2 {
	return &Jugador2{
		posX: posx,
		posY: posy,
		status: true,
		j2: img,
	}
}

func (q *Jugador2) Run(r rune) {

	if r == 'i' {

		// Mover hacia arriba
		q.posY -= 10 // Puedes ajustar el valor según tus necesidades
	}
	if r == 'k'{

		// Mover hacia abajo
		q.posY += 10 // Puedes ajustar el valor según tus necesidades
	}


	// Limitar la posición Y dentro de los límites de la ventana
	if q.posY <= 0 {
		q.posY = 0
	} else if q.posY >= 550 {
		q.posY = 550
	}

	// Actualizar la posición de la nave
	q.j2.Move(fyne.NewPos(q.posX, q.posY))
}
func (n *Jugador2) GetPosition2() (float32,float32){
	return n.posX,n.posY
}

func (n *Jugador2) SetStatus(status bool) {
	n.status = status
}
