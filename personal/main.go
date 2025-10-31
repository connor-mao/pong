package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
	"math/rand"
	"time"
)

const (
	screenWidth = 640
	screenHeight = 480
	ballWidth = 20
	ballHeight = 20

	paddleWidth = 10
	paddleHeight = 80
	paddleSpeed = 4
	paddleOffset = 30
)

type Game struct{
	ballX, ballY float64
	ballVX, ballVY float64
	ballImage *ebiten.Image
	ballColor color.Color

	leftPaddleY, rightPaddleY float64
	leftPaddle, rightPaddle *ebiten.Image
}

func newGame() *Game {
	ball := ebiten.NewImage(ballWidth, ballHeight)
	ball.Fill(color.White)

	left := ebiten.NewImage(paddleWidth, paddleHeight)
	left.Fill(color.White)

	right := ebiten.NewImage(paddleWidth, paddleHeight)
	right.Fill(color.White)

	return &Game{
		ballX: screenWidth/2 - ballWidth/2,
		ballY: screenHeight/2 - ballHeight/2,
		ballVX: 2,
		ballVY: 2,
		ballImage: ball,
		ballColor: color.White,

		leftPaddle: left,
		rightPaddle: right,
		leftPaddleY: screenHeight/2 - paddleHeight/2,
		rightPaddleY: screenHeight/2 - paddleHeight/2,

	}
}

func (g *Game) Update() error {
	//players
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		if(g.leftPaddleY > 0) {
			g.leftPaddleY -= paddleSpeed
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		if(g.leftPaddleY < screenHeight - paddleHeight) {
			g.leftPaddleY += paddleSpeed
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		if(g.rightPaddleY > 0){
			g.rightPaddleY -= paddleSpeed
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		if(g.rightPaddleY < screenHeight - paddleHeight) {
			g.rightPaddleY += paddleSpeed
		}
	}

	g.ballX += g.ballVX
	g.ballY += g.ballVY

	//bounce off walls
	//coordinates start at top left corner
	if g.ballX < 0 || g.ballX + ballWidth > screenWidth {
		g.ballVX = -g.ballVX
	}

	if g.ballY < 0 || g.ballY > screenHeight - ballHeight {
		g.ballVY = -g.ballVY
	}

	//bounce off paddles
	//left paddle
	if g.ballX <= paddleOffset + paddleWidth {
		if g.ballY + ballHeight >= g.leftPaddleY && g.ballY <= g.leftPaddleY + paddleHeight {
			g.ballVX = -g.ballVX
			g.ballColor = randomColor()
		}
	}

	if g.ballX + ballWidth >= screenWidth - paddleWidth - paddleOffset {
		if g.ballY + ballHeight >= g.rightPaddleY && g.ballY <= g.rightPaddleY + paddleHeight {
			g.ballVX = -g.ballVX
			g.ballColor = randomColor()
		}
	}
	return nil
}

//draw is called every frame to render the screen
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	g.ballImage.Fill(g.ballColor)

	ballOpts := &ebiten.DrawImageOptions{}
	ballOpts.GeoM.Translate(g.ballX, g.ballY)
	screen.DrawImage(g.ballImage, ballOpts)

	leftOpts := &ebiten.DrawImageOptions{}
	leftOpts.GeoM.Translate(paddleOffset, g.leftPaddleY)
	screen.DrawImage(g.leftPaddle, leftOpts)

	rightOpts := &ebiten.DrawImageOptions{}
	rightOpts.GeoM.Translate(screenWidth - paddleOffset - paddleWidth, g.rightPaddleY)
	screen.DrawImage(g.rightPaddle, rightOpts)
}

//layout sets screen size, returns width and height
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func randomColor() color.Color {
	return color.RGBA{
		R: uint8(rand.Intn(256)),
		G: uint8(rand.Intn(256)),
		B: uint8(rand.Intn(256)),
		A: 255,
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Pong")
	if err := ebiten.RunGame(newGame()); err != nil {
		log.Fatal(err)
	}
}
