package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
)

func main() {
	in, _ := os.Open("momiji.png")
	out, _ := os.Create("momijicopy.png")

	wtclist := []color.RGBA{}

	var i uint8

	for num := 0; num < 3; num++ {
		for i = 0; i < 120; i++ {
			switch num {
			case 0:
				wtcc := color.RGBA{60 + i, 60, 60, 255}
				wtclist = append(wtclist, wtcc)
			case 1:
				wtcc := color.RGBA{60, 60 + i, 60, 255}
				wtclist = append(wtclist, wtcc)
			case 2:
				wtcc := color.RGBA{60, 60, 60 + i, 255}
				wtclist = append(wtclist, wtcc)
			}
		}
	}

	aclist := [...]color.RGBA{
		{226, 71, 29, 255},
		{230, 20, 14, 255},
		{230, 99, 82, 255},
		{255, 255, 255, 255},
		{246, 158, 44, 255},
		{222, 196, 67, 255},
		{212, 170, 109, 255},
		{255, 255, 255, 255},
		{154, 123, 85, 255},
		{218, 139, 64, 255},
		{208, 132, 104, 255},
		{255, 255, 255, 255},
	}

	_ = aclist

	img, _ := png.Decode(in)
	bounds := img.Bounds()

	fmt.Println(len(wtclist))

	//dest := image.NewGray16(bounds)
	dest := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			color := img.At(x, y)
			for i := 0; i < len(wtclist); i++ {
				if color == wtclist[i] {
					//fmt.Println("success")
					color = aclist[rand.Intn(len(aclist))]
				}
			}
			//fmt.Println(color)
			dest.Set(x, y, color)
		}
	}
	png.Encode(out, dest)
}
