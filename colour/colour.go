package colour

import (
	"image"
	"image/color"
	"image/draw"
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

//

func testrun() {
	/*

		mkae one image setting a test pattern
	*/

	base := image.NewNRGBA64(image.Rect(0, 0, 2000, 2000))
	noSpace := space{space: "nothing"}
	noChange := newNRGBA64(noSpace, image.Rect(0, 0, 1000, 1000))
	b := bar{Space: noSpace}
	b.generate(noChange)

	changeSpace := space{space: "inverse"}
	change := newNRGBA64(noSpace, image.Rect(0, 0, 1000, 1000))
	cb := bar{Space: changeSpace}
	cb.generate(change)

	draw.Draw(base, image.Rect(0, 0, 1000, 1000), noChange, image.Point{}, draw.Over)
	draw.Draw(base, image.Rect(1000, 0, 2000, 1000), change, image.Point{}, draw.Over)

	f, _ := os.Create("./testdata/all.png")
	png.Encode(f, base)

}
