package colour

import (
	"image"
	"image/draw"
)

type space struct {
	space string
}

type Image interface {
	Space() space
	draw.Image
}

type nrgb struct {
	*image.NRGBA64
	Spacer space
}

func (n nrgb) Space() space {
	return n.Spacer
}

func newNRGBA64(s space, r image.Rectangle) Image {

	base := image.NewNRGBA64(r)

	return &nrgb{base, s}

}
