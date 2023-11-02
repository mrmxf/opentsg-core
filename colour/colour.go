package colour

import (
	"image"
	"image/color"
	"image/draw"
	"os"
)

// maxAlpha is the maximum color value returned by image.Color.RGBA.
const maxAlpha = 1<<16 - 1

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

	Draw(base, image.Rect(0, 0, 1000, 1000), noChange, image.Point{}, draw.Over)
	Draw(base, image.Rect(1000, 0, 2000, 1000), changeYCb, image.Point{}, draw.Over)
	Draw(base, image.Rect(0, 1000, 1000, 2000), chang601, image.Point{}, draw.Over)
	Draw(base, image.Rect(1000, 1000, 2000, 2000), chang709, image.Point{}, draw.Over)

	f, _ := os.Create("./testdata/all2.png")
	PngEncode(f, base)

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

	Draw(base, image.Rect(0, 0, 1000, 1000), noChange, image.Point{}, draw.Over)
	Draw(base, image.Rect(1000, 0, 2000, 1000), img709, image.Point{}, draw.Over)
	Draw(base, image.Rect(0, 1000, 1000, 2000), changP3, image.Point{}, draw.Over)
	Draw(base, image.Rect(1000, 1000, 2000, 2000), chang601, image.Point{}, draw.Over)

	f, _ := os.Create("./testdata/all2020.png")
	PngEncode(f, base)

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

	Draw(base, image.Rect(0, 0, 1000, 1000), noChange, image.Point{}, draw.Over)
	Draw(base, image.Rect(1000, 0, 2000, 1000), img709, image.Point{}, draw.Over)
	Draw(base, image.Rect(0, 1000, 1000, 2000), changP3, image.Point{}, draw.Over)
	Draw(base, image.Rect(1000, 1000, 2000, 2000), chang601, image.Point{}, draw.Over)

	f, _ := os.Create("./testdata/all709.png")
	PngEncode(f, base)

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

// Draw calls DrawMask with a nil mask.
func Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, op draw.Op) {
	DrawMask(dst, r, src, sp, nil, image.Point{}, op)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result of a Porter-Duff composition. A nil mask is treated as opaque.
/*
This version has been made to include the transform options for when colour space aware colours
are used. This works by using the transform function before placing the colour on the destination
image

Further more this uses NRGBA64 as a base to set colours, rather than the RGBA64 model
favoured by go. This means that the non alpha multiplied RGB values are used unless
required when alpha is neither 0 or maxAlpha. This leads to slight discrepancies
with the go value, but this is more accurate.

This function is only recommended when using NRGB64 images that are colour space aware and
you drawing with the same images.
*/
func DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, op draw.Op) {

	switch dst.(type) {
	case *NRGB64:

		// follow the draw.DrawMask generic code
		// with a few slight differences to ensure colour space is preserved.

		clip(dst, &r, src, &sp, mask, &mp)
		if r.Empty() {
			return
		}

		x0, x1, dx := r.Min.X, r.Max.X, 1
		y0, y1, dy := r.Min.Y, r.Max.Y, 1
		//set the colour to be colour space aware
		var out CNRGBA64
		sy := sp.Y + y0 - r.Min.Y
		my := mp.Y + y0 - r.Min.Y
		c := 0
		for y := y0; y != y1; y, sy, my = y+dy, sy+dy, my+dy {
			c++
			sx := sp.X + x0 - r.Min.X
			mx := mp.X + x0 - r.Min.X
			for x := x0; x != x1; x, sx, mx = x+dx, sx+dx, mx+dx {
				ma := uint32(maxAlpha)
				if mask != nil {
					_, _, _, ma = mask.At(mx, my).RGBA()
				}
				switch {
				//case op == draw.Over:
				// this differs from the go code
				// as it sets straight on top as does not change to rGBA64 like
				// teh draw.Drawmask function does
				//	dst.Set(x, y, src.At(sx, sy))

				case ma == 0:
					if op == draw.Over {
						// No-op.
					} else { // reset to a transparent pixel
						dst.Set(x, y, color.Transparent)
					}
				case ma == maxAlpha && op == draw.Src:
					dst.Set(x, y, src.At(sx, sy))
				default:

					/*


						we know the base is NRGBA64


					*/
					//	atc := color.NRGBA64Model.Convert(src.At(sx, sy)).(color.NRGBA64)
					srCol := src.At(sx, sy)
					var sr, sg, sb, sa uint32
					// this works in nrgb64, so we treat every colour as NRGB64
					if ncol, ok := srCol.(color.NRGBA64); ok {
						sr, sg, sb, sa = uint32(ncol.R), uint32(ncol.G), uint32(ncol.B), uint32(ncol.A)
					} else if cspace, ok := src.At(sx, sy).(*CNRGBA64); ok {

						// transform the colour before applying it
						tCol := transform(cspace.Space, dst.(*NRGB64).space, src.At(sx, sy))
						ncol := tCol.(*CNRGBA64)
						// making sure to cut out alpha multiplied values
						//	atc := tCol.(*CNRGBA64)
						//	sr, sg, sb, sa = uint32(atc.R), uint32(atc.G), uint32(atc.B), uint32(atc.A)
						sr, sg, sb, sa = uint32(ncol.R), uint32(ncol.G), uint32(ncol.B), uint32(ncol.A)
						// sr, sg, sb, sa = tCol.RGBA()
					} else {
						// convert these into non alpha multiplied values
						// this will be lossy so stick to NRGBA64 or CNRGBA64
						// if you want accurate values
						// @TODO figure out how to stop RGBA 64 causing overflow issues
						//	sr, sg, sb, sa = srCol.RGBA()

						nrgbCol := color.NRGBA64Model.Convert(srCol).(color.NRGBA64)

						sr, sg, sb, sa = uint32(nrgbCol.R), uint32(nrgbCol.G), uint32(nrgbCol.B), uint32(nrgbCol.A)
					}
					// NRGBA64 is non alpha multiplied
					/*
						if op == draw.Over {
							var tempout color.RGBA64
							dr, dg, db, da := dst.At(x, y).RGBA()
							a := maxAlpha - (sa * ma / maxAlpha)
							tempout.R = uint16((dr*a + sr*ma) / maxAlpha)
							tempout.G = uint16((dg*a + sg*ma) / maxAlpha)
							tempout.B = uint16((db*a + sb*ma) / maxAlpha)
							tempout.A = uint16((da*a + sa*ma) / maxAlpha)

							midOut := color.NRGBA64Model.Convert(tempout).(color.NRGBA64)
							out.R, out.G, out.B, out.A = midOut.R, midOut.G, midOut.B, midOut.A


						} else {
							out.R = uint16(sr * ma / maxAlpha)
							out.G = uint16(sg * ma / maxAlpha)
							out.B = uint16(sb * ma / maxAlpha)
							out.A = uint16(sa * ma / maxAlpha)
						}*/

					if op == draw.Src || sa == maxAlpha {
						out.R = uint16(sr * ma / maxAlpha)
						out.G = uint16(sg * ma / maxAlpha)
						out.B = uint16(sb * ma / maxAlpha)
						out.A = uint16(sa * ma / maxAlpha)

					} else {
						dstCol := dst.At(x, y).(*CNRGBA64)
						dr, dg, db, da := uint32(dstCol.R), uint32(dstCol.G), uint32(dstCol.B), uint32(dstCol.A)

						if da == 0 {
							out = CNRGBA64{R: uint16(sr), G: uint16(sg), B: uint16(sb), A: uint16(sa)}
						} else {

							// else get the alpha multiplied version of the dst and src RGBA values
							sr, sg, sb = ((sr * sa) / maxAlpha), ((sg * sa) / maxAlpha), ((sb * sa) / maxAlpha)
							dr, dg, db = ((dr * da) / maxAlpha), ((dg * da) / maxAlpha), ((db * da) / maxAlpha)

							// the alpha weighting can be changed
							// to a function for when different
							// clour spaces are usef
							a := maxAlpha - (sa * ma / maxAlpha)

							out.A = uint16((da*a + sa*ma) / maxAlpha)
							// divide by alpha to get th eno alpha multiplied version
							out.R = uint16((dr*a + sr*ma) / uint32(out.A))
							out.G = uint16((dg*a + sg*ma) / uint32(out.A))
							out.B = uint16((db*a + sb*ma) / uint32(out.A))

							// out.G = uint16((dg*da + sg*sa) / (da + sa))
							// out.B = uint16((db*da + sb*sa) / (da + sa))
							// out.A = uint16((da*a + sa*ma) / maxAlpha)

							//midOut := color.NRGBA64Model.Convert(tempout).(color.NRGBA64)
							//out.R, out.G, out.B, out.A = midOut.R, midOut.G, midOut.B, midOut.A

							/*
								var tempout color.RGBA64
								dr, dg, db, da := dst.At(x, y).RGBA()
								a := maxAlpha - (sa * ma / maxAlpha)
								tempout.R = uint16((dr*a + sr*ma) / maxAlpha)
								tempout.G = uint16((dg*a + sg*ma) / maxAlpha)
								tempout.B = uint16((db*a + sb*ma) / maxAlpha)
								tempout.A = uint16((da*a + sa*ma) / maxAlpha)

								midOut := color.NRGBA64Model.Convert(tempout).(color.NRGBA64)
								out.R, out.G, out.B, out.A = midOut.R, midOut.G, midOut.B, midOut.A*/
						}
					}

					// @TODO double check if this is needed
					// or double dipping transformations
					// assign the colour space if there is one
					if cspace, ok := src.At(sx, sy).(*CNRGBA64); ok {
						out.Space = cspace.Space
					}
					// The third argument is &out instead of out (and out is
					// declared outside of the inner loop) to avoid the implicit
					// conversion to color.Color here allocating memory in the
					// inner loop if sizeof(color.RGBA64) > sizeof(uintptr).

					dst.Set(x, y, &out)
				}
			}
		}

	default:
		draw.DrawMask(dst, r, src, sp, mask, mp, op)

	}

}

// clip clips r against each image's bounds (after translating into the
// destination image's coordinate space) and shifts the points sp and mp by
// the same amount as the change in r.Min.
// This the same as the draw standard library
func clip(dst draw.Image, r *image.Rectangle, src image.Image, sp *image.Point, mask image.Image, mp *image.Point) {
	orig := r.Min
	*r = r.Intersect(dst.Bounds())
	*r = r.Intersect(src.Bounds().Add(orig.Sub(*sp)))
	if mask != nil {
		*r = r.Intersect(mask.Bounds().Add(orig.Sub(*mp)))
	}
	dx := r.Min.X - orig.X
	dy := r.Min.Y - orig.Y
	if dx == 0 && dy == 0 {
		return
	}
	sp.X += dx
	sp.Y += dy
	if mp != nil {
		mp.X += dx
		mp.Y += dy
	}
}
