package colour

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
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

	Draw(base, image.Rect(0, 0, 1000, 1000), noChange, image.Point{}, draw.Over)
	Draw(base, image.Rect(1000, 0, 2000, 1000), img709, image.Point{}, draw.Over)
	Draw(base, image.Rect(0, 1000, 1000, 2000), changP3, image.Point{}, draw.Over)
	Draw(base, image.Rect(1000, 1000, 2000, 2000), chang601, image.Point{}, draw.Over)

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

	Draw(base, image.Rect(0, 0, 1000, 1000), noChange, image.Point{}, draw.Over)
	Draw(base, image.Rect(1000, 0, 2000, 1000), img709, image.Point{}, draw.Over)
	Draw(base, image.Rect(0, 1000, 1000, 2000), changP3, image.Point{}, draw.Over)
	Draw(base, image.Rect(1000, 1000, 2000, 2000), chang601, image.Point{}, draw.Over)

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

// Draw calls DrawMask with a nil mask.
func Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, op draw.Op) {
	DrawMask(dst, r, src, sp, nil, image.Point{}, op)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result of a Porter-Duff composition. A nil mask is treated as opaque.
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
		for y := y0; y != y1; y, sy, my = y+dy, sy+dy, my+dy {
			sx := sp.X + x0 - r.Min.X
			mx := mp.X + x0 - r.Min.X
			for x := x0; x != x1; x, sx, mx = x+dx, sx+dx, mx+dx {
				ma := uint32(maxAlpha)
				if mask != nil {
					_, _, _, ma = mask.At(mx, my).RGBA()
				}
				switch {
				case ma == 0:
					if op == draw.Over {
						// No-op.
					} else { // reset to a transparent pixel
						dst.Set(x, y, color.Transparent)
					}
				case ma == maxAlpha && op == draw.Src:
					dst.Set(x, y, src.At(sx, sy))
				default:
					sr, sg, sb, sa := src.At(sx, sy).RGBA()
					if cspace, ok := src.At(sx, sy).(*CNRGBA64); ok {
						// transform the colour before applying it
						tCol := transform(cspace.Space, dst.(*NRGB64).space, src.At(sx, sy))
						sr, sg, sb, sa = tCol.RGBA()
					}

					if op == draw.Over {

						dr, dg, db, da := dst.At(x, y).RGBA()

						a := maxAlpha - (sa * ma / maxAlpha)
						out.R = uint16((dr*a + sr*ma) / maxAlpha)
						out.G = uint16((dg*a + sg*ma) / maxAlpha)
						out.B = uint16((db*a + sb*ma) / maxAlpha)
						out.A = uint16((da*a + sa*ma) / maxAlpha)
					} else {
						out.R = uint16(sr * ma / maxAlpha)
						out.G = uint16(sg * ma / maxAlpha)
						out.B = uint16(sb * ma / maxAlpha)
						out.A = uint16(sa * ma / maxAlpha)
					}

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
