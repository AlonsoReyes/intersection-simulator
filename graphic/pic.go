package graphic

import (
	"github.com/faiface/pixel"
	"image"
	"os"
)


func SpriteSize(pd *pixel.Sprite) (float64, float64) {
	b := pd.Frame()
	width := b.W()
	heigth := b.H()
	return float64(width), float64(heigth)
}

func LoadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}