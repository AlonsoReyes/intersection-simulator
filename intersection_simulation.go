package main

import (
	"fmt"
	"github.com/AlonsoReyes/intersection-simulator/graphic"
	f "github.com/AlonsoReyes/intersection-simulator/intersection/fourway"
	v "github.com/AlonsoReyes/intersection-simulator/vehicle/car_generic"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	_ "image/png"
	"math"
	_ "math"
	"time"
	_ "time"
)

func createWindow(title string, minX, minY, width, height float64, vsync, smooth bool) *pixelgl.Window {
	cfg := pixelgl.WindowConfig{
		Title:  title,
		Bounds: pixel.R(minX, minY, width, height),
		VSync:  vsync,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.SetSmooth(smooth)
	return win
}

func backgroundSprite(win *pixelgl.Window, image string, width, height float64) (*pixel.Sprite, pixel.Matrix) {
	background, err := graphic.LoadPicture("src/github.com/AlonsoReyes/intersection-simulator/intersection/fourway/images/" + image)
	if err != nil {
		panic(err)
	}

	bgSprite := pixel.NewSprite(background, background.Bounds())

	bgWidth, bgHeight := graphic.SpriteSize(bgSprite)
	bgMat := pixel.IM
	bgMat = bgMat.ScaledXY(pixel.ZV, pixel.V(width/bgWidth, height/bgHeight))
	bgMat = bgMat.Moved(win.Bounds().Center())
	return bgSprite, bgMat
}

func run() {
	win := createWindow("Car-Sim", 0, 0, f.PictureLength, f.PictureLength, false, true)

	// Get background image, simplifies code
	bgSprite, bgMat := backgroundSprite(win, "inter.png", f.PictureLength, f.PictureLength)

	var (
		frames = 0
		second = time.Tick(time.Second)
	)

	last := time.Now()
	dt := time.Since(last).Seconds()

	carPic, err := graphic.LoadPicture("src/github.com/AlonsoReyes/intersection-simulator/vehicle/car_generic/images/redcar.png")
	if err != nil {
		panic(err)
	}

	carSprite := pixel.NewSprite(carPic, carPic.Bounds())

	lane := 3
	intention := v.RightIntention
	coopZoneLength := f.PictureLength
	dangerZoneLength := f.IntersectionLength

	testCar := v.CreateCar(lane, intention, coopZoneLength, dangerZoneLength)

	for !win.Closed() {
		dt = time.Since(last).Seconds()
		last = time.Now()

		// Clean Screen
		win.Clear(colornames.White)

		// Test here
		//fmt.Println(testCar.Speed)
		//		fmt.Println(testCar.Direction)
		testCar.Run(dt)

		mat := pixel.IM
		mat = mat.ScaledXY(pixel.ZV, pixel.V(0.1, 0.1))
		mat = mat.Rotated(pixel.ZV, testCar.Direction*math.Pi/180)
		mat = mat.Moved(pixel.V(testCar.Position.X, testCar.Position.Y))

		// Draw here
		bgSprite.Draw(win, bgMat)
		carSprite.Draw(win, mat)
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
