package colour

import "image/color"

var (
	Library = map[space]map[space]func(color.Color) color.Color{
		{space: "nothing"}: {
			space{space: "inverse"}: inverse},
	}
)

func transform(input, output space, cols color.Color) color.Color {

	if input.space == output.space {
		return cols
	}

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

	return Library[input][output]
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
