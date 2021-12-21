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

	var (
		i      uint8
		j      uint8
		n      uint8
		crange uint8
		srange uint8
	)

	crange = 20
	srange = 70

	for i = 0; i < crange; i++ {
		for j = 0; j < crange; j++ {
			for n = 0; n < crange; n++ {
				wtcc := color.RGBA{srange + i, srange + j, srange + n, 255}

				wtclist = append(wtclist, wtcc)
			}
		}
	}

	aclist := [...]color.RGBA{
		//{170,94,102,255},
		{170, 44, 48, 255},
	}

	img, _ := png.Decode(in)
	bounds := img.Bounds()

	fmt.Println(len(wtclist))

	//dest := image.NewGray16(bounds)
	dest := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			color := img.At(x, y)
			for i := 0; i < len(wtclist); i++ {
				if color == wtclist[i] && y < 400 {
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
