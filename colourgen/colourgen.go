// Package colourgen generates rgb values
package colourgen

import (
	"fmt"
	"image/color"
	"regexp"
	"strings"
)

// AssignRGBValues takes a string and the rgb value and returns a [3]int array of the rgb values for a colour.
// The only valid colours are grey, red, green and blue,
// the black and white colours are constants for the time being.
func AssignRGBValues(colour string, rgb, maxBlack, maxWhite int) ([3]int, error) {
	switch strings.ToLower(colour) {
	case "grey", "gray": // "black", "white",

		return [3]int{rgb, rgb, rgb}, nil

	case "black":

		return [3]int{maxBlack, maxBlack, maxBlack}, nil
	case "white":

		return [3]int{maxWhite, maxWhite, maxWhite}, nil
	case "red":

		return [3]int{rgb, 0, 0}, nil
	case "green":

		return [3]int{0, rgb, 0}, nil
	case "blue":

		return [3]int{0, 0, rgb}, nil
	default:

		return [3]int{0, 0, 0}, fmt.Errorf("%s Non specific colour called, rgb values set at 0", colour)
	}
}

// HexToColour takes a string and returns a colour value, extracting the rgba values from the string. When no alpha channel
// is found the alpha is set to be the max.
//
// Acceptable formats are #rgb, #rgba, #rrggbb, ##rrggbbaa, rgb(r,g,b), rgba(r,g,b,a), rgb12(r,g,b) and rgba12(r,g,b,a)
//
// The resulting value is either color.NRGBA or color.NRBGA64, 12 bit RGB values are represented in 16 bit NRGBA64.
func HexToColour(colorCode string) color.Color {
	var colour color.NRGBA
	regRRGGBB := regexp.MustCompile(`^#[A-Fa-f0-9]{6}$`)
	regRGB := regexp.MustCompile(`^#[A-Fa-f0-9]{3}$`)
	regRRGGBBAA := regexp.MustCompile(`^#[A-Fa-f0-9]{8}$`)
	regRGBA := regexp.MustCompile(`^#[A-Fa-f0-9]{4}$`)
	regcssRGBA := regexp.MustCompile(`^(rgba\()\b([01]?[0-9][0-9]?|2[0-4][0-9]|25[0-5]),\b([01]?[0-9][0-9]?|2[0-4][0-9]|25[0-5]),\b([01]?[0-9][0-9]?|2[0-4][0-9]|25[0-5]),\b([01]?[0-9][0-9]?|2[0-4][0-9]|25[0-5])\)$`)
	regcssRGB := regexp.MustCompile(`^(rgb\()\b([01]?[0-9][0-9]?|2[0-4][0-9]|25[0-5]),\b([01]?[0-9][0-9]?|2[0-4][0-9]|25[0-5]),\b([01]?[0-9][0-9]?|2[0-4][0-9]|25[0-5])\)$`)
	regcssRGB12 := regexp.MustCompile(`^rgb12\(([0-3]?[0-9]{1,3}|40[0-9][0-5]),([0-3]?[0-9]{1,3}|40[0-9][0-5]),([0-3]?[0-9]{1,3}|40[0-9][0-5])\)$`)
	regcssRGBA12 := regexp.MustCompile(`^rgba12\(([0-3]?[0-9]{1,3}|40[0-9][0-5]),([0-3]?[0-9]{1,3}|40[0-9][0-5]),([0-3]?[0-9]{1,3}|40[0-9][0-5]),([0-3]?[0-9]{1,3}|40[0-9][0-5])\)$`)
	// break down the string an attribute each 2 hex value to the rgb

	// check length as all are unqiue>
	switch {
	case regRRGGBB.MatchString(colorCode):
		
		return rrggbb(colorCode, colour)
	case regRGB.MatchString(colorCode):

		return rgb(colorCode, colour)
	case regRRGGBBAA.MatchString(colorCode):

		return rrggbbaa(colorCode, colour)
	case regRGBA.MatchString(colorCode):

		return rgba(colorCode, colour)
	case regcssRGBA.MatchString(colorCode):

		return cssrgba(colorCode, colour)
	case regcssRGB.MatchString(colorCode):

		return cssrgb(colorCode, colour)
	case regcssRGB12.MatchString(colorCode):

		return cssrgb12(colorCode)
	case regcssRGBA12.MatchString(colorCode):

		return cssrgba12(colorCode)
	}

	return colour
}

func rrggbb(hex string, c color.NRGBA) color.NRGBA {
	fmt.Sscanf(hex, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	c.A = 0xff
	
	return c
}

func rgb(hex string, c color.NRGBA) color.NRGBA {
	fmt.Sscanf(hex, "#%01x%01x%01x", &c.R, &c.G, &c.B)
	c.A = 0xff
	c.R <<= 4
	c.G <<= 4
	c.B <<= 4

	return c
}

func rrggbbaa(hex string, c color.NRGBA) color.NRGBA {
	fmt.Sscanf(hex, "#%02x%02x%02x%02x", &c.R, &c.G, &c.B, &c.A)

	return c
}

func rgba(hex string, c color.NRGBA) color.NRGBA {
	fmt.Sscanf(hex, "#%01x%01x%01x%01x", &c.R, &c.G, &c.B, &c.A)
	c.R <<= 4
	c.G <<= 4
	c.B <<= 4
	c.A <<= 4

	return c
}

func cssrgba(css string, c color.NRGBA) color.NRGBA {
	fmt.Sscanf(css, "rgba(%v,%v,%v,%v)", &c.R, &c.G, &c.B, &c.A)

	return c
}

func cssrgb(css string, c color.NRGBA) color.NRGBA {
	fmt.Sscanf(css, "rgb(%v,%v,%v)", &c.R, &c.G, &c.B)
	c.A = 0xff

	return c
}

func cssrgb12(css string) color.NRGBA64 {
	R, G, B := uint16(0), uint16(0), uint16(0)
	fmt.Sscanf(css, "rgb12(%v,%v,%v)", &R, &G, &B)
	c := color.NRGBA64{R << 4, G << 4, B << 4, 0xffff}

	return c
}

func cssrgba12(css string) color.NRGBA64 {
	R, G, B, A := uint16(0), uint16(0), uint16(0), uint16(0)
	fmt.Sscanf(css, "rgba12(%v,%v,%v,%v)", &R, &G, &B, &A)
	c := color.NRGBA64{R << 4, G << 4, B << 4, A << 4}

	return c
}

// ConvertNRGBA64 converts any colour into an NRGBA64 colour.
// The colours are returned as 8 bit colours shifted to 16 bit, unless already a color.NRGBA64 then that is returned without change.
// This is designed to preserve the colours 8 bit representation and not change the original value to be a warped 16 bit
// representation.
//
// Max 8 bit alpha channel values of 255 are set to max 16 bit values of 65535. This is to make solid images.
//
// e.g. rgb(250,4,100) becomes rgb(64000,1024,25600) in NRGBA64 form
func ConvertNRGBA64(original color.Color) color.NRGBA64 {
	if col, ok := original.(color.NRGBA64); ok {
		return col
	}

	nrgba := color.NRGBAModel.Convert(original).(color.NRGBA)

	convert := color.NRGBA64{R: uint16(nrgba.R) << 8, G: uint16(nrgba.G) << 8, B: uint16(nrgba.B) << 8}
	if nrgba.A == 0xff { // If alpha is the max value for an 8 bit number
		// round up to make solid for 16 bit images
		convert.A = 0xffff
	} else {
		convert.A = uint16(nrgba.A) << 8
	}

	return convert
}
