package colour

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

/*

basic draw function smpte ramps for example

colour inverter

*/

type bar struct {
	Space space
}

type bars struct {
	width float64
	color color.Color
}




func testrun() {
	/*

		mkae one image setting a test pattern

	*/
	noSpace := space{space: "nothing"}
	noChange := newNRGBA64(noSpace, image.Rect(0, 0, 1000, 1000))
	b := bar{Space: noSpace}
	b.generate(noChange)
	f, _ := os.Create("normal.png")
	png.Encode(f, noChange)

	changeSpace := space{space: "inverse"}
	change := newNRGBA64(noSpace, image.Rect(0, 0, 1000, 1000))
	cb := bar{Space: changeSpace}
	cb.generate(change)

	fc, _ := os.Create("inverse.png")
	png.Encode(fc, change)
}


