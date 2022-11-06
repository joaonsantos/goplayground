package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
	"strconv"
)

func convulsion(matrix [][]color.RGBA, mask [][]uint8, size int) color.RGBA {
	var (
		r uint32
		g uint32
		b uint32
		a uint32
	)
	div := uint32(size * size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			c := matrix[i][j]
			m := mask[i][j]

			r += uint32(c.R * m)
			g += uint32(c.G * m)
			b += uint32(c.B * m)
			a += uint32(c.A * m)
		}
	}
	return color.RGBA{
		R: uint8(r / div),
		G: uint8(g / div),
		B: uint8(b / div),
		A: uint8(a / div),
	}
}

func blurPix(src, dst *image.RGBA, mask [][]uint8, maskSize, x, y int) {
	offset := maskSize / 2
	b := src.Bounds()
	if x <= offset || y <= offset || x >= b.Dx()-offset || y >= b.Dy()-offset {
		// do nothing
		return
	}

	// create rgba numbers matrix size of mask
	m := make([][]color.RGBA, maskSize)
	for i := 0; i < maskSize; i++ {
		m[i] = make([]color.RGBA, maskSize)

		for j := 0; j < maskSize; j++ {
			x := (x - offset) + i
			y := (y - offset) + j

			m[i][j] = src.RGBAAt(x, y)
		}
	}

	// do convulsion with mask and save
	rgba := convulsion(m, mask, maskSize)
	dst.SetRGBA(x, y, rgba)
}

func createBlurMask(size int) [][]uint8 {
	m := make([][]uint8, size)
	for i := 0; i < size; i++ {
		m[i] = make([]uint8, size)
		for j := 0; j < size; j++ {
			m[i][j] = 1
		}
	}
	return m
}

func main() {
	f, err := os.OpenFile("in.jpg", os.O_RDONLY, 0o0400)
	if err != nil {
		panic(err)
	}
	img, err := jpeg.Decode(f)
	if err != nil {
		panic(err)
	}
	f.Close()

	b := img.Bounds()
	src := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	dst := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))

	draw.Draw(src, src.Bounds(), img, b.Min, draw.Src)
	draw.Draw(dst, dst.Bounds(), img, b.Min, draw.Src)

	maskSize := 3
	if args := os.Args; len(os.Args) > 1 {
		maskSize, err = strconv.Atoi(args[1])
		if err != nil {
			panic(err)
		}
	}
	mask := createBlurMask(maskSize)

	for i := 0; i < b.Dx(); i++ {
		for j := 0; j < b.Dy(); j++ {
			blurPix(src, dst, mask, maskSize, i, j)
		}
	}

	f, err = os.OpenFile("out.jpg", os.O_CREATE|os.O_WRONLY, 0o0600)
	if err != nil {
		panic(err)
	}
	jpeg.Encode(f, dst, nil)

}
