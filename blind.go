// Copyright © 2023 J. Salvador Arias <jsalarias@gmail.com>.
// All rights reserved.
// Distributed under BSD2 licenses that can be found in the LICENSE file.

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

// [Iridescent] color scheme.
//
// [Iridescent]: https://personal.sron.nl/~pault/#fig:scheme_iridescent
var iridescent = []color.RGBA{
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

// Gradient returns a color gradient from low to high.
// The input value must be scaled from 0 to 1.
//
// It is a continuous version of the [Iridescent]
// color scheme.
// Darker values indicate higher values
//
// [Iridescent]: https://personal.sron.nl/~pault/#fig:scheme_iridescent
func Gradient(v float64) color.RGBA {
	scaled := v * float64(len(iridescent)-1)
	pos := math.Floor(scaled)
	delta := scaled - pos

	c := iridescent[int(pos)]
	if int(pos) == len(iridescent)-1 {
		return c
	}

	next := iridescent[int(pos)+1]

	r := float64(c.R) + (float64(next.R)-float64(c.R))*delta
	g := float64(c.G) + (float64(next.G)-float64(c.G))*delta
	b := float64(c.B) + (float64(next.B)-float64(c.B))*delta
	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}
}
