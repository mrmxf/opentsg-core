package colour

import (
	"image"
	"image/color"
	"image/draw"
)

/*
Space will need some methods as transforms get more complicated
*/
type Space struct {
	Space string
	// Primaries let the space be declared as a string and the primaries
	// be sued for generating transformation matrices.
	// Have two seperate maps for data points
	TransformType string
	Primaries     Primaries
}

type Primaries struct {
	Red, Green, Blue, WhitePoint XY
}

type XY struct {
	X, Y int
}

type Image interface {
	Space() Space
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

type NRGB64 struct {
	base  *image.NRGBA64
	space Space
}

/*
wrap all the image functions with NRGB642


My current idea is using interface for NRGBa for our own colour tyoe that applies the transformation ar each set
replacing the interface with one that returns the colour space for prosperity
*/

func NewNRGBA64(s Space, r image.Rectangle) Image {

	base := image.NewNRGBA64(r)

	return &NRGB64{base: base, space: s}

}

func (n NRGB64) Bounds() image.Rectangle {
	return n.base.Bounds()
}

func (n NRGB64) Space() Space {
	return n.space
}

func (n NRGB64) At(x, y int) color.Color {
	/* can wrap
		NRGBA 64 colour as an tsg.colour
		if we want to preserve colour space
	 //	n.Base.NRGBA64At(x,y)

	*/

	baseCol := n.base.NRGBA64At(x, y)
	// return a colour space aware colour
	return &CNRGBA64{R: baseCol.R, G: baseCol.G, B: baseCol.B, A: baseCol.A, Space: n.space}

}

func (n NRGB64) ColorModel() color.Model {
	return n.base.ColorModel()
}

// utilise set for draw
func (n NRGB64) Set(x int, y int, c color.Color) {

	// update the colour if it has an explicit colour space
	if cmid, ok := c.(ColorSpace); ok {
		c = transform(cmid.GetSpace(), n.space, c)
	}

	n.base.Set(x, y, c)
}

/*
use draw and set

because the draw removes the colour transformation from set.

The only issue is add image, but this will be a beast unto itself

*/
