package main

import (
	"fmt"
	"golang.org/x/image/colornames"
	_ "math"
	"time"
	_ "time"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func createWindow(title string, minX, minY, width, height float64, vsync, smooth bool) *pixelgl.Window {
	cfg := pixelgl.WindowConfig{
		Title: title,
		Bounds: pixel.R(minX,minY, width, height),
		VSync: vsync,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.SetSmooth(smooth)
	return win
}

func backgroundSprite(win *pixelgl.Window, image string, width, height float64) (*pixel.Sprite, pixel.Matrix) {
	background, err := graphic.LoadPicture("crashing-cars/images/" + image)
	if err != nil {
		panic(err)
	}

	bgSprite := pixel.NewSprite(background, background.Bounds())

	bgWidth, bgHeight := graphic.SpriteSize(bgSprite)
	bgMat := pixel.IM
	bgMat = bgMat.ScaledXY(pixel.ZV, pixel.V(width / bgWidth, height / bgHeight))
	bgMat = bgMat.Moved(win.Bounds().Center())
	return bgSprite, bgMat
}

func run() {
	win := createWindow("Car-Sim", 0,0, carmdl.WinWidth, carmdl.WinHeigth, false, true)

	// Get background image, simplifies code
	bgSprite, bgMat := backgroundSprite(win, "clean-intersection.png", carmdl.WinWidth, carmdl.WinHeigth )

	var (
		frames = 0
		second = time.Tick(time.Second)
	)
	last := time.Now()
	dt := time.Since(last).Seconds()

	testCar, testMat, testSprite := carmdl.MakeDefaultCar(0, carmdl.LeftIntention)

	for !win.Closed() {
		dt = time.Since(last).Seconds()
		last = time.Now()

		// Clean Screen
		win.Clear(colornames.White)

		// Test here
		testCar.Move()
		testSprite.Draw(win, testMat)

		// Draw here
		bgSprite.Draw(win, bgMat)
		win.Update()

		frames++
		select {
		case <-second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", "Car-Sim", frames))
			frames = 0
		default:
		}
	}
	fmt.Println(dt)
}

func main() {
	pixelgl.Run(run)
}