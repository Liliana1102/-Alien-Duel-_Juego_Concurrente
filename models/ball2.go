package models

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Ball2 struct {
	posX, posY float32
	status bool	
	ball2 *canvas.Image
}

func NewBall2(posx float32, posy float32, img *canvas.Image) *Ball2 {
	return &Ball2{
		posX: posx,
		posY: posy,
		status: true,
		ball2: img,
	}
}

func (q *Ball2) Run(r rune, x float32, y float32) {
    var ban int 
	ban=0
    if r == 'j' {
        ban = 1 
    }

    // Si el estado está activo, la bola se mueve
    if ban == 1 {
        for q.status {
            q.posX += 10
            
            if q.posX >= 775 {
                q.posX = x
                q.posY = y
                q.ball2.Move(fyne.NewPos(q.posX, q.posY))
                q.status = false
            }
            q.ball2.Move(fyne.NewPos(q.posX, q.posY))
            fmt.Println("Coordenadas de la bola", q.posX, "  ", q.posY)
            time.Sleep(100 * time.Millisecond)
        }
    } else {
        

        q.posY = y
        q.ball2.Move(fyne.NewPos(q.posX, q.posY))
    }
}
func (p *Ball2) Getball2pos( )(float32,float32) {
	return p.posX,p.posY
	
}
func (p *Ball2) SetStatus(status bool) {
	p.status = status
}