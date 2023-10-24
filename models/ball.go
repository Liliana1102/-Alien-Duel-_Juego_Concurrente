package models

import (
	// "fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Ball struct {
	posX, posY float32
	status bool	
	ball *canvas.Image
}

func NewBall(posx float32, posy float32, img *canvas.Image) *Ball {
	return &Ball{
		posX: posx,
		posY: posy,
		status: true,
		ball: img,
	}
}

func (q *Ball) Run(r rune, x float32, y float32) {
    var ban int // Utilizamos un entero en lugar de float32 para representar un estado
	ban=0
    if r == ' ' {
        ban = 1 // Si se presionó espacio, activamos el estado
    }

    // Si el estado está activo, la bola se mueve
    if ban == 1 {
        for q.status {
            q.posX -= 10
            // Limitar la posición X dentro de los límites de la ventana
            if q.posX <= 0 {
                q.posX = x
                q.posY = y
                q.ball.Move(fyne.NewPos(q.posX, q.posY))
                q.status = false
            }
            q.ball.Move(fyne.NewPos(q.posX, q.posY))
            // fmt.Println("Coordenadas de la bola", q.posX, "  ", q.posY)
            time.Sleep(100 * time.Millisecond)
        }
    } else {
        // Si el estado no está activo (no se presionó espacio), la bola se pega a (x, y)

        q.posY = y
        q.ball.Move(fyne.NewPos(q.posX, q.posY))
    }
}
func (p *Ball) Getballpos( )(float32,float32) {
	return p.posX,p.posY
	
}

func (p *Ball) SetStatus(status bool) {
	p.status = status
}