package colour

import (
	"crypto/sha256"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math/rand"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTemp(t *testing.T) {
	//	testrun()
	testrun2()
	testrun2020()
	testrun709()

}

// go test ./colour/ -bench=. -benchtime=10s

/*
func BenchmarkNRGBA64Area(b *testing.B) {
	// decode to get the colour values
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		testrun2020()
	}
}

func BenchmarkNRGBA64ACESSet(b *testing.B) {
	// decode to get the colour values

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		testrun709()
	}
} */

func BenchmarkNRGBA64Draw(b *testing.B) {
	base := NewNRGBA64(ColorSpace{ColorSpace: "rec2020"}, image.Rect(0, 0, 1000, 1000))
	for n := 0; n < b.N; n++ {
		Draw(base, base.Bounds(), &image.Uniform{&CNRGBA64{R: 0xffff, A: 0xfff0, Space: ColorSpace{ColorSpace: "rec709"}}}, image.Point{}, draw.Src)
	}
}

func BenchmarkNRGBA64ImgDraw(b *testing.B) {
	base := NewNRGBA64(ColorSpace{ColorSpace: "rec2020"}, image.Rect(0, 0, 1000, 1000))
	for n := 0; n < b.N; n++ {
		draw.Draw(base, base.Bounds(), &image.Uniform{&CNRGBA64{R: 0xffff, A: 0xfff0, Space: ColorSpace{ColorSpace: "rec709"}}}, image.Point{}, draw.Src)
	}
}

func BenchmarkNRGBA64DrawMaxAlpha(b *testing.B) {
	base := NewNRGBA64(ColorSpace{ColorSpace: "rec2020"}, image.Rect(0, 0, 1000, 1000))
	for n := 0; n < b.N; n++ {
		Draw(base, base.Bounds(), &image.Uniform{&CNRGBA64{R: 0xffff, A: 0xffff, Space: ColorSpace{ColorSpace: "rec709"}}}, image.Point{}, draw.Src)
	}
}

func BenchmarkNRGBA64ImgDrawMaxAlpha(b *testing.B) {
	base := NewNRGBA64(ColorSpace{ColorSpace: "rec2020"}, image.Rect(0, 0, 1000, 1000))
	for n := 0; n < b.N; n++ {
		draw.Draw(base, base.Bounds(), &image.Uniform{&CNRGBA64{R: 0xffff, A: 0xffff, Space: ColorSpace{ColorSpace: "rec709"}}}, image.Point{}, draw.Src)
	}
}

func TestDraw(t *testing.T) {

	/*
		tests - see draw works the same when no colour colour space is applied

		check transformations of some small squares

	*/

	// check for any deviations from go
	for i := 0; i < 5; i++ {

		baseColour := color.NRGBA64{R: uint16(rand.Int63n(65535)), G: uint16(rand.Int63n(65535)), B: uint16(rand.Int63n(65535)), A: uint16(rand.Int63n(65535))}

		colourImplementation := NewNRGBA64(ColorSpace{ColorSpace: "rec2020"}, image.Rect(0, 0, 20, 20))
		Draw(colourImplementation, colourImplementation.Bounds(), &image.Uniform{baseColour}, image.Point{}, draw.Src)

		goImplementation := NewNRGBA64(ColorSpace{ColorSpace: "rec2020"}, image.Rect(0, 0, 20, 20))
		draw.Draw(goImplementation, goImplementation.Bounds(), &image.Uniform{baseColour}, image.Point{}, draw.Src)

		hnormal := sha256.New()
		htest := sha256.New()
		hnormal.Write(goImplementation.Pix())
		htest.Write(colourImplementation.Pix())

		fmt.Println(colourImplementation.At(0, 0))
		fmt.Println(goImplementation.At(0, 0))

		//td, _ := os.Create("r.png")
		//png.Encode(td, canvas)

		Convey("Checking that the go and colour implementations of draw produce the same result, when no colour space is involved", t, func() {
			Convey(fmt.Sprintf("Run using a colour of %v", baseColour), func() {
				Convey("The hashes of the image are identical", func() {
					So(htest.Sum(nil), ShouldResemble, hnormal.Sum(nil))
				})
			})
		})
	}

	colourImplementation := NewNRGBA64(ColorSpace{}, image.Rect(0, 0, 200, 200))
	goImplementation := NewNRGBA64(ColorSpace{}, image.Rect(0, 0, 200, 200))
	colours := []color.Color{
		color.RGBA64{R: 0x7FFF, A: 0x7FFF},
		&CNRGBA64{G: 0x7FFF, A: 0x7FFF},
		&CNRGBA64{B: 0x7FFF, A: 0x7FFF},
		&CNRGBA64{R: 0x6400, G: 0x6400, A: 16384},
		&CNRGBA64{R: uint16(rand.Int63n(65535)), G: uint16(rand.Int63n(65535)), B: uint16(rand.Int63n(65535)), A: uint16(rand.Int63n(65535))},
		&CNRGBA64{R: uint16(rand.Int63n(65535)), G: uint16(rand.Int63n(65535)), B: uint16(rand.Int63n(65535)), A: uint16(rand.Int63n(65535))},
		&CNRGBA64{R: uint16(rand.Int63n(65535)), G: uint16(rand.Int63n(65535)), B: uint16(rand.Int63n(65535)), A: uint16(rand.Int63n(65535))},
		&CNRGBA64{R: uint16(rand.Int63n(65535)), G: uint16(rand.Int63n(65535)), B: uint16(rand.Int63n(65535)), A: uint16(rand.Int63n(65535))},
		&CNRGBA64{R: uint16(rand.Int63n(65535)), G: uint16(rand.Int63n(65535)), B: uint16(rand.Int63n(65535)), A: uint16(rand.Int63n(65535))},
		&CNRGBA64{R: uint16(rand.Int63n(65535)), G: uint16(rand.Int63n(65535)), B: uint16(rand.Int63n(65535)), A: uint16(rand.Int63n(65535))},
		&CNRGBA64{R: uint16(rand.Int63n(65535)), G: uint16(rand.Int63n(65535)), B: uint16(rand.Int63n(65535)), A: uint16(rand.Int63n(65535))},
	}
	// check for any deviations from go
	for i := 0; i < 8; i++ {

		// baseColour := CNRGBA64{R: uint16(rand.Int63n(65535)), G: uint16(rand.Int63n(65535)), B: uint16(rand.Int63n(65535)), A: uint16(rand.Int63n(65535))}
		baseColour := colours[i]
		//	sr, sg, sb, sa := uint32(baseColour.R), uint32(baseColour.G), uint32(baseColour.B), uint32(baseColour.A)
		//	fmt.Println(sr, sg, sb, sa)
		//	sr, sg, sb = ((sr * sa) / maxAlpha), ((sg * sa) / maxAlpha), ((sb * sa) / maxAlpha)
		//	fmt.Println(sr, sg, sb, sa)

		nrgbCol := color.NRGBA64Model.Convert(baseColour).(color.NRGBA64)
		if (baseColour == color.RGBA64{R: 0x7FFF, A: 0x7FFF}) {
			fmt.Println("trigger")
		}

		sr, sg, sb, sa := uint32(nrgbCol.R), uint32(nrgbCol.G), uint32(nrgbCol.B), uint32(nrgbCol.A)
		sr, sg, sb = ((sr * sa) / maxAlpha), ((sg * sa) / maxAlpha), ((sb * sa) / maxAlpha)
		//	sr, sg, sb, sa := baseColour.RGBA()
		var out color.NRGBA64
		ma := uint32(0xffff)
		dstCol := colourImplementation.At(0, 0).(*CNRGBA64)
		dr, dg, db, da := uint32(dstCol.R), uint32(dstCol.G), uint32(dstCol.B), uint32(dstCol.A)
		dr, dg, db = ((dr * da) / maxAlpha), ((dg * da) / maxAlpha), ((db * da) / maxAlpha)
		a := maxAlpha - (sa * ma / maxAlpha)
		out.A = uint16((da*a + sa*ma) / maxAlpha)
		out.R = uint16((dr*a + sr*ma) / uint32(out.A))
		out.G = uint16((dg*a + sg*ma) / uint32(out.A))
		out.B = uint16((db*a + sb*ma) / uint32(out.A))
		fmt.Println(dr, dg, db, da, a)
		fmt.Println(((dr*a + sr*ma) / uint32(out.A)))
		fmt.Println(sr, sg, sb, sa)
		fmt.Println(out)

		Draw(colourImplementation, colourImplementation.Bounds(), &image.Uniform{baseColour}, image.Point{}, draw.Over)

		draw.Draw(goImplementation, goImplementation.Bounds(), &image.Uniform{baseColour}, image.Point{}, draw.Over)

		hnormal := sha256.New()
		htest := sha256.New()
		hnormal.Write(goImplementation.Pix())
		htest.Write(colourImplementation.Pix())
		fmt.Println(colourImplementation.At(0, 0), "my version")
		fmt.Println(goImplementation.At(0, 0))
		fmt.Println(colourImplementation.At(0, 0).RGBA()) //, "my version")
		fmt.Println(goImplementation.At(0, 0).RGBA())

		td, _ := os.Create(fmt.Sprintf("./testdata/draw/me%v.png", i))
		PngEncode(td, colourImplementation)
		td2, _ := os.Create(fmt.Sprintf("./testdata/draw/go%v.png", i))
		PngEncode(td2, goImplementation)

		Convey("Checking that the go and colour implementations of draw produce the same result, when no colour space is involved", t, func() {
			Convey(fmt.Sprintf("Run using a colour of %v", baseColour), func() {
				Convey("The hashes of the image are identical", func() {
					So(colourImplementation.At(0, 0), ShouldResemble, &baseColour)
				})
			})
		})
	}

	/*
		testColours := []CNRGBA64{{R: 35340, A: 0xffff}, {G: 30000, B: 40000, A: 0xf0f0}, {R: 0xffff, G: 0xffff, B: 0xffff}}

		target := []string{"fullalpha.png", "partialalpha.png", "noalpha.png"}

		for i, tcol := range testColours {
			colourImplementation := NewNRGBA64(ColorSpace{ColorSpace: "rec2020"}, image.Rect(0, 0, 1000, 1000))
			Draw(colourImplementation, colourImplementation.Bounds(), &image.Uniform{&tcol}, image.Point{}, draw.Src)
			//	goImplementation := image.NewNRGBA64(image.Rect(0, 0, 1000, 1000))
			//Draw(goImplementation, goImplementation.Bounds(), &image.Uniform{&tcol}, image.Point{}, draw.Src)

			baseFile, _ := os.Open("./testdata/draw/" + target[i])

			//PngEncode(basePng, colourImplementation.base)
			baseImage, _ := png.Decode(baseFile)

			testFormat := image.NewNRGBA64(baseImage.Bounds())
			Draw(testFormat, testFormat.Bounds(), baseImage, image.Point{}, draw.Src)
			//
			hnormal := sha256.New()
			htest := sha256.New()
			hnormal.Write(testFormat.Pix)
			htest.Write(colourImplementation.Pix())

			fmt.Println(baseImage.At(0, 0).RGBA())
			fmt.Println(testFormat.At(0, 0).RGBA())
			fmt.Println(colourImplementation.At(0, 0).RGBA())

			Convey("Checking that the transformation produces the expected results", t, func() {
				Convey(fmt.Sprintf("Run checking %v", target[i]), func() {
					Convey("The hashes of the image are identical", func() {
						So(htest.Sum(nil), ShouldResemble, hnormal.Sum(nil))
					})
				})
			})
		}

		base := NewNRGBA64(ColorSpace{}, image.Rect(0, 0, 2000, 1000))

		fmt.Println(base.At(500, 500))
		//base.Set(500, 500, &colour.CNRGBA64{R: 65335, A: 0xffff, Space: colour.ColorSpace{ColorSpace: "rec709"}})

		Draw(base, image.Rect(400, 400, 600, 600), &image.Uniform{&CNRGBA64{R: 65335, A: 0xffff, Space: ColorSpace{ColorSpace: "rec709"}}}, image.Point{}, draw.Src)
		fmt.Println(base.At(500, 500))*/
	/*
		f, _ := os.Create("./testdata/colour.png")
		png.Encode(f, base)

		basedraw := NewNRGBA64(ColorSpace{ColorSpace: "rec2020"}, image.Rect(0, 0, 1000, 1000))

		Draw(basedraw, base.Bounds(), &image.Uniform{color.NRGBA64{R: 0xffff, A: 0xfff0}}, image.Point{}, draw.Over)

		fdraw, _ := os.Create("./testdata/coloudrawr.png")
		png.Encode(fdraw, basedraw)*/

	// set some base test transformations
}
