// Color schemes copyright © 2022 Paul Tol <https://personal.sron.nl/~pault/>
// Code copyright © 2023 J. Salvador Arias <jsalarias@gmail.com>.
// All rights reserved.
// Distributed under BSD3 licenses that can be found in the LICENSE file.

// Package `blind` implement color gradients
// that are friendly for color blind people.
//
// The color schemes are taken from [Paul Tol page].
//
// [Paul Tol page]: https://personal.sron.nl/~pault/
package blind

import (
	"image/color"
	"math"
)

// A ColSeq is a discrete step color sequence.
type ColSeq []color.RGBA

// Sequential color schemes
var (
	// Incandescent color scheme:
	// https://personal.sron.nl/~pault/#fig:scheme_incandescent.
	//
	// This scheme is not print friendly.
	Incandescent = ColSeq{
		{R: 206, G: 255, B: 255, A: 255},
		{R: 198, G: 247, B: 214, A: 255},
		{R: 162, G: 244, B: 155, A: 255},
		{R: 187, G: 228, B: 83, A: 255},
		{R: 213, G: 206, B: 4, A: 255},
		{R: 231, G: 181, B: 3, A: 255},
		{R: 241, G: 153, B: 3, A: 255},
		{R: 246, G: 121, B: 11, A: 255},
		{R: 249, G: 73, B: 2, A: 255},
		{R: 228, G: 5, B: 21, A: 255},
		{R: 168, G: 0, B: 3, A: 255},
	}

	// Iridescent color scheme:
	// https://personal.sron.nl/~pault/#fig:scheme_iridescent.
	Iridescent = ColSeq{
		{R: 254, G: 251, B: 233, A: 255},
		{R: 252, G: 247, B: 213, A: 255},
		{R: 245, G: 243, B: 193, A: 255},
		{R: 234, G: 240, B: 181, A: 255},
		{R: 221, G: 236, B: 191, A: 255},
		{R: 208, G: 231, B: 202, A: 255},
		{R: 194, G: 227, B: 210, A: 255},
		{R: 181, G: 221, B: 216, A: 255},
		{R: 168, G: 216, B: 220, A: 255},
		{R: 155, G: 210, B: 225, A: 255},
		{R: 141, G: 203, B: 228, A: 255},
		{R: 129, G: 196, B: 231, A: 255},
		{R: 123, G: 188, B: 231, A: 255},
		{R: 126, G: 178, B: 228, A: 255},
		{R: 136, G: 165, B: 221, A: 255},
		{R: 147, G: 152, B: 210, A: 255},
		{R: 155, G: 138, B: 196, A: 255},
		{R: 157, G: 125, B: 178, A: 255},
		{R: 154, G: 112, B: 158, A: 255},
		{R: 144, G: 99, B: 136, A: 255},
		{R: 128, G: 87, B: 112, A: 255},
		{R: 104, G: 73, B: 87, A: 255},
		{R: 70, G: 53, B: 58, A: 255},
	}

	// Smooth rainbow scheme:
	// https://personal.sron.nl/~pault/#fig:scheme_rainbow_smooth.
	//
	// This scheme does not have to be used over the full range.
	Rainbow = ColSeq{
		{R: 232, G: 236, B: 251, A: 255},
		{R: 221, G: 216, B: 239, A: 255},
		{R: 209, G: 193, B: 225, A: 255},
		{R: 195, G: 168, B: 209, A: 255},
		{R: 181, G: 143, B: 194, A: 255},
		{R: 167, G: 120, B: 180, A: 255},
		{R: 155, G: 98, B: 167, A: 255},
		{R: 140, G: 78, B: 153, A: 255},
		{R: 111, G: 76, B: 155, A: 255}, // purple
		{R: 96, G: 89, B: 169, A: 255},
		{R: 85, G: 104, B: 184, A: 255},
		{R: 78, G: 121, B: 197, A: 255},
		{R: 77, G: 138, B: 198, A: 255},
		{R: 78, G: 150, B: 188, A: 255},
		{R: 84, G: 158, B: 179, A: 255},
		{R: 89, G: 165, B: 169, A: 255},
		{R: 96, G: 171, B: 158, A: 255},
		{R: 105, G: 177, B: 144, A: 255},
		{R: 119, G: 183, B: 125, A: 255},
		{R: 140, G: 188, B: 104, A: 255},
		{R: 166, G: 190, B: 84, A: 255},
		{R: 190, G: 188, B: 72, A: 255},
		{R: 209, G: 181, B: 65, A: 255},
		{R: 221, G: 170, B: 60, A: 255},
		{R: 228, G: 156, B: 57, A: 255},
		{R: 231, G: 140, B: 53, A: 255},
		{R: 230, G: 121, B: 50, A: 255},
		{R: 228, G: 99, B: 45, A: 255},
		{R: 223, G: 72, B: 40, A: 255},
		{R: 218, G: 34, B: 34, A: 255}, // red
		{R: 184, G: 34, B: 30, A: 255},
		{R: 149, G: 33, B: 27, A: 255},
		{R: 114, G: 30, B: 23, A: 255},
		{R: 82, G: 26, B: 19, A: 255},
	}

	// Smooth rainbow scheme from purple to red
	RainbowPurpleToRed = Rainbow[8:30]
)

// Gradient returns a color gradient from low to high.
// The input value must be scaled from 0 to 1.
//
// It is a continuous version of the [Rainbow (purple to red)]
// color scheme.
// Small values are in purple
// large values are in red.
//
// [Rainbow (purple to red)]: https://personal.sron.nl/~pault/#fig:scheme_rainbow_smooth
func Gradient(v float64) color.RGBA {
	return Sequential(RainbowPurpleToRed, v)
}

// Sequential returns a color gradient from low to high,
// interpolating from the indicated color sequence.
// The input value must be scaled from 0 to 1.
func Sequential(seq ColSeq, v float64) color.RGBA {
	scaled := v * float64(len(seq)-1)
	pos := math.Floor(scaled)
	delta := scaled - pos

	c := seq[int(pos)]
	if int(pos) == len(seq)-1 {
		return c
	}

	next := seq[int(pos)+1]

	r := float64(c.R) + (float64(next.R)-float64(c.R))*delta
	g := float64(c.G) + (float64(next.G)-float64(c.G))*delta
	b := float64(c.B) + (float64(next.B)-float64(c.B))*delta
	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}
}
