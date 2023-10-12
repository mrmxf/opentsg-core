package colour

import (
	"image"
	"image/color"
	"image/draw"
)

var (
	gray40   = color.NRGBA64{R: 414 << 6, G: 414 << 6, B: 414 << 6, A: 0xffff}
	white75  = color.NRGBA64{R: 721 << 6, G: 721 << 6, B: 721 << 6, A: 0xffff}
	yellow75 = color.NRGBA64{R: 721 << 6, G: 721 << 6, B: 64 << 6, A: 0xffff}
	cyan75   = color.NRGBA64{R: 64 << 6, G: 721 << 6, B: 721 << 6, A: 0xffff}
	green75  = color.NRGBA64{R: 64 << 6, G: 721 << 6, B: 64 << 6, A: 0xffff}
	mag75    = color.NRGBA64{R: 721 << 6, G: 64 << 6, B: 721 << 6, A: 0xffff}
	red75    = color.NRGBA64{R: 721 << 6, G: 64 << 6, B: 64 << 6, A: 0xffff}
	blue75   = color.NRGBA64{R: 64 << 6, G: 64 << 6, B: 721 << 6, A: 0xffff}
)

const (
	//widths
	d = 240 / 1920.0
	f = 205 / 1920.0
	c = 206 / 1920.0
	b = 1 / 12.0
	k = 309 / 1920.0
	g = 411 / 1920.0
	h = 171 / 1920.0
)

func (br bar) generate(canvas WidgetImage) {
	b := canvas.Bounds().Max
	w := 0.0
	twidth := 0.0

	fills := []bars{{width: d, color: gray40}, {width: f, color: white75}, {width: c, color: yellow75}, {width: c, color: cyan75}, {width: c, color: green75}, {width: c, color: mag75}, {width: c, color: red75}, {width: f, color: blue75}, {width: d, color: gray40}}

	for _, f := range fills {
		twidth += f.width * float64(b.X)
		area := image.Rect(int(w), int(0), int(w+f.width*float64(b.X)), b.Y)

		canvas.Draw(area, f.color, draw.Src, br.Space)
		//draw.Draw(canvas, area, fill, image.Point{}, draw.Src)

		w += f.width * float64(b.X)
	}

}
