package colour

import (
	"fmt"
	"image/color"
	"math"
)

var (
	Library = map[space]map[space]func(color.Color) color.Color{
		{space: "inverse"}: {
			space{space: "rec709"}: inverse},
	}
	// rec 601 here http://www.brucelindbloom.com/index.html?WorkingSpaceInfo.html under PAL/SECAM RGB
	RGBToXYZ = map[string][3][3]float64{
		"rec709":  {{0.4124564, 0.3575761, 0.1804375}, {0.2126729, 0.7151522, 0.0721750}, {0.0193339, 0.1191920, 0.9503041}},
		"rec601":  {{0.4306190, 0.3415419, 0.1783091}, {0.2220379, 0.7066384, 0.0713236}, {0.0201853, 0.1295504, 0.9390944}},
		"test709": {{0.41239079926595923, 0.3575843393838781, 0.18048078840183424}, {0.21263900587151024, 0.7151686787677562, 0.0721923153607337}, {0.019330818715591818, 0.119194779794626, 0.9505321522496605}},
	}
	XYZtoRGB = map[string][3][3]float64{

		"rec709": {{3.2404542, -1.5371385, -0.4985314}, {-0.9692660, 1.8760108, 0.0415560}, {0.0556434, -0.2040259, 1.0572252}},
		"rec601": {{3.0628971, -1.3931791, -0.4757517}, {-0.9692660, 1.8760108, 0.0415560}, {0.0678775, -0.2288548, 1.0693490}},

		"test709": {{3.240969941904524, -1.5373831775700946, -0.49861076029300366}, {-0.9692436362808796, 1.87596750150772, 0.041555057407175626}, {0.05563007969699367, -0.20397695888897657, 1.0569715142428788}},
	}
)

func transform(input, output space, cols color.Color) color.Color {

	/*	if input.space == output.space {
		return cols
	}*/

	// else get transformation
	tf := getTransform(input, output)
	// apply transformatoin

	return tf(cols)
}


func getTransform(input, output space) func(color.Color) color.Color {

	/*

		get multiple transforms, how would luts differ from a matrix

		how do we allow multiple types of transformation e.g. clipping and transformation matrix

		keep one matri for everything. Make a way to have strings and strcuts as part of the library

	*/

	// Also generate a method of going from x to y with matrces

	if fc, ok := Library[input][output]; ok {
		return fc
	}
	fmt.Println(output, input)
	return matrixTransform(RGBToXYZ[input.space], XYZtoRGB[output.space])
}

func inverse(c color.Color) color.Color {
	r, g, b, a := c.RGBA()
	if a == 0xffff {
		return color.NRGBA64{0xffff - uint16(r), 0xffff - uint16(g), 0xffff - uint16(b), 0xffff}
	}
	if a == 0 {
		return color.NRGBA64{0, 0, 0, 0}
	}
	// Since Color.RGBA returns an alpha-premultiplied color, we should have r <= a && g <= a && b <= a.
	r = (r * 0xffff) / a
	g = (g * 0xffff) / a
	b = (b * 0xffff) / a
	return color.NRGBA64{0xffff - uint16(r), 0xffff - uint16(g), 0xffff - uint16(b), uint16(a)}

}

func RGBTransform() {
	// using matrixes to go to XYZ and then back again
}

func matrixTransform(xyz, rgb [3][3]float64) func(color.Color) color.Color {

	return func(input color.Color) color.Color {
		R, G, B, A := input.RGBA()
		fmt.Println(R, G, B)
		//	r, g, b := float32(R)/65535, float32(G)/65535, float32(B)/65535
		r, g, b := float64(R), float64(G), float64(B)
		X := r*xyz[0][0] + g*xyz[0][1] + b*xyz[0][2]
		Y := r*xyz[1][0] + g*xyz[1][1] + b*xyz[1][2]
		Z := r*xyz[2][0] + g*xyz[2][1] + b*xyz[2][2]

		aR := math.Round(X*rgb[0][0] + Y*rgb[0][1] + Z*rgb[0][2]) //* 65535
		aG := math.Round(X*rgb[1][0] + Y*rgb[1][1] + Z*rgb[1][2]) //* 65535
		aB := math.Round(X*rgb[2][0] + Y*rgb[2][1] + Z*rgb[2][2]) //* 65535
		fmt.Println(aR, aG, aB)
		cols := []*float64{&aR, &aG, &aB}

		// default transform function of clipping high
		// colours and crushing low colours
		for _, c := range cols {
			if *c > 65535.0 {
				*c = 65535
			} else if *c < 0 {
				*c = 0
			}
		}
		// preserve the alpha channel
		return color.NRGBA64{R: uint16(aR), G: uint16(aG), B: uint16(aB), A: uint16(A)}

	}
}
