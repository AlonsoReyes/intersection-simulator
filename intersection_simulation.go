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
	"github.com/AlonsoReyes/intersection-simulator/graphic"
	f "github.com/AlonsoReyes/intersection-simulator/intersection/fourway"
	"github.com/AlonsoReyes/intersection-simulator/vehicle/car_generic"
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
	background, err := graphic.LoadPicture("intersection-simulator/intersection/fourway/images/" + image)
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
	win := createWindow("Car-Sim", 0,0, f.PictureLength, f.PictureLength, false, true)

	// Get background image, simplifies code
	bgSprite, bgMat := backgroundSprite(win, "clean-intersection.png", f.PictureLength, f.PictureLength )

	var (
		frames = 0
		second = time.Tick(time.Second)
	)

	last := time.Now()
	dt := time.Since(last).Seconds()


	carPic, err := graphic.LoadPicture("intersection-simulator/vehicle/car_generic/images/redcar.png")
	if err != nil {
		panic(err)
	}

	carSprite := pixel.NewSprite(carPic, carPic.Bounds())

	lane := 0
	intention := 0
	coopZoneLength := f.PictureLength
	dangerZoneLength := f.IntersectionLength
	testCar := car_generic.CreateCar(lane, intention, coopZoneLength, dangerZoneLength)


	for !win.Closed() {
		dt = time.Since(last).Seconds()
		last = time.Now()

		// Clean Screen
		win.Clear(colornames.White)

		// Test here
		testCar.Run(dt)

		mat := pixel.IM
		mat = mat.ScaledXY(pixel.ZV, pixel.V(0.1, 0.1))
		mat = mat.Rotated(pixel.ZV, testCar.Direction)
		mat = mat.Moved(pixel.V(testCar.Position.X, testCar.Position.Y))
		carSprite.Draw(win, mat)

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