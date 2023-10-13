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
/*
the colour should awlays be transformed before Set()
e.g. by running canvas.Transform()

You do not want to double transform.

What uses set.

These are special cases can we doing something tight for them
Add image
any textbox

*/

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

type NRGB642 struct {
	Base   *image.NRGBA64
	spacer space
}

/*
wrap all the image functions with NRGB642


My current idea is using interface for NRGBa for our own colour tyoe that applies the transformation ar each set
replacing the interface with one that returns the colour space for prosperity
*/

func newNRGBA642(s space, r image.Rectangle) draw.Image {

	base := image.NewNRGBA64(r)

	return &NRGB642{Base: base, spacer: s}

}

func (n NRGB642) Bounds() image.Rectangle {
	return n.Base.Bounds()
}

func (n NRGB642) At(x, y int) color.Color {
	return n.Base.At(x, y)

}

func (n NRGB642) ColorModel() color.Model {
	return n.Base.ColorModel()
}


// utilise set for draw
func (n NRGB642) Set(x int, y int, c color.Color) {

	// update the colour if it has an explicit colour space
	if cmid, ok := c.(ColorSpace); ok {
		c = transform(cmid.GetSpace(), n.spacer, c)
	}

	n.Base.Set(x, y, c)
}
/*
use draw and set

because the draw removes the colour transformation from set.

The only issue is add image, but this will be a beast unto itself

*/