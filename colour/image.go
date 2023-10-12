package colour

import (
	"image"
	"image/color"
	"image/draw"
)

/*
space will need some methods as transforms get more complicated
*/
type space struct {
	space string
}

type Image interface {
	Draw(image.Rectangle, color.Color, draw.Op, space)
	draw.Image // Draw include Set
}

// font needs draw.Image
// include the framework to do it properly
// and hope people just do it?

type WidgetImage interface {
	Draw(image.Rectangle, color.Color, draw.Op, space)
	image.Image // image does not include yet
}

// replicate this but for Aces
type NRGB64 struct {
	*image.NRGBA64
	spacer space
}

// Draw wraps image.Draw
func (n NRGB64) Draw(r image.Rectangle, src color.Color, op draw.Op, input space) {

	ts := transform(input, n.spacer, src)
	//ts := transformer(src, n.Spacer)

	draw.Draw(n, r, &image.Uniform{ts}, image.Point{}, op)

}



// This wraps th eimage library adding colourspaces
// to the image. While including all th ebasic image
// functionality
func newNRGBA64(s space, r image.Rectangle) Image {

	base := image.NewNRGBA64(r)

	return &NRGB64{base, s}

}
