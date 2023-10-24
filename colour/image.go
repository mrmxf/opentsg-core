package colour

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"

	"golang.org/x/image/tiff"
)

/*
ColorSpace will need some methods as transforms get more complicated
*/
type ColorSpace struct {
	ColorSpace string `json:"ColorSpace,omitempty" yaml:"ColorSpace,omitempty"`
	// Primaries let the space be declared as a string and the primaries
	// be sued for generating transformation matrices.
	// Have two seperate maps for data points
	TransformType string    `json:"TransformType,omitempty" yaml:"TransformType,omitempty"`
	Primaries     Primaries `json:"Primaries,omitempty" yaml:"Primaries,omitempty"`
}

type Primaries struct {
	Red        XY `json:"Red,omitempty" yaml:"Red,omitempty"`
	Green      XY `json:"Green,omitempty" yaml:"Green,omitempty"`
	Blue       XY `json:"Blue,omitempty" yaml:"Blue,omitempty"`
	WhitePoint XY `json:"WhitePoint,omitempty" yaml:"WhitePoint,omitempty"`
}

type XY struct {
	X int `json:"X,omitempty" yaml:"X,omitempty"`
	Y int `json:"Y,omitempty" yaml:"Y,omitempty"`
}

type Image interface {
	Space() ColorSpace
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
	space ColorSpace
}

/*
wrap all the image functions with NRGB642


My current idea is using interface for NRGBa for our own colour tyoe that applies the transformation ar each set
replacing the interface with one that returns the colour space for prosperity
*/

func NewNRGBA64(s ColorSpace, r image.Rectangle) *NRGB64 {

	base := image.NewNRGBA64(r)

	return &NRGB64{base: base, space: s}

}

func (n NRGB64) Bounds() image.Rectangle {
	return n.base.Bounds()
}

func (n NRGB64) Space() ColorSpace {
	return n.space
}

func (n NRGB64) Pix() []uint8 {
	return n.base.Pix
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

func (n NRGB64) BaseImage() *image.NRGBA64 {
	return n.base
}

// utilise set for draw
func (n NRGB64) Set(x int, y int, c color.Color) {

	// update the colour if it has an explicit colour space
	// and the base image is using colour spaces
	if cmid, ok := c.(Color); ok && (n.space != ColorSpace{}) {
		c = transform(cmid.GetColorSpace(), n.space, c)
	}

	n.base.Set(x, y, c)
}

/*
use draw and set

because the draw removes the colour transformation from set.

The only issue is add image, but this will be a beast unto itself

*/

func PngEncode(w io.Writer, m image.Image) error {

	// cut out the NRGB64 wrapper as png
	// doesn't know how to handle it correctly
	// and it changes the expected values when alpha is not 0xffff
	if mid, ok := m.(*NRGB64); ok {
		m = mid.base
	}

	return png.Encode(w, m)
}

func TiffEncode(w io.Writer, m image.Image, opt *tiff.Options) error {

	// cut out the NRGB64 wrapper as tiff
	// doesn't know how to handle it correctly
	// and it changes the expected values when alpha is not 0xffff
	if mid, ok := m.(*NRGB64); ok {
		m = mid.base
	}

	return tiff.Encode(w, m, opt)
}
