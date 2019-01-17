package car_generic

import (
	"github.com/faiface/pixel"
	"github.com/niclabs/intersection-simulator/graphic"
)

func GetCarImage(image string) (pixel.Picture, error) {
	carPic, err := graphic.LoadPicture("src/github.com/niclabs/intersection-simulator/vehicle/car_generic/images/" + image)
	if err != nil {
		panic(err)
	}

	return carPic, nil
}

func GetCarSprite(carPicture pixel.Picture) *pixel.Sprite {
	carSprite := pixel.NewSprite(carPicture, carPicture.Bounds())
	return carSprite
}

func GetCarMatrix(car *Car) pixel.Matrix {
	mat := pixel.IM
	mat = mat.ScaledXY(pixel.ZV, pixel.V(Scaling, Scaling))
	mat = mat.Rotated(pixel.ZV, car.GetDirectionInRadians())
	mat = mat.Moved(pixel.V(car.Position.X, car.Position.Y))
	return mat
}

func GetCarGraphic(image string, car *Car) (*pixel.Sprite, pixel.Matrix) {
	carPic, err := GetCarImage(image)
	if err != nil {
		panic(err)
	}
	sprite := GetCarSprite(carPic)
	mat := GetCarMatrix(car)
	return sprite, mat
}