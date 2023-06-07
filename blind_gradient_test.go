// Color schemes copyright © 2022 Paul Tol <https://personal.sron.nl/~pault/>
// Code copyright © 2023 J. Salvador Arias <jsalarias@gmail.com>.
// All rights reserved.
// Distributed under BSD3 licenses that can be found in the LICENSE file.

package blind_test

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/js-arias/blind"
)

const size = 500

type gradient struct{}

func (g gradient) ColorModel() color.Model { return color.RGBAModel }
func (g gradient) Bounds() image.Rectangle { return image.Rect(0, 0, size, 100) }
func (g gradient) At(x, y int) color.Color {
	v := float64(x) / (size - 1)
	return blind.Gradient(v)
}

// This example produces a gradient scale.
func Example_gradient() {
	f, err := os.Create("gradient.png")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
		fmt.Printf("ok")
	}()

	if err := png.Encode(f, gradient{}); err != nil {
		panic(err)
	}

	// Output:
	// ok
}
