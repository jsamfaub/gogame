package main

import (
	"os"
	//"fmt"
	//"bytes"
	"image"
	_ "image/jpeg"
	_ "image/png"
	//"embed"
	"log"
	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 320
	screenHeight = 240
	frameOX     = 0
	frameOY     = 0
	frameWidth  = 32
	frameHeight = 32
)

var (
	playerImage *ebiten.Image
	backgroundImage *ebiten.Image
	posx int=0;
	posy int=0;
	frameNum int= 8
	leftSpeed int = 0
	rightSpeed int = 0
	upSpeed int = 0
	downSpeed int = 0
	flip int = 0
	stand bool = true
)

type Game struct {
	count int
	keys []ebiten.Key
}

func (g *Game) Update() error {
	g.keys=inpututil.AppendPressedKeys(g.keys[:0])
	g.count++
	return nil
}
func MoveRight() {stand=false;flip=0;rightSpeed=2}
func MoveLeft() {stand=false;flip=3*frameHeight;leftSpeed=2}
func MoveUp() {stand=false;upSpeed=2}
func MoveDown() {stand=false;downSpeed=2}

func (g *Game) Draw(screen *ebiten.Image) {
	stand=true
	leftSpeed, rightSpeed, upSpeed, downSpeed = 0, 0, 0, 0
	for _, p := range g.keys {
		switch(p){
			case ebiten.KeyD: MoveRight()
			case ebiten.KeyA: MoveLeft()
			case ebiten.KeyW: MoveUp()
			case ebiten.KeyS: MoveDown()
			case ebiten.KeyQ: os.Exit(0)
		}
	}
	posx+=rightSpeed; posx-=leftSpeed;
	posy+=downSpeed; posy-=upSpeed;

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	sx, sy := 0, 0
	screen.DrawImage(backgroundImage.SubImage(image.Rect(sx, sy, screenWidth, screenHeight)).(*ebiten.Image), op)

	op.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
	op.GeoM.Translate(screenWidth/2,screenHeight/2)
	op.GeoM.Translate(float64(posx),float64(posy))
	var standPos=0
	if(stand) {
		frameNum=5
	} else {
		standPos=32
		frameNum=8
	}
	i := (g.count/5) % frameNum
	sx, sy = frameOX+i*frameWidth, frameOY+flip+standPos
	screen.DrawImage(playerImage.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	background, _, err := ebitenutil.NewImageFromFile("red.jpg")
	player, _, err := ebitenutil.NewImageFromFile("fullrunner.png")
	if err != nil {
		log.Fatal(err)
	}
	playerImage = ebiten.NewImageFromImage(player)
	backgroundImage = ebiten.NewImageFromImage(background)
	
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Animation (Ebiten Demo)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
