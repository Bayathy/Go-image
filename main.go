package main

import (
	"image"
	_ "image"
	_ "image/color"
	"image/png"
	"os"
)

func main() {
	in, _ := os.Open("momiji.png")
	out, _ := os.Create("momijicopy.png")
	img, _ := png.Decode(in)
	bounds := img.Bounds()

	//dest := image.NewGray16(bounds)
	dest := image.NewRGBA(bounds)


	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
					c := img.At(x,y)
					dest.Set(x, y, c)
			}
	}
	png.Encode(out, dest)
}
