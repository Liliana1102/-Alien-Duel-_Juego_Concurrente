package scenes

import (
	 "fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"guerra/models"


	
)

type MainMenuScene struct {
	window fyne.Window
}

var status =true


var n *models.Jugador1


var d *models.Jugador2

// ball

var b *models.Ball

// ball2

var b2 *models.Ball2

func NewMainMenuScene(window fyne.Window) *MainMenuScene {
	return &MainMenuScene{window: window,}
}

func (s *MainMenuScene) Show() {
	// imagen j1
	j1 := canvas.NewImageFromURI(storage.NewFileURI("./assets/j1.png"))

	j1.Resize(fyne.NewSize(50,50))
	j1.Move(fyne.NewPos(700,500))

	n =models.NewJugador1(700,500, j1)

	// j2
	j2 := canvas.NewImageFromURI(storage.NewFileURI("./assets/j2.png"))

	j2.Resize(fyne.NewSize(50,50))
	j2.Move(fyne.NewPos(100,500))

	d =models.NewJugador2(100,100, j2)

	//ball
	ball := canvas.NewImageFromURI(storage.NewFileURI("./assets/ball.png"))
	bx,by := n.GetPosition()
	ball.Resize(fyne.NewSize(25,25))
	ball.Move(fyne.NewPos(bx,by))
	

	b = models.NewBall(bx,by,ball)

	//ball2
	ball2:= canvas.NewImageFromURI(storage.NewFileURI("./assets/ball 2.png"))
	b2x,b2y:=d.GetPosition2()
	ball2.Resize(fyne.NewSize(25,25))
	ball2.Move(fyne.NewPos(b2x,b2y))

	b2 = models.NewBall2(b2x,b2y,ball2)


	botonIniciar := widget.NewButton("Start Game", s.StartGame)
	botonIniciar.Resize(fyne.NewSize(150,30))
	botonIniciar.Move(fyne.NewPos(300,10))

	botonDetener := widget.NewButton("Stop Game", s.StopGame)
	botonDetener.Resize(fyne.NewSize(150,30))
	botonDetener.Move(fyne.NewPos(300,50))

	// nave se manda a llamar para que salga en la ventana
	s.window.SetContent(container.NewWithoutLayout(ball,ball2,j1,j2,botonIniciar, botonDetener))
	s.window.Canvas().SetOnTypedRune(s.OnTypedRune)
}

func (s *MainMenuScene) StartGame() {
	fmt.Println(":)")
	go s.Collicions()
	
}

func (s *MainMenuScene) OnTypedRune(r rune) {

	go n.Run(r)
	go d.Run(r)
	x,y := n.GetPosition()
	go b.Run(r,x,y)
	b.SetStatus(true)
	x2,y2 := d.GetPosition2()
	go b2.Run(r,x2,y2)
	b2.SetStatus(true)
	

}

func (s *MainMenuScene) Collicions() {
	
	for status{


		n1x, n1y := d.GetPosition2()
		n2x, n2y := n.GetPosition()
		b1x, b1y := b.Getballpos()
		b2x, b2y := b2.Getball2pos()

		margenDeError := float32(10) // Margen de error permitido
		const anchoDelEnemigo float32 = 50
		const altoDelEnemigo float32 = 50
		// Lógica para detectar colisiones con margen de error
		if b1x+margenDeError >= n1x && b1x-margenDeError <= n1x+anchoDelEnemigo &&
			b1y+margenDeError >= n1y && b1y-margenDeError <= n1y+altoDelEnemigo {
			
			fmt.Println("Colisión con bala 1")
			s.window.Close()
			

		} else if b2x+margenDeError >= n2x && b2x-margenDeError <= n2x+anchoDelEnemigo &&
			b2y+margenDeError >= n2y && b2y-margenDeError <= n2y+altoDelEnemigo {
			
			fmt.Println("Colisión con bala 2")
			s.window.Close()
			// Haz algo con la colisión de la bala 2, por ejemplo, detener el juego
		}
	}
	
}
func (s *MainMenuScene) StopGame() {
	// p.SetStatus(false)
	n.SetStatus(false)
	s.window.Close()
}