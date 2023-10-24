package colour

import (
	"crypto/sha256"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math/rand"
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

func TestDraw(t *testing.T) {

	/*
		tests - see draw works the same when no colour colour space is applied

		check transformations of some small squares

	*/

	// check for any deviations from go
	for i := 0; i < 5; i++ {

		baseColour := color.NRGBA64{R: uint16(rand.Int63n(65535)), G: uint16(rand.Int63n(65535)), B: uint16(rand.Int63n(65535)), A: uint16(rand.Int63n(65535))}

		colourImplementation := NewNRGBA64(ColorSpace{ColorSpace: "rec2020"}, image.Rect(0, 0, 1000, 1000))
		Draw(colourImplementation, colourImplementation.Bounds(), &image.Uniform{baseColour}, image.Point{}, draw.Src)

		goImplementation := NewNRGBA64(ColorSpace{ColorSpace: "rec2020"}, image.Rect(0, 0, 1000, 1000))
		draw.Draw(goImplementation, goImplementation.Bounds(), &image.Uniform{baseColour}, image.Point{}, draw.Src)

		hnormal := sha256.New()
		htest := sha256.New()
		hnormal.Write(goImplementation.Pix())
		htest.Write(colourImplementation.Pix())

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

	/*
		f, _ := os.Create("./testdata/colour.png")
		png.Encode(f, base)

		basedraw := NewNRGBA64(ColorSpace{ColorSpace: "rec2020"}, image.Rect(0, 0, 1000, 1000))

		Draw(basedraw, base.Bounds(), &image.Uniform{color.NRGBA64{R: 0xffff, A: 0xfff0}}, image.Point{}, draw.Over)

		fdraw, _ := os.Create("./testdata/coloudrawr.png")
		png.Encode(fdraw, basedraw)*/

	// set some base test transformations
}
