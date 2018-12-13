package carmdl

import (
	"github.com/AlonsoReyes/intersection-simulator/graphic"
	"github.com/faiface/pixel"
	_ "math"
)


type Vehicle interface {
	Move() // Advances, accelerates and turns
	Advance() // Moves to next coordinates
	Accelerate()
	Turn()
	GetPosition()
}

func GetInitialPosition(lane int) (float64, float64) {
	return Lanes[lane].X, Lanes[lane].Y
}

func MakeCar(speed float64, lane, intention int) (*Car, pixel.Matrix, *pixel.Sprite)  {
	x, y := GetInitialPosition(lane)
	car := CreateCar(x, y, speed, intention, lane)
	car.Direction = InitialDirections[lane]
	mat, sprite := CarGraphicFun(car)
	return car, mat, sprite
}

func MakeDefaultCar(lane, intention int) (*Car, pixel.Matrix, *pixel.Sprite) {
	x, y := GetInitialPosition(lane)
	car := CreateCar(x, y, DefaultInitialSpeed, intention, lane)
	car.Direction = InitialDirections[lane]
	mat, sprite := CarGraphicFun(car)
	return car, mat, sprite
}

func CarGraphicFun(car *Car) (pixel.Matrix, *pixel.Sprite) {
	carPic, err := graphic.LoadPicture(DefaultCarImage)
	if err != nil {
		panic(err)
	}
	carSprite := pixel.NewSprite(carPic, carPic.Bounds())
	carSpriteW, carSpriteH := graphic.SpriteSize(carSprite)
	newW := CarWidth / carSpriteW
	newH := CarHeight / carSpriteH
	carMat := pixel.IM
	carMat = carMat.ScaledXY(pixel.ZV, pixel.V(newW, newH))
	carMat = carMat.Rotated(pixel.ZV,car.Direction)
	carMat = carMat.Moved(pixel.V(car.Pos.X, car.Pos.Y))

	return carMat, carSprite
}