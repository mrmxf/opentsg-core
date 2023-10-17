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

/*
func testrun() {
	/*

		mkae one image setting a test pattern


	base := image.NewNRGBA64(image.Rect(0, 0, 2000, 2000))
	noSpace := Space{Space: "rec709"}
	noChange := newNRGBA642(noSpace, image.Rect(0, 0, 1000, 1000))
	b := bar{Space: noSpace}
	b.generate(noChange)

	changeSpace := Space{Space: "inverse"}
	change := newNRGBA642(noSpace, image.Rect(0, 0, 1000, 1000))
	cb := bar{Space: changeSpace}
	cb.generate(change)

	change601 := Space{Space: "rec601"}
	chang601 := newNRGBA642(noSpace, image.Rect(0, 0, 1000, 1000))
	cb601 := bar{Space: change601}
	cb601.generate(chang601)

	change709 := Space{Space: "test709"}
	chang709 := newNRGBA642(noSpace, image.Rect(0, 0, 1000, 1000))
	cb709 := bar{Space: change709}
	cb709.generate(chang709)

	draw.Draw(base, image.Rect(0, 0, 1000, 1000), noChange, image.Point{}, draw.Over)
	draw.Draw(base, image.Rect(1000, 0, 2000, 1000), noChange, image.Point{}, draw.Over)
	draw.Draw(base, image.Rect(0, 1000, 1000, 2000), chang601, image.Point{}, draw.Over)
	draw.Draw(base, image.Rect(1000, 1000, 2000, 2000), chang709, image.Point{}, draw.Over)

	f, _ := os.Create("./testdata/all.png")
	png.Encode(f, base)

}*/

func testrun2() {
	/*

		mkae one image setting a test pattern
	*/

	base := image.NewNRGBA64(image.Rect(0, 0, 2000, 2000))
	noSpace := ColorSpace{ColorSpace: "rec709"}
	noChange := NewNRGBA64(noSpace, image.Rect(0, 0, 1000, 1000))
	b := bar{Space: noSpace}
	b.generate2(noChange)

	changeSpace := ColorSpace{ColorSpace: "rec709"}
	changeYCb := NewNRGBA64(noSpace, image.Rect(0, 0, 1000, 1000))
	cb := bar{Space: changeSpace}
	cb.generateYCbCr(changeYCb)

	change601 := ColorSpace{ColorSpace: "rec601"}
	chang601 := NewNRGBA64(noSpace, image.Rect(0, 0, 1000, 1000))
	cb601 := bar{Space: change601}
	cb601.generate2(chang601)

	change709 := ColorSpace{ColorSpace: "rec601"}
	chang709 := NewNRGBA64(noSpace, image.Rect(0, 0, 1000, 1000))
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
	noSpace := ColorSpace{ColorSpace: "rec2020"}
	noChange := NewNRGBA64(noSpace, image.Rect(0, 0, 1000, 1000))
	b := bar{Space: noSpace}
	b.generate2(noChange)

	changeSpace := ColorSpace{ColorSpace: "rec709"}
	img709 := NewNRGBA64(noSpace, image.Rect(0, 0, 1000, 1000))
	cb := bar{Space: changeSpace}
	cb.generate2(img709)

	change601 := ColorSpace{ColorSpace: "rec601"}
	chang601 := NewNRGBA64(noSpace, image.Rect(0, 0, 1000, 1000))
	cb601 := bar{Space: change601}
	cb601.generate2(chang601)

	change709 := ColorSpace{ColorSpace: "p3"}
	changP3 := NewNRGBA64(noSpace, image.Rect(0, 0, 1000, 1000))
	cb709 := bar{Space: change709}
	cb709.generate2(changP3)

	draw.Draw(base, image.Rect(0, 0, 1000, 1000), noChange, image.Point{}, draw.Over)
	draw.Draw(base, image.Rect(1000, 0, 2000, 1000), img709, image.Point{}, draw.Over)
	draw.Draw(base, image.Rect(0, 1000, 1000, 2000), changP3, image.Point{}, draw.Over)
	draw.Draw(base, image.Rect(1000, 1000, 2000, 2000), chang601, image.Point{}, draw.Over)

	f, _ := os.Create("./testdata/all2020.png")
	png.Encode(f, base)

}

func testrun709() {
	/*

		mkae one image setting a test pattern
	*/

	base := image.NewNRGBA64(image.Rect(0, 0, 2000, 2000))
	noSpace := ColorSpace{ColorSpace: "rec709"}
	noChange := NewNRGBA64(noSpace, image.Rect(0, 0, 1000, 1000))
	b := bar{Space: ColorSpace{ColorSpace: "rec2020"}}
	b.generate2(noChange)

	changeSpace := ColorSpace{ColorSpace: "rec709"}
	img709 := NewNRGBA64(noSpace, image.Rect(0, 0, 1000, 1000))
	cb := bar{Space: changeSpace}
	cb.generate2(img709)

	change601 := ColorSpace{ColorSpace: "rec601"}
	chang601 := NewNRGBA64(noSpace, image.Rect(0, 0, 1000, 1000))
	cb601 := bar{Space: change601}
	cb601.generate2(chang601)

	change709 := ColorSpace{ColorSpace: "p3"}
	changP3 := NewNRGBA64(noSpace, image.Rect(0, 0, 1000, 1000))
	cb709 := bar{Space: change709}
	cb709.generate2(changP3)

	draw.Draw(base, image.Rect(0, 0, 1000, 1000), noChange, image.Point{}, draw.Over)
	draw.Draw(base, image.Rect(1000, 0, 2000, 1000), img709, image.Point{}, draw.Over)
	draw.Draw(base, image.Rect(0, 1000, 1000, 2000), changP3, image.Point{}, draw.Over)
	draw.Draw(base, image.Rect(1000, 1000, 2000, 2000), chang601, image.Point{}, draw.Over)

	f, _ := os.Create("./testdata/all709.png")
	png.Encode(f, base)

}

type CNRGBA64 struct {
	// color.NRGBA64
	R, G, B, A uint16
	Space      ColorSpace
}

type Color interface {
	color.Color
	GetColorSpace() ColorSpace
	UpdateColorSpace(ColorSpace)
}

func (c *CNRGBA64) GetColorSpace() ColorSpace {
	return c.Space
}

func (c *CNRGBA64) UpdateColorSpace(s ColorSpace) {
	c.Space = s
}

func (c *CNRGBA64) RGBA() (R, G, B, A uint32) {
	return color.NRGBA64{R: c.R, G: c.G, B: c.B, A: c.A}.RGBA()
}

type CYCbCr struct {
	Y, Cb, Cr uint8
	Space     ColorSpace
}

func (c *CYCbCr) GetSpace() ColorSpace {
	return c.Space
}

func (c *CYCbCr) UpdateSpace(s ColorSpace) {
	c.Space = s
}

func (c *CYCbCr) RGBA() (R, G, B, A uint32) {
	return color.YCbCr{Y: c.Y, Cb: c.Cb, Cr: c.Cr}.RGBA()
}
