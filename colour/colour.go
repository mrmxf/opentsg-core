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

//

func testrun() {
	/*

		mkae one image setting a test pattern
	*/

	base := image.NewNRGBA64(image.Rect(0, 0, 2000, 2000))
	noSpace := space{space: "rec709"}
	noChange := newNRGBA64(noSpace, image.Rect(0, 0, 1000, 1000))
	b := bar{Space: noSpace}
	b.generate(noChange)

	changeSpace := space{space: "inverse"}
	change := newNRGBA64(noSpace, image.Rect(0, 0, 1000, 1000))
	cb := bar{Space: changeSpace}
	cb.generate(change)

	change601 := space{space: "rec601"}
	chang601 := newNRGBA64(noSpace, image.Rect(0, 0, 1000, 1000))
	cb601 := bar{Space: change601}
	cb601.generate(chang601)

	change709 := space{space: "test709"}
	chang709 := newNRGBA64(noSpace, image.Rect(0, 0, 1000, 1000))
	cb709 := bar{Space: change709}
	cb709.generate(chang709)

	draw.Draw(base, image.Rect(0, 0, 1000, 1000), noChange, image.Point{}, draw.Over)
	draw.Draw(base, image.Rect(1000, 0, 2000, 1000), noChange, image.Point{}, draw.Over)
	draw.Draw(base, image.Rect(0, 1000, 1000, 2000), chang601, image.Point{}, draw.Over)
	draw.Draw(base, image.Rect(1000, 1000, 2000, 2000), chang709, image.Point{}, draw.Over)

	f, _ := os.Create("./testdata/all.png")
	png.Encode(f, base)

}

func testrun2() {
	/*

		mkae one image setting a test pattern
	*/

	base := image.NewNRGBA64(image.Rect(0, 0, 2000, 2000))
	noSpace := space{space: "rec709"}
	noChange := newNRGBA642(noSpace, image.Rect(0, 0, 1000, 1000))
	b := bar{Space: noSpace}
	b.generate2(noChange)

	changeSpace := space{space: "rec709"}
	changeYCb := newNRGBA642(noSpace, image.Rect(0, 0, 1000, 1000))
	cb := bar{Space: changeSpace}
	cb.generateYCbCr(changeYCb)

	change601 := space{space: "rec601"}
	chang601 := newNRGBA642(noSpace, image.Rect(0, 0, 1000, 1000))
	cb601 := bar{Space: change601}
	cb601.generate2(chang601)

	change709 := space{space: "rec601"}
	chang709 := newNRGBA642(noSpace, image.Rect(0, 0, 1000, 1000))
	cb709 := bar{Space: change709}
	cb709.generateYCbCr(chang709)

	draw.Draw(base, image.Rect(0, 0, 1000, 1000), noChange, image.Point{}, draw.Over)
	draw.Draw(base, image.Rect(1000, 0, 2000, 1000), changeYCb, image.Point{}, draw.Over)
	draw.Draw(base, image.Rect(0, 1000, 1000, 2000), chang601, image.Point{}, draw.Over)
	draw.Draw(base, image.Rect(1000, 1000, 2000, 2000), chang709, image.Point{}, draw.Over)

	f, _ := os.Create("./testdata/all2.png")
	png.Encode(f, base)

}

func testrun2020() {
	/*

		mkae one image setting a test pattern
	*/

	base := image.NewNRGBA64(image.Rect(0, 0, 2000, 2000))
	noSpace := space{space: "rec2020"}
	noChange := newNRGBA642(noSpace, image.Rect(0, 0, 1000, 1000))
	b := bar{Space: noSpace}
	b.generate2(noChange)

	changeSpace := space{space: "rec709"}
	changeYCb := newNRGBA642(noSpace, image.Rect(0, 0, 1000, 1000))
	cb := bar{Space: changeSpace}
	cb.generate2(changeYCb)

	change601 := space{space: "rec601"}
	chang601 := newNRGBA642(noSpace, image.Rect(0, 0, 1000, 1000))
	cb601 := bar{Space: change601}
	cb601.generate2(chang601)

	change709 := space{space: "p3"}
	chang709 := newNRGBA642(noSpace, image.Rect(0, 0, 1000, 1000))
	cb709 := bar{Space: change709}
	cb709.generate2(chang709)

	draw.Draw(base, image.Rect(0, 0, 1000, 1000), noChange, image.Point{}, draw.Over)
	draw.Draw(base, image.Rect(1000, 0, 2000, 1000), changeYCb, image.Point{}, draw.Over)
	draw.Draw(base, image.Rect(0, 1000, 1000, 2000), chang601, image.Point{}, draw.Over)
	draw.Draw(base, image.Rect(1000, 1000, 2000, 2000), chang709, image.Point{}, draw.Over)

	f, _ := os.Create("./testdata/all2020.png")
	png.Encode(f, base)

}

type CNRGBA64 struct {
	// color.NRGBA64
	R, G, B, A uint16
	Space      space
}

type ColorSpace interface {
	color.Color
	GetSpace() space
	UpdateSpace(space)
}

func (c *CNRGBA64) GetSpace() space {
	return c.Space
}

func (c *CNRGBA64) UpdateSpace(s space) {
	c.Space = s
}

func (c *CNRGBA64) RGBA() (R, G, B, A uint32) {
	return color.NRGBA64{R: c.R, G: c.G, B: c.B, A: c.A}.RGBA()
}

type CYCbCr struct {
	Y, Cb, Cr uint8
	Space     space
}

func (c *CYCbCr) GetSpace() space {
	return c.Space
}

func (c *CYCbCr) UpdateSpace(s space) {
	c.Space = s
}

func (c *CYCbCr) RGBA() (R, G, B, A uint32) {
	return color.YCbCr{Y: c.Y, Cb: c.Cb, Cr: c.Cr}.RGBA()
}
