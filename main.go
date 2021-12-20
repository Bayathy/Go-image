package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	img, _ := png.Decode(os.Stdin)
	bounds := img.Bounds()
	dest := image.NewGray16(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := color.Gray16Model.Convert(img.At(x, y))
			gray, _ := c.(color.Gray16)
			dest.Set(x, y, gray)
		}
	}
	png.Encode(os.Stdout, dest)
}
